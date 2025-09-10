# 27. TCP Programming in Go

TCP (Transmission Control Protocol) is a connection-oriented protocol that provides reliable, ordered, and error-checked delivery of a stream of bytes. It is the foundation of many higher-level protocols, including HTTP. In this lesson, we'll explore how to build a simple TCP server and client in Go using the `net` package.

## TCP Server

A TCP server listens for incoming connections on a specific address and port.

Here's how to create a simple echo server that reads data from a client and writes it back:

```go
package main

import (
    "io"
    "log"
    "net"
)

// handleConnection reads data from the connection and writes it back.
func handleConnection(conn net.Conn) {
    defer conn.Close()
    log.Printf("Serving %s\n", conn.RemoteAddr().String())
    if _, err := io.Copy(conn, conn); err != nil {
        log.Println("Error:", err)
    }
    log.Printf("Finished serving %s\n", conn.RemoteAddr().String())
}

func main() {
    // Listen for incoming connections.
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal("Error listening:", err)
    }
    defer listener.Close()
    log.Println("Listening on :8080")

    for {
        // Wait for a connection.
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Error accepting connection:", err)
            continue
        }

        // Handle the connection in a new goroutine.
        go handleConnection(conn)
    }
}
```

Let's break down the server code:
1.  `net.Listen("tcp", ":8080")`: This starts a TCP listener on port 8080.
2.  `listener.Accept()`: This method blocks until a new client connects and then returns a `net.Conn` object representing the connection.
3.  We handle each connection in a new goroutine using `go handleConnection(conn)`. This allows the server to handle multiple clients concurrently.
4.  The `handleConnection` function uses `io.Copy(conn, conn)` to read data from the connection and write it back to the same connection. This is a simple way to create an echo server.

## TCP Client

A TCP client connects to a TCP server to send and receive data.

Here's a simple client that connects to our echo server, sends a message, and reads the response:

```go
package main

import (
    "io"
    "log"
    "net"
    "os"
)

func main() {
    // Connect to the server.
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Error connecting:", err)
    }
    defer conn.Close()
    log.Println("Connected to server")

    // Send a message to the server.
    if _, err := conn.Write([]byte("Hello, Server!")); err != nil {
        log.Fatal("Error writing:", err)
    }

    // Read the response from the server.
    if _, err := io.Copy(os.Stdout, conn); err != nil {
        log.Fatal("Error reading:", err)
    }
}
```
In this client:
1.  `net.Dial("tcp", "localhost:8080")` connects to the TCP server at `localhost:8080`.
2.  `conn.Write()` sends a message to the server.
3.  `io.Copy(os.Stdout, conn)` reads the response from the server and writes it to standard output.

TCP is a powerful protocol for building reliable networked applications. The `net` package in Go provides a simple yet powerful interface for working with TCP, making it easy to build robust clients and servers.

In the next lesson, we will look at UDP programming, which is another important protocol in the TCP/IP suite.
