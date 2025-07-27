# Module 9: Tooling, Dependencies, and Testing

In this module, we'll cover the essential tools for managing dependencies and testing your Go code.

## Go Modules: Managing dependencies (`go.mod`, `go.sum`)

Go Modules is the official dependency management system for Go. It was introduced in Go 1.11.

To start a new project with Go Modules, you can use the `go mod init` command. This will create a `go.mod` file in your project's root directory.

```bash
go mod init example.com/myproject
```

The `go.mod` file defines the module's path and its dependencies. When you import a new package in your code, the Go tools will automatically add it to your `go.mod` file.

The `go.sum` file is a generated file that contains the checksums of the direct and indirect dependencies of your module. This file is used to verify the integrity of your dependencies.

## Writing Unit Tests with the `testing` package

The `testing` package provides support for automated testing of Go packages. To write a new test suite, create a file whose name ends in `_test.go`.

Here's an example of a simple test function:

```go
// mymath.go
package mymath

func Add(a, b int) int {
    return a + b
}
```

```go
// mymath_test.go
package mymath

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

To run the tests, you can use the `go test` command.

```bash
go test
```

## Benchmarking your code for performance analysis

The `testing` package also provides support for benchmarking. A benchmark function is similar to a test function, but it is run multiple times to measure the performance of your code.

Here's an example of a benchmark function:

```go
// mymath_test.go
package mymath

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

To run the benchmarks, you can use the `go test -bench=.` command.

```bash
go test -bench=.
```

This will run all of the benchmark functions in your package and print the results.

This module has introduced you to the essential tools for managing dependencies and testing your Go code. With these tools, you can build robust and reliable Go applications.

**Previous:** [Module 8: Data Processing and Serialization](../8/README.md)
**Next:** [Module 10: Client-Side Networking](../10/README.md)
