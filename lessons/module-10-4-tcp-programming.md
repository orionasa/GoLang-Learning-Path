# 🔧 Module 10.4: Low-Level TCP Programming

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master the Foundation of Network Communication**

*From Zero to Go Hero - Module 10.4*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 2.5-3 hours  | Advanced       | Modules 10.1-10.3, Basic networking concepts |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand TCP/IP fundamentals and the `net` package
- ✅ Build TCP servers that handle multiple concurrent connections
- ✅ Create robust TCP clients with proper connection management
- ✅ Implement custom protocols over TCP
- ✅ Handle binary data and message framing
- ✅ Apply concurrency patterns for scalable network applications

---

## 🌐 Understanding TCP and the net Package

<div style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 TCP: The Reliable Highway of the Internet

TCP (Transmission Control Protocol) is like having a **reliable postal service** for your data. Unlike UDP (which is like throwing postcards and hoping they arrive), TCP ensures every piece of data gets delivered in the right order, like a careful librarian organizing books on shelves!

</div>

### 🚄 **1. TCP vs HTTP Comparison**
```
┌─────────────────────────────────────────────────────────────┐
│                    NETWORK STACK LAYERS                     │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  HTTP Layer      ████████████████████████████████████ 100% │
│  (Application)   ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  TCP Layer       ████████████████████████████████████ 100% │
│  (Transport)     ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  IP Layer        ████████████████████████████████████ 100% │
│  (Network)       ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  Physical        ████████████████████████████████████ 100% │
│  (Hardware)      ████████████████████████████████████ 100% │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 🧵 **2. Key TCP Concepts**
<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**TCP Fundamentals** - Understanding the building blocks:

- **Connection-oriented** - Establish connection before data transfer
- **Reliable delivery** - Guarantees data arrives and in order
- **Flow control** - Manages data transmission rate
- **Error detection** - Checksums ensure data integrity
- **Full-duplex** - Data flows both directions simultaneously

</div>

---

## 🛠️ Building TCP Servers and Clients

### **Step 1: Basic TCP Echo Server**

```go
package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    "time"
)

type EchoServer struct {
    address string
}

func NewEchoServer(address string) *EchoServer {
    return &EchoServer{address: address}
}

func (s *EchoServer) Start() error {
    listener, err := net.Listen("tcp", s.address)
    if err != nil {
        return fmt.Errorf("failed to start server: %w", err)
    }
    defer listener.Close()
    
    fmt.Printf("🚀 Echo server started on %s\n", s.address)
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("❌ Failed to accept connection: %v\n", err)
            continue
        }
        
        fmt.Printf("🔗 New connection from %s\n", conn.RemoteAddr())
        go s.handleConnection(conn)
    }
}

func (s *EchoServer) handleConnection(conn net.Conn) {
    defer conn.Close()
    
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)
    
    writer.WriteString("🎉 Welcome to Echo Server!\n")
    writer.Flush()
    
    for {
        conn.SetReadDeadline(time.Now().Add(5 * time.Minute))
        
        message, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        
        message = strings.TrimSpace(message)
        if strings.ToLower(message) == "quit" {
            writer.WriteString("👋 Goodbye!\n")
            writer.Flush()
            break
        }
        
        response := fmt.Sprintf("Echo: %s\n", message)
        writer.WriteString(response)
        writer.Flush()
    }
}

func main() {
    server := NewEchoServer(":8080")
    server.Start()
}
```

### **Step 2: TCP Client**

```go
package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
    "time"
)

type TCPClient struct {
    address string
    conn    net.Conn
}

func NewTCPClient(address string) *TCPClient {
    return &TCPClient{address: address}
}

func (c *TCPClient) Connect() error {
    conn, err := net.DialTimeout("tcp", c.address, 10*time.Second)
    if err != nil {
        return fmt.Errorf("failed to connect: %w", err)
    }
    
    c.conn = conn
    fmt.Printf("✅ Connected to %s\n", conn.RemoteAddr())
    return nil
}

func (c *TCPClient) StartSession() error {
    defer c.conn.Close()
    
    go c.readServerMessages()
    
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("💬 Type messages (type 'quit' to exit):")
    
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        
        message := scanner.Text()
        if strings.TrimSpace(message) == "" {
            continue
        }
        
        _, err := c.conn.Write([]byte(message + "\n"))
        if err != nil {
            fmt.Printf("❌ Failed to send: %v\n", err)
            break
        }
        
        if strings.ToLower(strings.TrimSpace(message)) == "quit" {
            break
        }
    }
    
    return nil
}

func (c *TCPClient) readServerMessages() {
    reader := bufio.NewReader(c.conn)
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        fmt.Printf("📨 Server: %s", message)
    }
}

func main() {
    client := NewTCPClient("localhost:8080")
    
    if err := client.Connect(); err != nil {
        fmt.Printf("❌ Connection failed: %v\n", err)
        return
    }
    
    client.StartSession()
}
```

### **Step 3: Multi-Client Chat Server**

```go
package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "net"
    "strings"
    "sync"
    "time"
)

type Message struct {
    From      string    `json:"from"`
    Content   string    `json:"content"`
    Timestamp time.Time `json:"timestamp"`
    Type      string    `json:"type"`
}

type Client struct {
    ID       string
    Username string
    Conn     net.Conn
    Writer   *bufio.Writer
}

type ChatServer struct {
    clients    map[string]*Client
    clientsMux sync.RWMutex
    broadcast  chan Message
    register   chan *Client
    unregister chan *Client
}

func NewChatServer() *ChatServer {
    return &ChatServer{
        clients:    make(map[string]*Client),
        broadcast:  make(chan Message, 100),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (s *ChatServer) Start(address string) error {
    listener, err := net.Listen("tcp", address)
    if err != nil {
        return err
    }
    defer listener.Close()
    
    fmt.Printf("💬 Chat server started on %s\n", address)
    
    go s.messageHub()
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go s.handleConnection(conn)
    }
}

func (s *ChatServer) messageHub() {
    for {
        select {
        case client := <-s.register:
            s.clientsMux.Lock()
            s.clients[client.ID] = client
            s.clientsMux.Unlock()
            
            joinMsg := Message{
                From:      "System",
                Content:   fmt.Sprintf("%s joined", client.Username),
                Timestamp: time.Now(),
                Type:      "join",
            }
            s.broadcast <- joinMsg
            
        case client := <-s.unregister:
            s.clientsMux.Lock()
            if _, exists := s.clients[client.ID]; exists {
                delete(s.clients, client.ID)
                client.Conn.Close()
            }
            s.clientsMux.Unlock()
            
            leaveMsg := Message{
                From:      "System",
                Content:   fmt.Sprintf("%s left", client.Username),
                Timestamp: time.Now(),
                Type:      "leave",
            }
            s.broadcast <- leaveMsg
            
        case message := <-s.broadcast:
            s.broadcastMessage(message)
        }
    }
}

func (s *ChatServer) handleConnection(conn net.Conn) {
    defer conn.Close()
    
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)
    
    writer.WriteString("👤 Enter username: ")
    writer.Flush()
    
    username, err := reader.ReadString('\n')
    if err != nil {
        return
    }
    username = strings.TrimSpace(username)
    
    client := &Client{
        ID:       fmt.Sprintf("%s-%d", conn.RemoteAddr(), time.Now().Unix()),
        Username: username,
        Conn:     conn,
        Writer:   writer,
    }
    
    s.register <- client
    defer func() { s.unregister <- client }()
    
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        
        content := strings.TrimSpace(line)
        if content == "" {
            continue
        }
        
        message := Message{
            From:      client.Username,
            Content:   content,
            Timestamp: time.Now(),
            Type:      "chat",
        }
        
        s.broadcast <- message
    }
}

func (s *ChatServer) broadcastMessage(message Message) {
    s.clientsMux.RLock()
    defer s.clientsMux.RUnlock()
    
    jsonData, _ := json.Marshal(message)
    
    for _, client := range s.clients {
        client.Writer.WriteString(string(jsonData) + "\n")
        client.Writer.Flush()
    }
}

func main() {
    server := NewChatServer()
    server.Start(":8081")
}
```

---

## 🎯 Practical Project: Binary Protocol File Transfer

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 Project: Custom Binary Protocol

Build a file transfer system with a custom binary protocol that demonstrates advanced TCP programming.

</div>

```go
package main

import (
    "encoding/binary"
    "fmt"
    "io"
    "net"
    "os"
    "path/filepath"
)

// Protocol constants
const (
    ProtocolVersion = 1
    ChunkSize       = 8192
)

// Message types
const (
    MsgTypeFileRequest  = 1
    MsgTypeFileResponse = 2
    MsgTypeFileData     = 3
    MsgTypeFileEnd      = 4
    MsgTypeError        = 5
)

// Header represents protocol header
type Header struct {
    Version uint8
    MsgType uint8
    Length  uint32
}

// FileServer handles file transfer requests
type FileServer struct {
    address string
    baseDir string
}

func NewFileServer(address, baseDir string) *FileServer {
    return &FileServer{address: address, baseDir: baseDir}
}

func (s *FileServer) Start() error {
    listener, err := net.Listen("tcp", s.address)
    if err != nil {
        return err
    }
    defer listener.Close()
    
    fmt.Printf("📁 File server started on %s\n", s.address)
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go s.handleConnection(conn)
    }
}

func (s *FileServer) handleConnection(conn net.Conn) {
    defer conn.Close()
    
    for {
        header, err := s.readHeader(conn)
        if err != nil {
            break
        }
        
        switch header.MsgType {
        case MsgTypeFileRequest:
            s.handleFileRequest(conn, header)
        }
    }
}

func (s *FileServer) readHeader(conn net.Conn) (*Header, error) {
    headerBytes := make([]byte, 6)
    _, err := io.ReadFull(conn, headerBytes)
    if err != nil {
        return nil, err
    }
    
    return &Header{
        Version: headerBytes[0],
        MsgType: headerBytes[1],
        Length:  binary.BigEndian.Uint32(headerBytes[2:6]),
    }, nil
}

func (s *FileServer) writeHeader(conn net.Conn, msgType uint8, length uint32) error {
    header := make([]byte, 6)
    header[0] = ProtocolVersion
    header[1] = msgType
    binary.BigEndian.PutUint32(header[2:6], length)
    
    _, err := conn.Write(header)
    return err
}

func (s *FileServer) handleFileRequest(conn net.Conn, header *Header) {
    filenameBytes := make([]byte, header.Length)
    _, err := io.ReadFull(conn, filenameBytes)
    if err != nil {
        return
    }
    
    filename := string(filenameBytes)
    filePath := filepath.Join(s.baseDir, filename)
    
    file, err := os.Open(filePath)
    if err != nil {
        s.sendError(conn, "File not found")
        return
    }
    defer file.Close()
    
    fileInfo, _ := file.Stat()
    
    // Send file response
    response := fmt.Sprintf("%s:%d", filename, fileInfo.Size())
    s.writeHeader(conn, MsgTypeFileResponse, uint32(len(response)))
    conn.Write([]byte(response))
    
    // Send file data
    buffer := make([]byte, ChunkSize)
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            break
        }
        
        s.writeHeader(conn, MsgTypeFileData, uint32(n))
        conn.Write(buffer[:n])
    }
    
    // Send end marker
    s.writeHeader(conn, MsgTypeFileEnd, 0)
}

func (s *FileServer) sendError(conn net.Conn, message string) {
    s.writeHeader(conn, MsgTypeError, uint32(len(message)))
    conn.Write([]byte(message))
}

func main() {
    server := NewFileServer(":8082", "./files")
    server.Start()
}
```

---

## 🧪 Practice Exercises

### **Exercise 1: Protocol Design**
Create a custom protocol for:
1. Message authentication
2. Compression support
3. Heartbeat mechanism

### **Exercise 2: Connection Pooling**
Build a client that:
1. Maintains connection pools
2. Handles connection reuse
3. Implements connection health checks

### **Exercise 3: Load Balancer**
Create a TCP load balancer that:
1. Distributes connections across backends
2. Handles backend failures
3. Implements different balancing algorithms

---

## 🎯 Key Takeaways

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**🎓 What You've Learned:**

1. **TCP Fundamentals** - Connection-oriented, reliable communication
2. **Server Architecture** - Concurrent connection handling with goroutines
3. **Protocol Design** - Custom binary protocols and message framing
4. **Connection Management** - Timeouts, cleanup, and resource management
5. **Real-world Patterns** - Chat servers, file transfer, and custom protocols

</div>

---

## 🚀 Module 10 Complete!

🎉 **Congratulations!** You've mastered client-side networking in Go! You now have the skills to:

- Build HTTP clients for REST APIs
- Handle authentication and error scenarios
- Implement fault-tolerant network applications
- Create low-level TCP servers and clients
- Design custom network protocols

**What's next?** Continue your Go journey with advanced topics like web frameworks, databases, and production deployment!

---

<div align="center">

**Happy Network Programming! 🔧**

*Remember: Great network applications are built on solid foundations!*

</div>
