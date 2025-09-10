# 15. Advanced Concurrency Patterns in Go

You've learned about goroutines and channels, the basic building blocks of concurrency in Go. Now, let's explore some common and powerful patterns that you can build with them.

## Fan-Out, Fan-In

The fan-out, fan-in pattern is a way to distribute a set of tasks among multiple goroutines for parallel processing (fan-out), and then collect the results from them (fan-in).

### Fan-Out

Fan-out involves starting multiple goroutines to handle work from a single input channel. This is useful when you have a number of independent tasks that can be processed in parallel.

### Fan-In

Fan-in is the process of combining multiple channels into a single channel. This is useful for collecting the results from the goroutines that were fanned-out.

Here's an example that demonstrates the fan-out, fan-in pattern:

```go
package main

import (
    "fmt"
    "sync"
)

// This function takes a channel of numbers, squares them, and sends the result to an output channel.
func square(in <-chan int, out chan<- int) {
    for n := range in {
        out <- n * n
    }
}

func main() {
    in := make(chan int)
    out := make(chan int)

    // Fan-out: Start 3 goroutines to do the squaring work.
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            square(in, out)
        }()
    }

    // Send numbers to the input channel.
    go func() {
        for i := 1; i <= 10; i++ {
            in <- i
        }
        close(in)
    }()

    // Fan-in: Start a goroutine to wait for all the squaring goroutines to finish, then close the output channel.
    go func() {
        wg.Wait()
        close(out)
    }()

    // Read from the output channel until it's closed.
    for n := range out {
        fmt.Println(n)
    }
}
```
In this example, we use `sync.WaitGroup` to coordinate the goroutines. This is a common way to handle the fan-in part of the pattern.

## Worker Pools

A worker pool is a collection of goroutines that are available to perform a set of tasks. This pattern is useful when you have a large number of tasks to perform and you want to limit the number of concurrent goroutines to control resource usage.

Here's an example of a worker pool:

```go
package main

import (
    "fmt"
    "time"
)

// The worker function takes a channel of jobs and a channel of results.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("worker %d started job %d\n", id, j)
        time.Sleep(time.Second) // Simulate work
        fmt.Printf("worker %d finished job %d\n", id, j)
        results <- j * 2
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start 3 workers.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs to the jobs channel.
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect the results.
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

In this example, we have a fixed number of workers (3) that are processing jobs from a `jobs` channel. The `main` goroutine sends jobs to the `jobs` channel and then collects the results from the `results` channel.

These are just a couple of the many concurrency patterns you can use in Go. By combining goroutines and channels in different ways, you can build powerful and efficient concurrent systems.

In the next lesson, we will dive into another important aspect of concurrency: mutexes.
