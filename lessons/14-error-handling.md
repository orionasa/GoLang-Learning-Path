# 14. Error Handling in Go

Error handling is a critical part of writing robust and reliable software. In Go, error handling is an explicit and deliberate process. Go does not have exceptions like other languages such as Java or Python. Instead, functions that can fail return an `error` value as their last return value.

## The `error` Type

The `error` type is a built-in interface in Go. An `error` is anything that implements the `Error()` method, which returns an error message as a string.

```go
type error interface {
    Error() string
}
```

## Returning and Handling Errors

By convention, functions that can fail return an `error` as their last return value. If the function succeeds, the `error` value is `nil`. If it fails, the `error` value will contain information about the error that occurred.

You've already seen this pattern in the lesson on functions:

```go
import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

In this example, the `divide` function returns an error if the divisor is zero. The `main` function checks if the `err` value is `nil`. If it's not `nil`, it prints the error and returns. This `if err != nil` pattern is very common in Go code.

## Creating Custom Error Types

You can create your own custom error types by defining a struct that implements the `error` interface. This is useful when you want to provide more context about an error.

```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("code %d: %s", e.Code, e.Message)
}

func doSomething() error {
    return &MyError{Code: 500, Message: "something went wrong"}
}
```

You can then use a type assertion to check for your custom error type and handle it accordingly.

```go
err := doSomething()
if err != nil {
    var myErr *MyError
    if errors.As(err, &myErr) {
        // Handle the custom error
    } else {
        // Handle other errors
    }
}
```

The `errors.As` function is the recommended way to check if an error is of a specific type.

## Wrapping Errors

Sometimes, you want to add more context to an error without losing the original error information. You can do this by "wrapping" the error. Go 1.13 introduced a new way to wrap errors using the `%w` verb in `fmt.Errorf`.

```go
func readFile() error {
    err := // ... some error from reading a file
    if err != nil {
        return fmt.Errorf("could not read file: %w", err)
    }
    return nil
}
```

You can then use `errors.Is` to check if an error in a chain of wrapped errors matches a specific error, and `errors.As` to find the first error in the chain that matches a specific type.

## `panic` and `recover`

Go has a `panic` function that stops the ordinary flow of control and begins panicking. When a function `F` calls `panic`, the execution of `F` stops, any deferred functions in `F` are executed, and then `F` returns to its caller. To the caller, `F` then behaves like a call to `panic`. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes.

`panic` is typically used for unrecoverable errors, such as a logical error in the program (e.g., an index out of bounds).

The `recover` function is a built-in function that regains control of a panicking goroutine. `recover` is only useful inside deferred functions. During normal execution, a call to `recover` will return `nil` and have no other effect. If the current goroutine is panicking, a call to `recover` will capture the value given to `panic` and resume normal execution.

While `panic` and `recover` are available, they should be used sparingly. Idiomatic Go code uses `error` values for error handling.

This concludes our advanced-level lessons on concurrency and error handling. With this knowledge, you are now well-equipped to write robust and concurrent Go programs.
