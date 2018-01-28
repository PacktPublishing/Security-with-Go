package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
)

func checkArgs() (string, string, string) {
	if len(os.Args) != 4 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1], os.Args[2], os.Args[3]
}

func printUsage() {
	fmt.Println(os.Args[0] + ` - Open an SSH shell

Usage:
  ` + os.Args[0] + ` <username> <host> <privateKeyFile>

Example:
  ` + os.Args[0] + ` nanodano devdungeon.com:22 ~/.ssh/id_rsa
`)
}

func getKeySigner(privateKeyFile string) ssh.Signer {
	privateKeyData, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal("Error loading private key file. ", err)
	}

	privateKey, err := ssh.ParsePrivateKey(privateKeyData)
	if err != nil {
		log.Fatal("Error parsing private key. ", err)
	}
	return privateKey
}

func main() {
	username, host, privateKeyFile := checkArgs()

	privateKey := getKeySigner(privateKeyFile)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server. ", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Pipe the standard buffers together
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	// Get psuedo-terminal
	err = session.RequestPty(
		"vt100", // or "linux", "xterm"
		40,      // Height
		80,      // Width
		// https://godoc.org/golang.org/x/crypto/ssh#TerminalModes
		// POSIX terminal mode flags defined in RFC 4254 Section 8.
		// https://tools.ietf.org/html/rfc4254#section-8
		ssh.TerminalModes{
			ssh.ECHO: 0,
		})
	if err != nil {
		log.Fatal("Error requesting psuedo-terminal. ", err)
	}

	// Run shell until it is exited
	err = session.Shell()
	if err != nil {
		log.Fatal("Error executing command. ", err)
	}
	session.Wait()
}
