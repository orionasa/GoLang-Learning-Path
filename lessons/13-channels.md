# 13. Channels

In the previous lesson, you learned about goroutines. Now, we'll explore channels, which are the primary way for goroutines to communicate and synchronize their execution. The Go philosophy is: "Do not communicate by sharing memory; instead, share memory by communicating."

## What is a Channel?

A channel is a typed conduit through which you can send and receive values with the channel operator, `<-`.

## Creating a Channel

You can create a channel using the `make` function.

```go
// Create a channel of type int
ch := make(chan int)
```

By default, channels are unbuffered, which means they will only accept a send operation if there is a corresponding receive operation ready to receive the value. Unbuffered channels are used for synchronization.

## Sending and Receiving from a Channel

You can send a value to a channel using the `<-` operator.

```go
ch <- 1 // Send 1 to the channel ch
```

You can receive a value from a channel using the same operator.

```go
x := <-ch // Receive from ch and assign the value to x
```

Here's a simple example of using a channel to communicate between two goroutines:

```go
package main

import "fmt"

func main() {
    messages := make(chan string)

    go func() {
        messages <- "ping"
    }()

    msg := <-messages
    fmt.Println(msg) // prints "ping"
}
```

In this example, the `main` goroutine creates a channel and then launches a new goroutine. The new goroutine sends the string "ping" to the channel. The `main` goroutine then blocks until it receives the message from the channel.

## Buffered Channels

You can also create a buffered channel by providing a capacity as the second argument to `make`.

```go
ch := make(chan int, 100)
```

A buffered channel will only block a send operation when the buffer is full. It will only block a receive operation when the buffer is empty.

## Ranging Over a Channel

You can use a `for...range` loop to receive values from a channel until it is closed.

```go
func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}
```

In this example, we create a buffered channel, send two values to it, and then close it. The `for...range` loop receives and prints the two values from the channel and then terminates when the channel is closed.

Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

## `select` Statement

The `select` statement lets a goroutine wait on multiple communication operations. A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```go
func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

The `select` statement is a powerful tool for building concurrent systems in Go.

Channels are a key feature of Go's concurrency model. They provide a safe and elegant way for goroutines to communicate and synchronize.

In the next lesson, we will cover error handling in Go.
