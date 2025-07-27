# 🎛️ Lesson 6.1: The Select Statement

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master Channel Multiplexing with Select**

*Module 6 - Advanced Concurrency Patterns*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 90 minutes   | Intermediate   | Basic Channels, Goroutines |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand what the `select` statement does and why it's powerful
- ✅ Handle multiple channels simultaneously without blocking
- ✅ Implement timeouts and non-blocking channel operations
- ✅ Build a real-world multi-channel server application
- ✅ Master select statement patterns and best practices

---

## What is the Select Statement?

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 The Channel Multiplexer

The `select` statement is like a **traffic controller** for channels. It lets your program wait on multiple channel operations simultaneously and proceeds with whichever one is ready first.

**Think of it like:** A restaurant server who can take orders from multiple tables - they serve whoever is ready first!

</div>

### 🚄 **1. Basic Select Syntax**

```
┌─────────────────────────────────────┐
│         Select Statement Flow        │
├─────────────────────────────────────┤
│  Channel A  ████████████████████ Ready │
│  Channel B  ███████████████████▌ Wait  │
│  Channel C  ████▌               Wait   │
│  Default    ████████████████████ Exec  │
└─────────────────────────────────────┘
```

The basic syntax looks similar to a `switch` statement:

```go
select {
case msg1 := <-channel1:
    // Handle message from channel1
case msg2 := <-channel2:
    // Handle message from channel2
case channel3 <- value:
    // Send value to channel3
default:
    // Execute if no channels are ready
}
```

### 🧵 **2. Your First Select Example**

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Multi-Channel Message Handler** - Let's build a simple message router

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create channels for different message types
    urgent := make(chan string)
    normal := make(chan string)
    
    // Start goroutines to send messages
    go func() {
        time.Sleep(1 * time.Second)
        urgent <- "🚨 URGENT: System alert!"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        normal <- "📧 Normal: Daily report ready"
    }()
    
    // Handle messages as they arrive
    for i := 0; i < 2; i++ {
        select {
        case msg := <-urgent:
            fmt.Println("Priority handling:", msg)
        case msg := <-normal:
            fmt.Println("Standard handling:", msg)
        }
    }
}
```

**Output:**
```
Priority handling: 🚨 URGENT: System alert!
Standard handling: 📧 Normal: Daily report ready
```

</div>

### 🎯 **3. Select Statement Patterns**

- **Non-blocking Operations** - Use `default` case to avoid blocking
- **Timeouts** - Combine with `time.After()` for timeout handling
- **Channel Multiplexing** - Handle multiple input sources
- **Load Balancing** - Distribute work across multiple workers

### 🏢 **4. Select vs Switch Comparison**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🔄 Switch Statement</strong><br>
Evaluates expressions<br>
Synchronous execution<br>
Single value comparison
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🎛️ Select Statement</strong><br>
Waits on channel operations<br>
Concurrent execution<br>
Multiple channel readiness
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>⚡ Performance</strong><br>
Immediate evaluation<br>
No goroutine overhead<br>
CPU-bound operations
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🌐 Concurrency</strong><br>
Blocks until ready<br>
Goroutine coordination<br>
I/O-bound operations
</div>

</div>

---

## Non-Blocking Operations: Step-by-Step Guide

### **Step 1: Understanding Blocking vs Non-Blocking**

Without `select`, channel operations can block indefinitely:

```go
// ❌ This will block forever if no one sends to ch
msg := <-ch

// ❌ This will block if channel buffer is full
ch <- "message"
```

With `select` and `default`, you can make operations non-blocking:

```go
// ✅ Non-blocking receive
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No message available")
}

// ✅ Non-blocking send
select {
case ch <- "message":
    fmt.Println("Message sent")
default:
    fmt.Println("Channel full, message dropped")
}
```

### **Step 2: Building a Non-Blocking Message Queue**

Create `message_queue.go`:

```go
package main

import (
    "fmt"
    "time"
)

type MessageQueue struct {
    messages chan string
    capacity int
}

func NewMessageQueue(capacity int) *MessageQueue {
    return &MessageQueue{
        messages: make(chan string, capacity),
        capacity: capacity,
    }
}

// Non-blocking send
func (mq *MessageQueue) TrySend(message string) bool {
    select {
    case mq.messages <- message:
        return true // Message sent successfully
    default:
        return false // Queue is full
    }
}

// Non-blocking receive
func (mq *MessageQueue) TryReceive() (string, bool) {
    select {
    case msg := <-mq.messages:
        return msg, true // Message received
    default:
        return "", false // No messages available
    }
}

func (mq *MessageQueue) Status() {
    fmt.Printf("Queue status: %d/%d messages\n", 
        len(mq.messages), mq.capacity)
}

func main() {
    // Create a small queue for demonstration
    queue := NewMessageQueue(2)
    
    // Try to send messages
    fmt.Println("=== Sending Messages ===")
    messages := []string{"Hello", "World", "Go", "Rocks"}
    
    for _, msg := range messages {
        if queue.TrySend(msg) {
            fmt.Printf("✅ Sent: %s\n", msg)
        } else {
            fmt.Printf("❌ Failed to send: %s (queue full)\n", msg)
        }
        queue.Status()
    }
    
    fmt.Println("\n=== Receiving Messages ===")
    // Try to receive all messages
    for i := 0; i < 5; i++ {
        if msg, ok := queue.TryReceive(); ok {
            fmt.Printf("✅ Received: %s\n", msg)
        } else {
            fmt.Println("❌ No messages available")
        }
        queue.Status()
    }
}
```

**Output:**
```
=== Sending Messages ===
✅ Sent: Hello
Queue status: 1/2 messages
✅ Sent: World
Queue status: 2/2 messages
❌ Failed to send: Go (queue full)
Queue status: 2/2 messages
❌ Failed to send: Rocks (queue full)
Queue status: 2/2 messages

=== Receiving Messages ===
✅ Received: Hello
Queue status: 1/2 messages
✅ Received: World
Queue status: 0/2 messages
❌ No messages available
Queue status: 0/2 messages
❌ No messages available
Queue status: 0/2 messages
❌ No messages available
Queue status: 0/2 messages
```

---

## Timeouts and Deadlines

### **Step 3: Implementing Timeouts**

<div style="background: #fff3cd; border-left: 4px solid #ffc107; padding: 15px; margin: 10px 0;">

**⏰ Why Timeouts Matter:** In real applications, you can't wait forever for a response. Timeouts prevent your program from hanging indefinitely.

</div>

```go
package main

import (
    "fmt"
    "time"
)

func fetchData(delay time.Duration) <-chan string {
    ch := make(chan string)
    go func() {
        time.Sleep(delay)
        ch <- "Data fetched successfully!"
    }()
    return ch
}

func main() {
    fmt.Println("=== Timeout Examples ===")
    
    // Example 1: Fast operation (succeeds)
    fmt.Println("\n1. Fast operation:")
    select {
    case data := <-fetchData(500 * time.Millisecond):
        fmt.Println("✅", data)
    case <-time.After(1 * time.Second):
        fmt.Println("❌ Timeout: Operation took too long")
    }
    
    // Example 2: Slow operation (times out)
    fmt.Println("\n2. Slow operation:")
    select {
    case data := <-fetchData(2 * time.Second):
        fmt.Println("✅", data)
    case <-time.After(1 * time.Second):
        fmt.Println("❌ Timeout: Operation took too long")
    }
    
    // Example 3: Multiple timeouts
    fmt.Println("\n3. Multiple operations with different timeouts:")
    
    fastCh := fetchData(300 * time.Millisecond)
    slowCh := fetchData(1500 * time.Millisecond)
    
    for i := 0; i < 2; i++ {
        select {
        case data := <-fastCh:
            fmt.Println("✅ Fast:", data)
        case data := <-slowCh:
            fmt.Println("✅ Slow:", data)
        case <-time.After(1 * time.Second):
            fmt.Println("❌ Timeout reached")
        }
    }
}
```

---

## Real-World Project: Multi-Channel Server

### **Step 4: Building a Concurrent Server**

Let's build a server that handles multiple types of requests simultaneously:

Create `multi_server.go`:

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Request types
type Request struct {
    ID       int
    Type     string
    Data     string
    Response chan string
}

// Server handles multiple request types
type Server struct {
    httpRequests  chan Request
    dbRequests    chan Request
    cacheRequests chan Request
    shutdown      chan bool
}

func NewServer() *Server {
    return &Server{
        httpRequests:  make(chan Request, 10),
        dbRequests:    make(chan Request, 5),
        cacheRequests: make(chan Request, 20),
        shutdown:      make(chan bool),
    }
}

func (s *Server) Start() {
    fmt.Println("🚀 Server starting...")
    
    go s.handleRequests()
    
    // Simulate incoming requests
    go s.simulateTraffic()
}

func (s *Server) handleRequests() {
    requestCount := 0
    
    for {
        select {
        case req := <-s.httpRequests:
            requestCount++
            go s.processHTTPRequest(req, requestCount)
            
        case req := <-s.dbRequests:
            requestCount++
            go s.processDBRequest(req, requestCount)
            
        case req := <-s.cacheRequests:
            requestCount++
            go s.processCacheRequest(req, requestCount)
            
        case <-s.shutdown:
            fmt.Println("🛑 Server shutting down...")
            return
            
        case <-time.After(5 * time.Second):
            fmt.Printf("📊 Server idle... Processed %d requests so far\n", requestCount)
        }
    }
}

func (s *Server) processHTTPRequest(req Request, count int) {
    // Simulate HTTP processing time
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    response := fmt.Sprintf("HTTP Response #%d for %s", count, req.Data)
    req.Response <- response
}

func (s *Server) processDBRequest(req Request, count int) {
    // Simulate database query time
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    response := fmt.Sprintf("DB Result #%d for %s", count, req.Data)
    req.Response <- response
}

func (s *Server) processCacheRequest(req Request, count int) {
    // Simulate fast cache lookup
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    response := fmt.Sprintf("Cache Hit #%d for %s", count, req.Data)
    req.Response <- response
}

func (s *Server) simulateTraffic() {
    requests := []struct {
        channel chan Request
        reqType string
        data    string
    }{
        {s.httpRequests, "HTTP", "GET /users"},
        {s.cacheRequests, "CACHE", "user:123"},
        {s.dbRequests, "DB", "SELECT * FROM users"},
        {s.httpRequests, "HTTP", "POST /orders"},
        {s.cacheRequests, "CACHE", "product:456"},
        {s.dbRequests, "DB", "UPDATE inventory"},
        {s.httpRequests, "HTTP", "GET /products"},
        {s.cacheRequests, "CACHE", "session:789"},
    }
    
    for i, reqData := range requests {
        time.Sleep(time.Duration(rand.Intn(800)) * time.Millisecond)
        
        responseCh := make(chan string)
        req := Request{
            ID:       i + 1,
            Type:     reqData.reqType,
            Data:     reqData.data,
            Response: responseCh,
        }
        
        // Send request to appropriate channel
        select {
        case reqData.channel <- req:
            fmt.Printf("📤 Sent %s request: %s\n", req.Type, req.Data)
            
            // Wait for response with timeout
            select {
            case response := <-responseCh:
                fmt.Printf("📥 %s\n", response)
            case <-time.After(2 * time.Second):
                fmt.Printf("⏰ Timeout waiting for %s response\n", req.Type)
            }
            
        default:
            fmt.Printf("❌ %s queue full, dropping request: %s\n", req.Type, req.Data)
        }
        
        fmt.Println()
    }
    
    // Shutdown server after processing
    time.Sleep(2 * time.Second)
    s.shutdown <- true
}

func main() {
    server := NewServer()
    server.Start()
    
    // Keep main goroutine alive
    time.Sleep(15 * time.Second)
    fmt.Println("🏁 Demo completed!")
}
```

---

## 🎯 Practice Exercises

<div style="background: #e7f3ff; border-left: 4px solid #2196f3; padding: 15px; margin: 10px 0;">

### Exercise 1: Channel Race
Create a program where multiple goroutines send messages to different channels, and use `select` to determine which message arrives first.

### Exercise 2: Timeout Handler
Build a function that tries to read from a channel with a configurable timeout. Return either the value or a timeout error.

### Exercise 3: Load Balancer
Create a simple load balancer that distributes work across multiple worker channels using `select`.

</div>

---

## 🔑 Key Takeaways

- **`select` is for channels** what `switch` is for values
- **Non-blocking operations** prevent your program from hanging
- **Timeouts** are essential for robust concurrent programs
- **Default cases** make channel operations non-blocking
- **Multiple channels** can be handled simultaneously

---

## 🚀 What's Next?

Great job mastering the `select` statement! You now know how to:
- Handle multiple channels simultaneously
- Implement timeouts and non-blocking operations
- Build concurrent servers with proper channel management

**Next up:** [Lesson 6.2 - WaitGroup Coordination](./module-6-2-waitgroup.md) where you'll learn to coordinate multiple goroutines and wait for their completion.

---

<div align="center">

**🎉 Congratulations on completing Lesson 6.1! 🎉**

*You're becoming a Go concurrency expert!*

</div>
