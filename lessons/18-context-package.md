# 18. The `context` Package

In concurrent and networked Go programs, it's often necessary to manage the lifecycle of a request across multiple goroutines. The `context` package provides a powerful and standardized way to handle cancellation, deadlines, and request-scoped values.

## What is `context.Context`?

A `context.Context` is an object that carries a deadline, a cancellation signal, and other request-scoped values across API boundaries and between processes.

The `context` package is a core part of the Go standard library and is widely used in network clients and servers.

## Creating a Context

You typically create a `context.Context` at the beginning of a request's lifecycle. There are a few ways to create a context:

*   `context.Background()`: Returns a non-nil, empty context. It is never canceled, has no values, and no deadline. It is typically used by the `main` function, initialization, and tests, and as the top-level context for incoming requests.
*   `context.TODO()`: Also returns a non-nil, empty context. It should be used when you are unsure which context to use or if the function will be updated later to use a context.

## Deriving Contexts

The real power of the `context` package comes from creating derived contexts. You can create a new context from an existing one, and the new context will inherit the properties of its parent.

There are four functions for creating derived contexts:

*   `context.WithCancel()`: Returns a copy of the parent context and a `CancelFunc`. Calling the `CancelFunc` cancels the context and all contexts derived from it.
*   `context.WithDeadline()`: Returns a copy of the parent context with a deadline. The context is automatically canceled when the deadline is exceeded.
*   `context.WithTimeout()`: Similar to `WithDeadline`, but takes a timeout duration instead of a specific time.
*   `context.WithValue()`: Returns a copy of the parent context with a new key-value pair.

## Cancellation

Here's an example of using `context.WithCancel` to gracefully shut down a goroutine:

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println(name, "is stopping")
            return
        default:
            fmt.Println(name, "is working")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go worker(ctx, "worker 1")

    time.Sleep(2 * time.Second)
    cancel() // Cancel the context

    // Wait for the worker to stop
    time.Sleep(1 * time.Second)
    fmt.Println("main function finished")
}
```

In this example, the `worker` goroutine periodically checks if its context has been canceled using `ctx.Done()`. The `main` goroutine creates a cancellable context, starts the worker, waits for a couple of seconds, and then calls `cancel()` to signal the worker to stop.

## Deadlines and Timeouts

You can use `context.WithTimeout` to create a context that will be automatically canceled after a certain duration. This is very useful for operations that should not take too long, like network requests.

```go
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel() // Always call cancel, even with a timeout

    select {
    case <-time.After(2 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err()) // prints "context deadline exceeded"
    }
}
```

## Request-Scoped Values

You can use `context.WithValue` to pass request-scoped values down the call chain. This is useful for passing information like a request ID or an authenticated user.

```go
func processRequest(ctx context.Context) {
    requestID, ok := ctx.Value("requestID").(string)
    if !ok {
        requestID = "unknown"
    }
    fmt.Println("Processing request with ID:", requestID)
}

func main() {
    ctx := context.WithValue(context.Background(), "requestID", "12345")
    processRequest(ctx)
}
```

It is important to use a custom type for context keys to avoid collisions with keys used in other packages.

The `context` package is an essential tool for writing robust and scalable concurrent programs in Go. It provides a clean and consistent way to manage the lifecycle of requests and prevent resource leaks.
