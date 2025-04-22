# 🌐 SSHVPNModule

**A Go package for secure SSH connections, command execution, and file downloads.**

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## 📖 Table of Contents

- [About](#about)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Creating an SSH Client](#creating-an-ssh-client)
  - [Executing Commands](#executing-commands)
  - [Downloading Files](#downloading-files)
- [Example](#example)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)

---

## ℹ️ About

`SSHVPNModule` is a lightweight Go package designed to streamline SSH operations. It enables secure connections to remote servers, executes commands with multiple inputs, and downloads files effortlessly. Built on `golang.org/x/crypto/ssh`, it’s perfect for developers needing reliable SSH functionality in their Go applications.

---

## 🚀 Features

- 🔑 **Secure Authentication**: Password-based SSH access with robust security.
- 🖥️ **Interactive Commands**: Execute commands with multiple stdin inputs.
- 📂 **File Downloads**: Fetch remote files using `sudo cat` for privileged access.
- 🛡️ **Error Handling**: Comprehensive error management for dependable operations.
- ⚡ **Simple API**: Intuitive interface for quick integration.

---

## 🛠️ Installation

1. Ensure you have **Go 1.18+** installed.
2. Install the required dependency:

   ```bash
   go get golang.org/x/crypto/ssh
   ```

3. Add `SSHVPNModule` to your project:

   ```bash
   go get github.com/your-repo/SSHVPNModule
   ```

---

## 📚 Usage

### Creating an SSH Client

Initialize an SSH client with server credentials:

```go
package main

import (
    "fmt"
    "your-repo/SSHVPNModule"
)

func main() {
    client, err := SSHVPNModule.NewSSHClient("password", "username", "server-ip", "22")
    if err != nil {
        fmt.Printf("Failed to create SSH client: %s\n", err)
        return
    }
    defer client.(*SSHVPNModule.sshClientImpl).client.Close()
}
```

### Executing Commands

Run a command with multiple inputs:

```go
command := "some-interactive-command"
inputs := []string{"input1", "input2", "input3"}
output, err := client.ExecuteSSHCommandWithMultipleInputs(command, inputs)
if err != nil {
    fmt.Printf("Command execution failed: %s\n", err)
    return
}
fmt.Printf("Command output: %s\n", output)
```

### Downloading Files

Download a file from the remote server:

```go
fileContent, err := client.DownloadFileSSH("/path/to/remote/file")
if err != nil {
    fmt.Printf("File download failed: %s\n", err)
    return
}
fmt.Printf("Downloaded file content: %s\n", string(fileContent))
```

---

## 💻 Example

A complete example showcasing `SSHVPNModule`:

```go
package main

import (
    "fmt"
    "your-repo/SSHVPNModule"
)

func main() {
    // Initialize SSH client
    client, err := SSHVPNModule.NewSSHClient("my-password", "my-username", "192.168.1.100", "22")
    if err != nil {
        fmt.Printf("Failed to create SSH client: %s\n", err)
        return
    }
    defer client.(*SSHVPNModule.sshClientImpl).client.Close()

    // Execute a command
    command := "echo 'Hello' && read input"
    inputs := []string{"World"}
    output, err := client.ExecuteSSHCommandWithMultipleInputs(command, inputs)
    if err != nil {
        fmt.Printf("Command execution failed: %s\n", err)
        return
    }
    fmt.Printf("Command output: %s\n", output)

    // Download a file
    fileContent, err := client.DownloadFileSSH("/etc/hostname")
    if err != nil {
        fmt.Printf("File download failed: %s\n", err)
        return
    }
    fmt.Printf("File content: %s\n", string(fileContent))
}
```

---

## 📦 Dependencies

- [`golang.org/x/crypto/ssh`](https://pkg.go.dev/golang.org/x/crypto/ssh): Core SSH protocol implementation.

Install it with:

```bash
go get golang.org/x/crypto/ssh
```

---

## 🤝 Contributing

We’d love your contributions! Follow these steps:

1. 🍴 Fork the repository.
2. 🌿 Create a feature branch (`git checkout -b feature/YourFeature`).
3. 💾 Commit your changes (`git commit -m 'Add YourFeature'`).
4. 🚀 Push to the branch (`git push origin feature/YourFeature`).
5. 📬 Open a Pull Request.

Please adhere to the project’s coding standards and include tests where possible.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

🌟 **Star this repo** if you find it helpful! Report issues or suggest improvements via GitHub Issues. Let’s build something great together!
