package sftp

import (
	"testing"
)

func TestSftpConnect(t *testing.T) {

	t.Log("Don't log secrets.  Don't do it.\n")

	client, err := SftpConnect()
	check(err)
	defer client.Close()

	pwd, err := client.Getwd()
	check(err)
	t.Log("Current directory: ", pwd)
}
