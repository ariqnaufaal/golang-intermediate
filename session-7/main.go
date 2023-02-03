package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	/*
		const SSH_ADDRESS = "0.0.0.0:22"
		const SSH_USERNAME = "user"
		const SSH_PASSWORD = "password"

		sshConfig := &ssh.ClientConfig{
			User:            SSH_USERNAME,
			HostKeyCallback: ssh.IInsecureIgnoreHostKey(),
			Auth: []ssh.AuthMethod{
				ssh.Password(SSH_PASSWORD),
			},
		}
	*/

	// Otentikasi menggunakan identity file

	const SSH_ADDRESS = "192.168.0.24:22"
	const SSH_USERNAME = "user"
	const SSH_KEY = "path/to/file/identity.pem"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			PublicKeyFile(SSH_KEY),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}

	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}
}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParseDSAPrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}
