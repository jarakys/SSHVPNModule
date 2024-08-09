package SSHVPNModule

import (
	"bytes"
	"golang.org/x/crypto/ssh"
)

func ExecuteSSHCommandWithMultipleInputs(client *ssh.Client, command string, inputs []string) (string, error) {
	session, err := client.NewSession()
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
