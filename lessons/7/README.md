# Module 7: File System I/O

In this module, we'll learn how to work with files and directories in Go. We'll cover reading from and writing to files, as well as navigating the file system.

## Reading files using the `os` and `io` packages

The `os` package provides a platform-independent interface to operating system functionality. The `io` package provides basic interfaces to I/O primitives.

To read a file, you can use the `os.ReadFile` function. This function reads the entire contents of a file into a byte slice.

```go
import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("test.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println(string(data))
}
```

For more control over reading, you can use `os.Open` to get an `*os.File` value, which implements the `io.Reader` interface.

## Writing data to files

To write data to a file, you can use the `os.WriteFile` function. This function writes a byte slice to a file.

```go
import (
    "fmt"
    "os"
)

func main() {
    data := []byte("Hello, World!")
    err := os.WriteFile("test.txt", data, 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }
}
```

The third argument to `os.WriteFile` is a file mode, which specifies the file's permissions.

## Buffering I/O for performance with `bufio`

The `bufio` package implements buffered I/O. It wraps an `io.Reader` or `io.Writer` object, creating another object (a `Reader` or `Writer`) that also implements the interface but provides buffering and some help for textual I/O.

```go
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
    }
}
```

## Navigating the file system

The `os` package also provides functions for working with directories.

*   `os.Mkdir`: Creates a new directory.
*   `os.MkdirAll`: Creates a new directory and any necessary parents.
*   `os.ReadDir`: Reads the contents of a directory.
*   `os.Remove`: Removes a file or empty directory.
*   `os.RemoveAll`: Removes a path and any children it contains.

```go
import (
    "fmt"
    "os"
)

func main() {
    err := os.Mkdir("testdir", 0755)
    if err != nil {
        fmt.Println("Error creating directory:", err)
        return
    }

    files, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }

    err = os.Remove("testdir")
    if err != nil {
        fmt.Println("Error removing directory:", err)
        return
    }
}
```

This module has covered the basics of file system I/O in Go. You can now read from and write to files, and work with directories.

**Previous:** [Module 6: Advanced Concurrency Patterns](../6/README.md)
**Next:** [Module 8: Data Processing and Serialization](../8/README.md)
