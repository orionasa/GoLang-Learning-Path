# Module 1: Introduction & Setup

Welcome to the first module of our GoLang learning path! In this module, we'll cover the basics of Go, including why it's a popular choice for developers, how to set up your development environment, and how to write your first Go program.

## Why Go?

Go, also known as GoLang, is an open-source programming language developed by Google. It's known for its simplicity, efficiency, and strong support for concurrent programming. Here are a few reasons why developers choose Go:

*   **Concurrency:** Go's concurrency model, which uses goroutines and channels, makes it easy to write programs that can handle many tasks at once. This is especially useful for building network services and web backends.
*   **Performance:** Go is a compiled language, which means it's translated directly into machine code. This results in fast execution speeds, comparable to languages like C++ and Java.
*   **Simplicity:** Go has a clean and simple syntax that's easy to learn. It has a small standard library and a limited number of language features, which helps to keep code readable and maintainable.

## Installing and configuring the Go toolchain

Before you can start writing Go code, you'll need to install the Go toolchain on your computer. You can download the latest version of Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)

Once you've installed Go, you'll need to set up your `GOPATH` and `GOMODULES` environment variables.

*   `GOPATH` is the path to your workspace, where your Go code and dependencies will be stored.
*   `GOMODULES` is a feature that allows you to manage dependencies for your projects.

You can set these environment variables in your shell's configuration file (e.g., `~/.bashrc` or `~/.zshrc`).

```bash
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

## Your First Program: "Hello, World!"

Now that you've set up your development environment, you're ready to write your first Go program. Create a new file called `main.go` and add the following code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

This program defines a `main` package and imports the `fmt` package, which provides functions for formatting and printing text. The `main` function is the entry point of the program, and it uses the `fmt.Println` function to print "Hello, World!" to the console.

## Using the Go CLI

The Go command-line interface (CLI) is a powerful tool for working with Go code. Here are a few common commands:

*   `go run`: Compiles and runs a Go program.
*   `go build`: Compiles a Go program into an executable file.
*   `go fmt`: Formats Go code according to the language's style guidelines.

To run your "Hello, World!" program, open a terminal and navigate to the directory where you saved `main.go`. Then, run the following command:

```bash
go run main.go
```

You should see the following output:

```
Hello, World!
```

## Recommended IDEs and setup

You can write Go code in any text editor, but we recommend using an integrated development environment (IDE) for a more productive development experience. Here are a few popular IDEs for Go development:

*   **Visual Studio Code (VS Code):** A lightweight and extensible code editor with excellent support for Go. You'll need to install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) to get features like code completion, debugging, and testing.
*   **GoLand:** A full-featured IDE for Go development from JetBrains. It offers advanced features like code analysis, refactoring, and a built-in debugger.

That's it for Module 1! In the next module, we'll dive deeper into Go's syntax and data types.

**Next:** [Module 2: Basic Syntax and Data Types](../2/README.md)
