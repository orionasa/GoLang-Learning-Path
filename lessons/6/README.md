# Module 6: Advanced Concurrency Patterns

In this module, we'll explore some advanced concurrency patterns in Go, including the `select` statement and the `sync` package.

## The `select` Statement for managing multiple channels

The `select` statement lets a goroutine wait on multiple communication operations. A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```go
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
```

## The `sync` Package: `WaitGroup` for managing groups of Goroutines

The `sync` package provides synchronization primitives, such as `WaitGroup`, which can be used to wait for a collection of goroutines to finish.

A `WaitGroup` waits for a collection of goroutines to finish. The main goroutine calls `Add` to set the number of goroutines to wait for. Then, each of the goroutines runs and calls `Done` when finished. At the same time, `Wait` can be used to block until all goroutines have finished.

```go
import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
}
```

## The `sync` Package: `Mutex` for protecting shared data

A `Mutex` is a mutual exclusion lock. It is used to protect shared data from being accessed by multiple goroutines at the same time.

The zero value for a `Mutex` is an unlocked mutex.

```go
import (
    "fmt"
    "sync"
    "time"
)

// SafeCounter is safe to use concurrently.
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

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}
```

In this module, you've learned about some of the more advanced concurrency patterns in Go. These patterns will help you write more robust and efficient concurrent programs.

**Previous:** [Module 5: Introduction to Concurrency](../5/README.md)
**Next:** [Module 7: File System I/O](../7/README.md)
