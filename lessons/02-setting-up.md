# 2. Setting up the Go Environment

Now that you have an idea of what Go is, it's time to set up your development environment. In this lesson, we will walk you through the steps to install Go on your system and configure your workspace.

## Installing Go

The first step is to download and install the Go distribution for your operating system. You can find the official installation packages on the [Go downloads page](https://go.dev/dl/).

### Windows

1.  Download the MSI installer for Windows.
2.  Run the installer and follow the on-screen instructions. By default, Go will be installed in `C:\Go`.
3.  The installer will automatically add the `C:\Go\bin` directory to your system's PATH environment variable.

### macOS

1.  Download the package installer for macOS.
2.  Run the installer and follow the instructions. Go will be installed in `/usr/local/go`.
3.  The installer will add the `/usr/local/go/bin` directory to your PATH environment variable.

### Linux

1.  Download the tarball for Linux.
2.  Extract the archive to `/usr/local` using the following command (you may need to run it as root):
    ```bash
    tar -C /usr/local -xzf go<version>.linux-amd64.tar.gz
    ```
3.  Add `/usr/local/go/bin` to your PATH environment variable. You can do this by adding the following line to your `$HOME/.profile` or `/etc/profile`:
    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```

## Verifying the Installation

Once the installation is complete, you can verify it by opening a new terminal or command prompt and running the following command:

```bash
go version
```

This command should print the installed version of Go, for example:

```
go version go1.24.3 linux/amd64
```

## Your First Go Program

Now that you have Go installed, let's write a simple "Hello, World!" program to make sure everything is working correctly.

1.  Create a new directory for your project, for example, `hello`.
2.  Navigate into the `hello` directory and create a new file named `main.go`.
3.  Open `main.go` in your favorite text editor and add the following code:

    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
    ```

4.  To run the program, open a terminal in the `hello` directory and use the `go run` command:

    ```bash
    go run main.go
    ```

You should see the output `Hello, World!` printed to the console.

## Go Workspace

In older versions of Go, you had to set up a specific workspace structure with `GOPATH`. With the introduction of Go modules, this is no longer a strict requirement. You can now create your Go projects in any directory on your system.

When you run `go mod init <module_name>` in your project's root directory, Go creates a `go.mod` file that tracks your project's dependencies. We will cover Go modules in more detail in a later lesson.

## Recommended Tools

To enhance your Go development experience, we recommend the following tools:

*   **Visual Studio Code (VS Code):** A popular and lightweight code editor with excellent support for Go. You can install the official [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=golang.Go) extension, which provides features like IntelliSense, code navigation, and debugging.
*   **GoLand:** A full-featured IDE for Go development from JetBrains. It offers advanced features like code completion, refactoring, and a built-in debugger.

Congratulations! You now have a working Go development environment. In the next lesson, we will dive into the basic syntax of the Go language.
