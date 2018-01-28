package main

import (
	"golang.org/x/crypto/ssh"
	"log"
)

var username = "username"
var password = "password"
var host = "example.com:22"

func main() {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server. ", err)
	}

	log.Println(string(client.ClientVersion()))
}
