# 28. UDP Programming in Go

UDP (User Datagram Protocol) is a connectionless protocol that, unlike TCP, does not guarantee delivery, ordering, or error checking. It is useful for applications where speed is more important than reliability, such as video streaming, online gaming, and DNS.

In this lesson, we'll explore how to build a simple UDP server and client in Go using the `net` package.

## UDP Server

A UDP server listens for incoming datagrams on a specific port. Since UDP is connectionless, the server does not establish a persistent connection with a client. It simply receives datagrams from any client that sends them to its address.

Here's a simple UDP echo server:

```go
package main

import (
    "log"
    "net"
)

func main() {
    // Listen for incoming UDP packets.
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        log.Fatal("Error resolving UDP address:", err)
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        log.Fatal("Error listening:", err)
    }
    defer conn.Close()
    log.Println("Listening on :8080")

    buffer := make([]byte, 1024)

    for {
        // Read from the UDP connection.
        n, remoteAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            log.Println("Error reading from UDP:", err)
            continue
        }
        log.Printf("Received %d bytes from %s: %s\n", n, remoteAddr, string(buffer[:n]))

        // Write the same data back to the client.
        _, err = conn.WriteToUDP(buffer[:n], remoteAddr)
        if err != nil {
            log.Println("Error writing to UDP:", err)
        }
    }
}
```

Let's break down the server code:
1.  `net.ResolveUDPAddr()` is used to get a UDP address.
2.  `net.ListenUDP()` creates a UDP connection that listens on the specified address.
3.  `conn.ReadFromUDP()` reads a datagram from the connection. It returns the number of bytes read, the address of the client that sent the datagram, and an error.
4.  `conn.WriteToUDP()` sends a datagram to a specific client address.

## UDP Client

A UDP client sends datagrams to a server.

Here's a simple client that sends a message to our UDP echo server and waits for the response:

```go
package main

import (
    "log"
    "net"
)

func main() {
    // Resolve the server address.
    serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        log.Fatal("Error resolving server address:", err)
    }

    // Create a UDP connection.
    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        log.Fatal("Error dialing:", err)
    }
    defer conn.Close()
    log.Println("Connected to server")

    // Send a message to the server.
    _, err = conn.Write([]byte("Hello, UDP Server!"))
    if err != nil {
        log.Fatal("Error writing:", err)
    }

    // Read the response from the server.
    buffer := make([]byte, 1024)
    n, _, err := conn.ReadFromUDP(buffer)
    if err != nil {
        log.Fatal("Error reading:", err)
    }
    log.Printf("Received response: %s\n", string(buffer[:n]))
}
```
The client code is similar to the TCP client, but it uses UDP-specific functions from the `net` package.

## TCP vs. UDP

| Feature        | TCP                               | UDP                                |
|----------------|-----------------------------------|------------------------------------|
| **Connection** | Connection-oriented               | Connectionless                     |
| **Reliability**| Reliable (guaranteed delivery)    | Unreliable (no guaranteed delivery)|
| **Ordering**   | Ordered delivery                  | Unordered delivery                 |
| **Speed**      | Slower (due to overhead)          | Faster (less overhead)             |
| **Use Cases**  | HTTP, FTP, email                  | Video streaming, online gaming, DNS|

Choosing between TCP and UDP depends on the requirements of your application. If you need reliable, ordered delivery of data, TCP is the right choice. If you need low-latency, high-speed communication and can tolerate some packet loss, UDP is a better option.

In the next lesson, we will look at Protocol Buffers, a data serialization format that is often used in high-performance networked applications.
