# 3. Go Basic Syntax

Now that you have your Go environment set up, let's dive into the basic syntax of the language. In this lesson, we will cover the fundamental building blocks of a Go program.

## Packages

Every Go program is made up of packages. A package is a way to group related Go source files together. The `main` package is a special package that is used to create an executable program.

A Go program starts running in the `main` package. The `main` function inside the `main` package is the entry point of the program.

Here is the basic structure of a Go program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

In this example, the `package main` line declares that this file belongs to the `main` package.

## Imports

The `import` keyword is used to import packages into your program. This allows you to use the functions and variables defined in those packages. In the example above, we imported the `fmt` package, which provides functions for formatted I/O (like printing to the console).

You can import multiple packages by enclosing them in parentheses:

```go
import (
    "fmt"
    "math"
)
```

## Functions

A function is a block of code that performs a specific task. In Go, a function is defined using the `func` keyword, followed by the function name, a list of parameters, the return type, and the function body.

The `main` function is a special function that is the entry point of every executable Go program. It takes no arguments and returns no value.

```go
func main() {
    // This is the function body
    // It contains the code that will be executed when the program runs
}
```

We will cover functions in more detail in a later lesson.

## Comments

Comments are used to explain the code and are ignored by the compiler. Go supports two types of comments:

*   **Single-line comments:** These start with `//` and continue to the end of the line.

    ```go
    // This is a single-line comment
    ```

*   **Multi-line comments:** These start with `/*` and end with `*/`.

    ```go
    /*
    This is a multi-line comment.
    It can span across multiple lines.
    */
    ```

## Semicolons

In Go, you don't need to use semicolons at the end of statements. The Go compiler automatically inserts them for you at the end of each line. This is a feature called "automatic semicolon insertion".

This helps in keeping the code clean and readable.

## A Note on Formatting

Go has a standard code formatting style that is enforced by the `gofmt` tool. This tool automatically formats your Go code according to the official style guide. This ensures that all Go code looks the same, making it easier to read and understand.

Most Go editors and IDEs are configured to run `gofmt` automatically when you save a file.

In the next lesson, we will explore variables and data types in Go.
