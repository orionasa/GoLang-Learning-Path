# 21. File Operations in Go

Working with files is a common task in many applications. Go provides a comprehensive set of tools for file operations in its `os` and `io` packages. In this lesson, we'll cover the basics of reading from and writing to files.

## Reading Files

There are several ways to read a file in Go.

### Reading an Entire File

The easiest way to read an entire file into memory is to use the `os.ReadFile()` function.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("test.txt")
    if err != nil {
        // handle error
    }
    fmt.Println(string(data))
}
```

This function returns the contents of the file as a byte slice.

### More Control with `os.Open`

For more control over reading a file, you can use `os.Open()`. This function returns an `*os.File` object, which you can then use to read the file in chunks.

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        // handle error
    }
    defer file.Close()

    buffer := make([]byte, 1024)
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break // End of file
        }
        if err != nil {
            // handle error
        }
        fmt.Print(string(buffer[:n]))
    }
}
```
In this example, we read the file in chunks of 1024 bytes until we reach the end of the file (`io.EOF`).

## Writing Files

### Writing a Byte Slice to a File

The `os.WriteFile()` function is the easiest way to write a byte slice to a file.

```go
package main

import (
    "os"
)

func main() {
    data := []byte("Hello, World!")
    err := os.WriteFile("output.txt", data, 0644)
    if err != nil {
        // handle error
    }
}
```
The third argument to `os.WriteFile()` is the file permissions. `0644` is a common permission setting that allows the owner to read and write the file, and everyone else to only read it.

### More Control with `os.Create`

For more control, you can use `os.Create()` to create a new file for writing. This returns an `*os.File` object that you can write to.

```go
package main

import (
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        // handle error
    }
    defer file.Close()

    _, err = file.WriteString("Hello, Go!")
    if err != nil {
        // handle error
    }
}
```

## Working with Directories

The `os` package also provides functions for working with directories.

*   `os.Mkdir()`: Creates a new directory.
*   `os.MkdirAll()`: Creates a directory and any necessary parents.
*   `os.ReadDir()`: Reads the contents of a directory.
*   `os.Remove()`: Removes a file or an empty directory.
*   `os.RemoveAll()`: Removes a path and any children it contains.

File operations are a fundamental part of many Go programs. The `os` and `io` packages provide a powerful and easy-to-use set of tools for working with the file system.

In the next lesson, we'll build on what we've learned about networking and file operations to make API calls from a Go program.
