package SSHVPNModule

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

type SSHClient interface {
	ExecuteSSHCommandWithMultipleInputs(command string, inputs []string) (string, error)
	DownloadFileSSH(file string) ([]byte, error)
}

type sshClientImpl struct {
	client *ssh.Client
}

func NewSSHClient(password string, login string, serverIp string, port string) (SSHClient, error) {
	config := &ssh.ClientConfig{
		User: login,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:" + port, serverIp), config)
	if err != nil {
		return nil, err
	}
	return &sshClientImpl{client: client}, nil
}

func (s *sshClientImpl) ExecuteSSHCommandWithMultipleInputs(command string, inputs []string) (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	stdin, err := session.StdinPipe()
	if err != nil {
		return "", err
	}

	err = session.Start(command)
	if err != nil {
		return "", err
	}

	// Отправка данных во stdin
	for _, input := range inputs {
		_, err = stdin.Write([]byte(input + "\n"))
		if err != nil {
			return "", err
		}
	}
	stdin.Close()

	err = session.Wait()
	if err != nil {
		return "", err
	}

	return stdoutBuf.String(), nil
}

func (s *sshClientImpl) DownloadFileSSH(file string) ([]byte, error) {
	session, err := s.client.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s\n", err)
		return []byte{}, err
	}
	defer session.Close()

	// Create a pipe to capture the remote file's content
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	// Execute the command to read the file content
	catCommand := fmt.Sprintf("sudo cat %s", file)
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
