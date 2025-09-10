# 12. Goroutines

Welcome to the advanced section of our Go learning path! In this lesson, we'll introduce you to one of Go's most powerful features: goroutines. Goroutines are the foundation of concurrent programming in Go.

## What is a Goroutine?

A goroutine is a lightweight thread of execution. Goroutines are managed by the Go runtime, which allows you to have thousands, or even hundreds of thousands, of goroutines running concurrently in a single program.

Compared to traditional operating system threads, goroutines are much cheaper to create and manage. They start with a small stack size that can grow and shrink as needed, which makes them very memory efficient.

## Creating a Goroutine

Creating a goroutine is incredibly simple. You just need to use the `go` keyword before a function call. The function will then be executed in a new goroutine, concurrently with the calling goroutine.

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

In this example, we have a `say` function that prints a string five times. In the `main` function, we call `say("world")` in a new goroutine. The `main` goroutine then continues to execute `say("hello")`.

You will see the output of both "hello" and "world" interleaved, demonstrating that they are running concurrently.

## Goroutines and the `main` Function

It's important to note that when the `main` function finishes, the program exits, and all other goroutines are terminated.

If we were to just call `go say("world")` and nothing else in `main`, the program would likely exit before the "world" goroutine had a chance to print anything. In the example above, the `main` goroutine is kept busy by its own call to `say("hello")`, which gives the other goroutine time to run.

In real-world applications, you'll need a way to wait for your goroutines to finish. This is where channels and other synchronization primitives come in.

## Anonymous Goroutines

You can also use the `go` keyword with an anonymous function to create a goroutine.

```go
func main() {
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // ...
}
```

Goroutines are a fundamental part of Go's concurrency model. They provide a simple and powerful way to write concurrent programs. However, when you have multiple goroutines running, you often need them to communicate with each other.

In the next lesson, we will learn about channels, which are the primary way for goroutines to communicate and synchronize their execution.
