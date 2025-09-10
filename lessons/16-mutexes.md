# 16. Mutexes in Go

While Go's philosophy is to "share memory by communicating" (using channels), there are times when you need to use traditional locking mechanisms to protect shared data. In Go, this is done using mutexes from the `sync` package.

## What is a Mutex?

A mutex (short for "mutual exclusion") is a synchronization primitive that can be used to protect shared data from being accessed by multiple goroutines at the same time.

The `sync` package provides two types of mutexes:
*   `sync.Mutex`: A regular mutex that provides exclusive access.
*   `sync.RWMutex`: A reader/writer mutex that allows multiple readers to access the data simultaneously, but only one writer at a time.

## `sync.Mutex`

A `sync.Mutex` has two methods: `Lock()` and `Unlock()`. When a goroutine calls `Lock()`, it acquires the lock. If another goroutine has already acquired the lock, the call to `Lock()` will block until the lock is released. When the goroutine is done with the shared data, it calls `Unlock()` to release the lock.

Here's an example of using a `sync.Mutex` to protect a shared counter:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// SafeCounter is a counter that is safe to use concurrently.
type SafeCounter struct {
    mu sync.Mutex
    v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
    c.mu.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    c.v[key]++
    c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
    c.mu.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    defer c.mu.Unlock()
    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }

    // Wait for a bit for the goroutines to finish
    // (In a real program, you would use a WaitGroup)
    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}
```
In this example, the `SafeCounter` struct embeds a `sync.Mutex`. The `Inc` and `Value` methods lock the mutex before accessing the shared map `v` and unlock it when they are done. This ensures that only one goroutine can access the map at a time, preventing race conditions.

The `defer` statement is often used with `Unlock()` to ensure that the mutex is always unlocked, even if the function returns early or panics.

## `sync.RWMutex`

A `sync.RWMutex` is a reader/writer mutex. It is useful when you have a situation where there are many more reads than writes. It allows multiple goroutines to read the data at the same time, but it provides exclusive access for writing.

A `sync.RWMutex` has the following methods:
*   `Lock()` and `Unlock()`: For exclusive write access.
*   `RLock()` and `RUnlock()`: For shared read access.

Here's how you might use an `RWMutex`:

```go
type SafeConfig struct {
    mu     sync.RWMutex
    config map[string]string
}

func (c *SafeConfig) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.config[key]
}

func (c *SafeConfig) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.config[key] = value
}
```
In this example, the `Get` method uses `RLock()` to acquire a read lock, which allows multiple goroutines to call `Get` concurrently. The `Set` method uses `Lock()` to acquire a write lock, which ensures that only one goroutine can be writing to the config at a time, and that no other goroutines can be reading while the write is in progress.

Mutexes are a fundamental tool for writing correct concurrent programs in Go. They are essential when you need to protect shared state that is accessed by multiple goroutines.
