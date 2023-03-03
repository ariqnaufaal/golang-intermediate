package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"

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

	const SSH_ADDRESS = "0.0.0.0:22"
	const SSH_USERNAME = "testsftp"
	const SSH_PASSWORD = "password"
	// const SSH_KEY = "path/to/file/identity.pem"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}

	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	// create session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session. " + err.Error())
	}
	//session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// test session
	stdinBuf, _ := session.StdinPipe()
	if err := session.Shell(); err != nil {
		panic(err)
	}
	stdinBuf.Write([]byte("echo hello\n"))
	stdinBuf.Write([]byte("ls -l ~/\n"))
	// err = session.Run("ls -l ~/")
	// if err != nil {
	// 	log.Fatal("Command execution error. " + err.Error())
	// }

	// test sftp
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed create client sftp client. " + err.Error())
	}

	//err = session.Run("touch ~/test-file.txt")
	//if err != nil {
	//	log.Fatal("Command execution error. " + err.Error())
	//}test_dir
	//session.Close()

	/*
		// create file with os library, change permissions and write text in file
			dir := "test_dir"
			os.Mkdir(dir, 0777)
			fileName := path.Join(dir, "test-file.txt")

			ioutil.WriteFile(fileName, []byte("test"), 0666)
	*/
	// D:/Multipolar-Projects/Training Golang Intermediate Programming/GLIM_Hacktiv8/golang-intermediate/session-7/test_dir/

	fDestination, err := sftpClient.Create("test-file.txt")
	if err != nil {
		log.Fatal("Failed to create destination file. " + err.Error())
	}

	fSource, err := os.Open("test-file.txt")
	if err != nil {
		log.Fatal("Failed to read source file. " + err.Error())
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Fatal("Failed copy source file into destination file. " + err.Error())
	}

	log.Println("File copied.")

	// create new session
}

/*

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

*/
