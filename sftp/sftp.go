package sftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	host    = os.Getenv("GALLERY_VM_PUBLIC_GO_IP")
	port    = os.Getenv("GALLERY_VM_PUBLIC_GO_PORT")
	userVM  = os.Getenv("GALLERY_VM_PUBLIC_GO_USER")
	privKey = os.Getenv("GALLERY_VM_PUBLIC_GO_PRIVKEY_ADDR")
)

func SftpConnect() (*sftp.Client, error) {

	priv, err := ioutil.ReadFile(privKey)
	check(err)
	signer, _ := ssh.ParsePrivateKey(priv)
	config := &ssh.ClientConfig{
		User: userVM,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connect to ssh
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := ssh.Dial("tcp", addr, config)
	check(err)

	// sftp client
	client, err := sftp.NewClient(conn)
	check(err)

	return client, err
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}
