# 🔒 Lesson 6.3: Mutex & Data Protection

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Protect Shared Data with Mutexes**

*Module 6 - Advanced Concurrency Patterns*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 120 minutes  | Intermediate   | Goroutines, WaitGroup |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand race conditions and why they're dangerous
- ✅ Use `sync.Mutex` to protect shared data
- ✅ Master `sync.RWMutex` for read-heavy workloads
- ✅ Build thread-safe data structures
- ✅ Recognize when to use mutexes vs channels

---

## Understanding Race Conditions

<div style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 The Data Protection Problem

A **race condition** occurs when multiple goroutines access shared data simultaneously, and at least one of them modifies it. This can lead to unpredictable and incorrect results.

**Think of it like:** Multiple people trying to edit the same document at the same time without coordination - chaos ensues!

</div>

### 🚄 **1. Race Condition Example**

```
┌─────────────────────────────────────┐
│        Race Condition Scenario      │
├─────────────────────────────────────┤
│  Goroutine 1: Read  ████████████ 50  │
│  Goroutine 2: Read  ████████████ 50  │
│  Goroutine 1: +1    ████████████ 51  │
│  Goroutine 2: +1    ████████████ 51  │
│  Expected: 52, Got: 51 ❌           │
└─────────────────────────────────────┘
```

<div style="background: #f8f9fa; border-left: 4px solid #dc3545; padding: 15px; margin: 10px 0;">

**⚠️ Dangerous Code - Race Condition Example**

```go
package main

import (
    "fmt"
    "sync"
)

// Unsafe counter - DON'T USE THIS!
type UnsafeCounter struct {
    value int
}

func (c *UnsafeCounter) Increment() {
    c.value++ // NOT thread-safe!
}

func (c *UnsafeCounter) Value() int {
    return c.value
}

func main() {
    counter := &UnsafeCounter{}
    var wg sync.WaitGroup
    
    // 1000 goroutines incrementing 1000 times each
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                counter.Increment()
            }
        }()
    }
    
    wg.Wait()
    
    expected := 1000000
    actual := counter.Value()
    
    fmt.Printf("Expected: %d\n", expected)
    fmt.Printf("Actual:   %d\n", actual)
    
    if expected != actual {
        fmt.Printf("❌ Race condition! Lost %d increments\n", expected-actual)
    }
}
```

</div>

### 🧵 **2. Introducing Mutex**

`sync.Mutex` ensures only one goroutine can access protected data at a time:

```go
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()         // 🔒 Lock
    defer c.mu.Unlock() // 🔓 Unlock
    c.value++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

### 🎯 **3. Mutex vs RWMutex vs Channels**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🔒 Mutex</strong><br>
Exclusive access<br>
Simple locking<br>
Read/Write same cost
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📖 RWMutex</strong><br>
Multiple readers<br>
Single writer<br>
Read-heavy workloads
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📡 Channels</strong><br>
Data communication<br>
Goroutine coordination<br>
"Don't share memory"
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🎯 When to Use</strong><br>
Protecting state<br>
Simple synchronization<br>
Performance critical
</div>

</div>

---

## Building Thread-Safe Data Structures

### **Step 1: Thread-Safe Counter**

Create `safe_counter.go`:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// SafeCounter is a thread-safe counter
type SafeCounter struct {
    mu    sync.Mutex
    value int64
}

func NewSafeCounter() *SafeCounter {
    return &SafeCounter{}
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *SafeCounter) Add(delta int64) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value += delta
}

func (c *SafeCounter) Value() int64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := NewSafeCounter()
    var wg sync.WaitGroup
    
    start := time.Now()
    
    // 100 goroutines, 1000 operations each
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                counter.Increment()
            }
        }()
    }
    
    wg.Wait()
    duration := time.Since(start)
    
    fmt.Printf("Expected: 100000\n")
    fmt.Printf("Actual:   %d\n", counter.Value())
    fmt.Printf("Duration: %v\n", duration)
    fmt.Println("✅ Thread-safe operations successful!")
}
```

### **Step 2: Thread-Safe Map with RWMutex**

Create `safe_map.go`:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// SafeMap is a thread-safe map
type SafeMap struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

func NewSafeMap() *SafeMap {
    return &SafeMap{
        data: make(map[string]interface{}),
    }
}

func (sm *SafeMap) Set(key string, value interface{}) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (interface{}, bool) {
    sm.mu.RLock() // Read lock allows multiple readers
    defer sm.mu.RUnlock()
    value, exists := sm.data[key]
    return value, exists
}

func (sm *SafeMap) Delete(key string) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    delete(sm.data, key)
}

func (sm *SafeMap) Len() int {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    return len(sm.data)
}

func main() {
    safeMap := NewSafeMap()
    var wg sync.WaitGroup
    
    // Test concurrent writes
    fmt.Println("📝 Testing concurrent writes...")
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                key := fmt.Sprintf("key_%d_%d", id, j)
                safeMap.Set(key, fmt.Sprintf("value_%d_%d", id, j))
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("✅ Wrote %d items\n", safeMap.Len())
    
    // Test concurrent reads
    fmt.Println("\n📖 Testing concurrent reads...")
    start := time.Now()
    
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 500; j++ {
                key := fmt.Sprintf("key_%d_%d", j%50, j%100)
                safeMap.Get(key)
            }
        }(i)
    }
    
    wg.Wait()
    duration := time.Since(start)
    
    fmt.Printf("✅ Completed 50,000 reads in %v\n", duration)
}
```

### **Step 3: Advanced Cache with TTL**

Create `thread_safe_cache.go`:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
}

func (item *CacheItem) IsExpired() bool {
    return time.Now().After(item.ExpiresAt)
}

type ThreadSafeCache struct {
    mu    sync.RWMutex
    items map[string]*CacheItem
    hits  int64
    misses int64
}

func NewThreadSafeCache() *ThreadSafeCache {
    cache := &ThreadSafeCache{
        items: make(map[string]*CacheItem),
    }
    go cache.cleanup()
    return cache
}

func (c *ThreadSafeCache) Set(key string, value interface{}, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    c.items[key] = &CacheItem{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
    }
}

func (c *ThreadSafeCache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, exists := c.items[key]
    if !exists || item.IsExpired() {
        c.misses++
        return nil, false
    }
    
    c.hits++
    return item.Value, true
}

func (c *ThreadSafeCache) Stats() (hits, misses int64, hitRatio float64) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    total := c.hits + c.misses
    if total == 0 {
        return c.hits, c.misses, 0
    }
    
    return c.hits, c.misses, float64(c.hits) / float64(total)
}

func (c *ThreadSafeCache) cleanup() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        c.mu.Lock()
        expired := 0
        for key, item := range c.items {
            if item.IsExpired() {
                delete(c.items, key)
                expired++
            }
        }
        if expired > 0 {
            fmt.Printf("🧹 Cleaned up %d expired items\n", expired)
        }
        c.mu.Unlock()
    }
}

func main() {
    cache := NewThreadSafeCache()
    var wg sync.WaitGroup
    
    // Test concurrent access
    fmt.Println("💾 Testing thread-safe cache...")
    
    // Writers
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                key := fmt.Sprintf("key_%d_%d", id, j)
                value := fmt.Sprintf("value_%d_%d", id, j)
                ttl := time.Duration(1+j%5) * time.Second
                cache.Set(key, value, ttl)
            }
        }(i)
    }
    
    // Readers
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 200; j++ {
                key := fmt.Sprintf("key_%d_%d", j%10, j%100)
                cache.Get(key)
            }
        }(i)
    }
    
    wg.Wait()
    
    hits, misses, hitRatio := cache.Stats()
    fmt.Printf("📊 Cache Stats:\n")
    fmt.Printf("   Hits: %d, Misses: %d\n", hits, misses)
    fmt.Printf("   Hit Ratio: %.2f%%\n", hitRatio*100)
    
    fmt.Println("✅ Cache test completed!")
}
```

---

## 🎯 Practice Exercises

<div style="background: #e7f3ff; border-left: 4px solid #2196f3; padding: 15px; margin: 10px 0;">

### Exercise 1: Thread-Safe Bank Account
Create a thread-safe bank account that supports deposits, withdrawals, and balance checks with proper mutex protection.

### Exercise 2: Concurrent Hit Counter
Build a web hit counter that can handle thousands of concurrent requests safely.

### Exercise 3: Thread-Safe LRU Cache
Implement a Least Recently Used cache with mutex protection and automatic eviction.

</div>

---

## ⚠️ Best Practices & Pitfalls

<div style="background: #fff3cd; border-left: 4px solid #ffc107; padding: 15px; margin: 10px 0;">

### Best Practices:
1. **Always use `defer` for Unlock()** - Ensures unlock even if panic occurs
2. **Keep critical sections small** - Minimize time holding locks
3. **Use RWMutex for read-heavy workloads** - Better performance
4. **Avoid nested locks** - Can cause deadlocks

### Common Pitfalls:
1. **Forgetting to unlock** - Causes deadlocks
2. **Copying mutexes** - Pass by pointer, not value
3. **Locking in wrong order** - Can cause deadlocks
4. **Holding locks too long** - Reduces concurrency

</div>

---

## 🔑 Key Takeaways

- **Mutexes protect shared data** from race conditions
- **RWMutex optimizes** read-heavy workloads
- **Always use defer** for unlock operations
- **Keep critical sections minimal** for better performance
- **Choose the right tool**: channels for communication, mutexes for protection

---

## 🚀 Module 6 Complete!

🎉 **Congratulations!** You've mastered all of Module 6's advanced concurrency patterns:

- ✅ **Select Statement** - Channel multiplexing and timeouts
- ✅ **WaitGroup** - Goroutine coordination and worker pools  
- ✅ **Mutex** - Data protection and thread-safe structures

You're now equipped to build robust, concurrent Go applications! 

**Next up:** [Module 7 - File System I/O](../module-7-file-system-io.md) where you'll learn to work with files and directories.

---

<div align="center">

**🎉 You're becoming a Go concurrency expert! 🎉**

*Ready for the next challenge?*

</div>
