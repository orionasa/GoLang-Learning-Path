# 7. Packages in Go

In Go, packages are used to organize and reuse code. A package is a collection of source files in the same directory that are compiled together. You've already used packages like `fmt` and `math` from the standard library. In this lesson, you'll learn how to create your own packages.

## Creating a Package

To create a package, you simply create a new directory and place your Go source files inside it. All files in the directory must have the same package name, declared at the top of the file with the `package` keyword.

Let's create a simple `calculator` package.

1.  Create a new directory named `calculator` inside your project.
2.  Inside the `calculator` directory, create a file named `add.go` with the following content:

    ```go
    package calculator

    // Add returns the sum of two integers.
    // Note that the function name starts with a capital letter.
    func Add(a, b int) int {
        return a + b
    }
    ```

3.  Create another file named `subtract.go` in the same directory:

    ```go
    package calculator

    // Subtract returns the difference between two integers.
    func Subtract(a, b int) int {
        return a - b
    }
    ```

We now have a `calculator` package with two functions.

## Exported Names

In Go, if a name (like a function or variable) starts with a capital letter, it is considered "exported". Exported names are accessible from other packages. Names that start with a lowercase letter are not exported and are only accessible within the same package.

In our `calculator` package, the `Add` and `Subtract` functions are exported because they start with a capital letter.

## Using a Package

Now, let's use our `calculator` package in our `main` program.

Assuming your project has the following structure:

```
my-project/
├── go.mod
├── main.go
└── calculator/
    ├── add.go
    └── subtract.go
```

Your `main.go` file would look like this:

```go
package main

import (
    "fmt"
    "my-project/calculator" // Import the local package
)

func main() {
    sum := calculator.Add(5, 3)
    fmt.Println("Sum:", sum) // Output: Sum: 8

    difference := calculator.Subtract(5, 3)
    fmt.Println("Difference:", difference) // Output: Difference: 2
}
```

To import a local package, you use the module path followed by the path to the package directory. If your `go.mod` file defines the module as `my-project`, then the import path for the `calculator` package is `my-project/calculator`.

## Go Modules

Go modules are used to manage dependencies in your project. A module is a collection of related Go packages that are versioned together as a single unit.

You initialize a module by running the `go mod init <module-path>` command in the root of your project. This creates a `go.mod` file. The module path is typically the URL of your repository, for example, `github.com/your-username/my-project`.

When you import a package, Go uses the module's `go.mod` file to find and download the required dependencies.

## The Standard Library

Go comes with a rich standard library that is organized into packages. These packages provide a wide range of functionality, from basic I/O (`fmt`, `io`) to networking (`net/http`) and cryptography (`crypto/aes`).

You can explore the standard library packages on the [official Go documentation website](https://pkg.go.dev/std).

In the next lesson, we will look at arrays and slices, which are used to store collections of data.
