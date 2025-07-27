# Module 5: Introduction to Concurrency

Concurrency is one of Go's most powerful features. In this module, we'll introduce the basic concepts of concurrency in Go, including goroutines and channels.

## Understanding Concurrency vs. Parallelism

*   **Concurrency** is about dealing with lots of things at once. It's about the composition of independently executing processes.
*   **Parallelism** is about doing lots of things at once. It's about the simultaneous execution of computations.

Go provides concurrency primitives, such as goroutines and channels, which can be used to write concurrent code that may or may not be parallel.

## Goroutines: The lightweight threads of Go

A goroutine is a lightweight thread of execution. Goroutines are managed by the Go runtime, and they are much cheaper to create than traditional threads.

To start a new goroutine, you simply use the `go` keyword followed by a function call.

```go
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

In this example, the `say("world")` call is executed in a new goroutine, while the `say("hello")` call is executed in the main goroutine.

## Channels: The primary way to communicate between Goroutines

Channels are a typed conduit through which you can send and receive values with the `<-` operator. Channels are the primary way to communicate between goroutines.

```go
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
}
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

## Unbuffered vs. Buffered Channels

### Unbuffered Channels

An unbuffered channel is a channel with a capacity of 0. When you send a value to an unbuffered channel, the send operation blocks until another goroutine is ready to receive the value.

```go
c := make(chan int) // unbuffered channel
```

### Buffered Channels

A buffered channel is a channel with a capacity greater than 0. When you send a value to a buffered channel, the send operation only blocks if the buffer is full.

```go
c := make(chan int, 10) // buffered channel with a capacity of 10
```

Buffered channels can be useful for managing the flow of data between goroutines, but they can also introduce complexity and potential deadlocks if not used carefully.

This module has provided a brief introduction to concurrency in Go. In the next module, we'll explore more advanced concurrency patterns.

**Previous:** [Module 4: Structs, Methods & Interfaces](../4/README.md)
**Next:** [Module 6: Advanced Concurrency Patterns](../6/README.md)
