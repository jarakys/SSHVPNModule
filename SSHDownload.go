package SSHVPNModule

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func DownloadFileSSH(client *ssh.Client, file string) ([]byte, error) {
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s\n", err)
		return []byte{}, err
	}
	defer session.Close()

	// Create a pipe to capture the remote file's content
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	// Execute the command to read the file content
	catCommand := fmt.Sprintf("sudo cat %s", file+".ovpn")
	if err := session.Run(catCommand); err != nil {
		fmt.Printf("Failed to run command: %s\n", err)
		return nil, err
	}

	// Get the file content as a byte slice
	fileContent := stdoutBuf.Bytes()

	// Optional: Write the content to a local file (for demonstration purposes)
	fmt.Println("File downloaded successfully")
	fmt.Printf("File content: %s\n", string(fileContent))

	return fileContent, nil
}
