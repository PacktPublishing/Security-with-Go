package main

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

var username = "username"
var host = "example.com:22"
var privateKeyFile = "/home/user/.ssh/id_rsa"

// Known hosts only reads FIRST entry
var knownHostsFile = "/home/user/.ssh/known_hosts"

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

func loadServerPublicKey(knownHostsFile string) ssh.PublicKey {
	publicKeyData, err := ioutil.ReadFile(knownHostsFile)
	if err != nil {
		log.Fatal("Error loading server public key file. ", err)
	}

	_, _, publicKey, _, _, err := ssh.ParseKnownHosts(publicKeyData)
	if err != nil {
		log.Fatal("Error parsing server public key. ", err)
	}
	return publicKey
}

func main() {
	userPrivateKey := getKeySigner(privateKeyFile)
	serverPublicKey := loadServerPublicKey(knownHostsFile)

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(userPrivateKey),
		},
		HostKeyCallback: ssh.FixedHostKey(serverPublicKey),
		// Acceptable host key algorithms (Allow all)
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server. ", err)
	}

	log.Println(string(client.ClientVersion()))
}
