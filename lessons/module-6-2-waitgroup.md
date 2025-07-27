# 👥 Lesson 6.2: WaitGroup Coordination

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Coordinate Multiple Goroutines with WaitGroup**

*Module 6 - Advanced Concurrency Patterns*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 90 minutes   | Intermediate   | Goroutines, Basic Channels |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand what `sync.WaitGroup` is and when to use it
- ✅ Coordinate multiple goroutines and wait for their completion
- ✅ Build worker pool patterns with proper synchronization
- ✅ Handle errors and results from multiple goroutines
- ✅ Master WaitGroup best practices and common pitfalls

---

## What is WaitGroup?

<div style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 The Goroutine Coordinator

`sync.WaitGroup` is like a **team coordinator** that waits for all team members to finish their tasks before proceeding. It's perfect when you need to ensure all goroutines complete before continuing.

**Think of it like:** A project manager waiting for all team members to submit their reports before compiling the final presentation!

</div>

### 🚄 **1. WaitGroup Lifecycle**

```
┌─────────────────────────────────────┐
│         WaitGroup Lifecycle         │
├─────────────────────────────────────┤
│  Add(n)     ████████████████████ Set │
│  Go func()  ███████████████████▌ Run │
│  Done()     ████████████▌       Dec │
│  Wait()     ████▌               Block│
└─────────────────────────────────────┘
```

The WaitGroup has three main methods:
- **`Add(n)`** - Increment the counter by n
- **`Done()`** - Decrement the counter by 1 (call when goroutine finishes)
- **`Wait()`** - Block until counter reaches 0

### 🧵 **2. Basic WaitGroup Pattern**

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Simple WaitGroup Example** - Wait for multiple tasks to complete

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    // Always call Done() when the goroutine finishes
    defer wg.Done()
    
    fmt.Printf("Worker %d starting\n", id)
    
    // Simulate work
    time.Sleep(time.Duration(id) * time.Second)
    
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    // Start 3 workers
    for i := 1; i <= 3; i++ {
        wg.Add(1) // Increment counter
        go worker(i, &wg)
    }
    
    fmt.Println("Waiting for all workers to finish...")
    wg.Wait() // Block until all Done() calls are made
    fmt.Println("All workers completed!")
}
```

**Output:**
```
Waiting for all workers to finish...
Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 1 finished
Worker 2 finished
Worker 3 finished
All workers completed!
```

</div>

### 🎯 **3. WaitGroup vs Channels**

- **WaitGroup** - When you only need to wait for completion (no data exchange)
- **Channels** - When you need to exchange data between goroutines
- **Both** - When you need data exchange AND coordination

### 🏢 **4. Common Use Cases**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🔄 Parallel Processing</strong><br>
Process multiple items<br>
simultaneously and wait<br>
for all to complete
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🌐 Web Scraping</strong><br>
Fetch multiple URLs<br>
concurrently and wait<br>
for all responses
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📁 File Processing</strong><br>
Process multiple files<br>
in parallel and wait<br>
for completion
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🔍 Database Queries</strong><br>
Execute multiple queries<br>
concurrently and wait<br>
for all results
</div>

</div>

---

## Worker Pool Pattern: Step-by-Step Guide

### **Step 1: Understanding Worker Pools**

A worker pool is a pattern where you have a fixed number of workers processing tasks from a queue:

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Task 1    │    │   Task 2    │    │   Task 3    │
│   Task 4    │───▶│   Task 5    │───▶│   Task 6    │
│   Task 7    │    │   Task 8    │    │   Task 9    │
└─────────────┘    └─────────────┘    └─────────────┘
     Queue           Worker 1          Worker 2
```

### **Step 2: Building a Worker Pool**

Create `worker_pool.go`:

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Task represents work to be done
type Task struct {
    ID   int
    Data string
}

// Result represents the output of a task
type Result struct {
    TaskID int
    Output string
    Error  error
}

// WorkerPool manages a pool of workers
type WorkerPool struct {
    numWorkers int
    tasks      chan Task
    results    chan Result
    wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(numWorkers int) *WorkerPool {
    return &WorkerPool{
        numWorkers: numWorkers,
        tasks:      make(chan Task, 100),
        results:    make(chan Result, 100),
    }
}

// Start launches all workers
func (wp *WorkerPool) Start() {
    fmt.Printf("🚀 Starting %d workers...\n", wp.numWorkers)
    
    for i := 1; i <= wp.numWorkers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

// worker processes tasks from the task channel
func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    
    for task := range wp.tasks {
        fmt.Printf("👷 Worker %d processing task %d\n", id, task.ID)
        
        // Simulate work with random duration
        workTime := time.Duration(rand.Intn(1000)) * time.Millisecond
        time.Sleep(workTime)
        
        // Simulate occasional errors
        var result Result
        if rand.Float32() < 0.1 { // 10% chance of error
            result = Result{
                TaskID: task.ID,
                Error:  fmt.Errorf("task %d failed", task.ID),
            }
        } else {
            result = Result{
                TaskID: task.ID,
                Output: fmt.Sprintf("Processed: %s (by worker %d)", task.Data, id),
            }
        }
        
        wp.results <- result
        fmt.Printf("✅ Worker %d completed task %d\n", id, task.ID)
    }
    
    fmt.Printf("🛑 Worker %d shutting down\n", id)
}

// AddTask adds a task to the pool
func (wp *WorkerPool) AddTask(task Task) {
    wp.tasks <- task
}

// Stop closes the task channel and waits for workers to finish
func (wp *WorkerPool) Stop() {
    close(wp.tasks)
    wp.wg.Wait()
    close(wp.results)
    fmt.Println("🏁 All workers stopped")
}

// GetResults returns the results channel
func (wp *WorkerPool) GetResults() <-chan Result {
    return wp.results
}

func main() {
    // Create worker pool with 3 workers
    pool := NewWorkerPool(3)
    
    // Start the workers
    pool.Start()
    
    // Add tasks
    tasks := []Task{
        {1, "Process user data"},
        {2, "Generate report"},
        {3, "Send email"},
        {4, "Update database"},
        {5, "Backup files"},
        {6, "Validate input"},
        {7, "Transform data"},
        {8, "Create thumbnail"},
    }
    
    fmt.Printf("📋 Adding %d tasks...\n", len(tasks))
    for _, task := range tasks {
        pool.AddTask(task)
    }
    
    // Collect results in a separate goroutine
    var resultsWG sync.WaitGroup
    resultsWG.Add(1)
    
    go func() {
        defer resultsWG.Done()
        
        successCount := 0
        errorCount := 0
        
        for result := range pool.GetResults() {
            if result.Error != nil {
                fmt.Printf("❌ Task %d failed: %v\n", result.TaskID, result.Error)
                errorCount++
            } else {
                fmt.Printf("✅ Task %d: %s\n", result.TaskID, result.Output)
                successCount++
            }
        }
        
        fmt.Printf("\n📊 Summary: %d successful, %d failed\n", successCount, errorCount)
    }()
    
    // Stop the pool and wait for all workers to finish
    pool.Stop()
    
    // Wait for result collection to finish
    resultsWG.Wait()
    
    fmt.Println("🎉 Demo completed!")
}
```

### **Step 3: Advanced Worker Pool with Context**

Create `advanced_worker_pool.go`:

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Job represents work with context support
type Job struct {
    ID      int
    Payload string
    Timeout time.Duration
}

// JobResult contains the result of a job
type JobResult struct {
    JobID    int
    Result   string
    Duration time.Duration
    Error    error
}

// AdvancedWorkerPool with context support and graceful shutdown
type AdvancedWorkerPool struct {
    workerCount int
    jobs        chan Job
    results     chan JobResult
    wg          sync.WaitGroup
    ctx         context.Context
    cancel      context.CancelFunc
}

func NewAdvancedWorkerPool(workerCount int) *AdvancedWorkerPool {
    ctx, cancel := context.WithCancel(context.Background())
    
    return &AdvancedWorkerPool{
        workerCount: workerCount,
        jobs:        make(chan Job, 50),
        results:     make(chan JobResult, 50),
        ctx:         ctx,
        cancel:      cancel,
    }
}

func (awp *AdvancedWorkerPool) Start() {
    fmt.Printf("🚀 Starting %d advanced workers...\n", awp.workerCount)
    
    for i := 1; i <= awp.workerCount; i++ {
        awp.wg.Add(1)
        go awp.worker(i)
    }
}

func (awp *AdvancedWorkerPool) worker(id int) {
    defer awp.wg.Done()
    
    for {
        select {
        case job, ok := <-awp.jobs:
            if !ok {
                fmt.Printf("🛑 Worker %d: jobs channel closed\n", id)
                return
            }
            
            awp.processJob(id, job)
            
        case <-awp.ctx.Done():
            fmt.Printf("🛑 Worker %d: context cancelled\n", id)
            return
        }
    }
}

func (awp *AdvancedWorkerPool) processJob(workerID int, job Job) {
    start := time.Now()
    
    // Create job-specific context with timeout
    jobCtx, cancel := context.WithTimeout(awp.ctx, job.Timeout)
    defer cancel()
    
    fmt.Printf("👷 Worker %d starting job %d (timeout: %v)\n", 
        workerID, job.ID, job.Timeout)
    
    // Simulate work with potential timeout
    done := make(chan JobResult, 1)
    
    go func() {
        // Simulate work
        time.Sleep(time.Duration(job.ID*200) * time.Millisecond)
        
        done <- JobResult{
            JobID:    job.ID,
            Result:   fmt.Sprintf("Processed '%s' by worker %d", job.Payload, workerID),
            Duration: time.Since(start),
        }
    }()
    
    select {
    case result := <-done:
        awp.results <- result
        fmt.Printf("✅ Worker %d completed job %d in %v\n", 
            workerID, job.ID, result.Duration)
            
    case <-jobCtx.Done():
        result := JobResult{
            JobID:    job.ID,
            Duration: time.Since(start),
            Error:    fmt.Errorf("job %d timed out after %v", job.ID, job.Timeout),
        }
        awp.results <- result
        fmt.Printf("⏰ Worker %d: job %d timed out\n", workerID, job.ID)
    }
}

func (awp *AdvancedWorkerPool) SubmitJob(job Job) {
    select {
    case awp.jobs <- job:
        // Job submitted successfully
    case <-awp.ctx.Done():
        fmt.Printf("❌ Cannot submit job %d: pool is shutting down\n", job.ID)
    }
}

func (awp *AdvancedWorkerPool) Shutdown() {
    fmt.Println("🛑 Initiating graceful shutdown...")
    
    // Close jobs channel to signal no more jobs
    close(awp.jobs)
    
    // Wait for all workers to finish current jobs
    awp.wg.Wait()
    
    // Close results channel
    close(awp.results)
    
    fmt.Println("🏁 Graceful shutdown completed")
}

func (awp *AdvancedWorkerPool) ForceShutdown() {
    fmt.Println("🚨 Force shutdown initiated...")
    awp.cancel()
    awp.wg.Wait()
    close(awp.results)
    fmt.Println("🏁 Force shutdown completed")
}

func (awp *AdvancedWorkerPool) Results() <-chan JobResult {
    return awp.results
}

func main() {
    pool := NewAdvancedWorkerPool(3)
    pool.Start()
    
    // Submit jobs with different timeouts
    jobs := []Job{
        {1, "Quick task", 1 * time.Second},
        {2, "Medium task", 800 * time.Millisecond},
        {3, "Slow task", 500 * time.Millisecond}, // This will timeout
        {4, "Another quick task", 1 * time.Second},
        {5, "Final task", 2 * time.Second},
    }
    
    // Submit all jobs
    for _, job := range jobs {
        pool.SubmitJob(job)
    }
    
    // Collect results
    go func() {
        for result := range pool.Results() {
            if result.Error != nil {
                fmt.Printf("❌ Job %d failed: %v\n", result.JobID, result.Error)
            } else {
                fmt.Printf("✅ Job %d: %s (took %v)\n", 
                    result.JobID, result.Result, result.Duration)
            }
        }
    }()
    
    // Let jobs run for a bit
    time.Sleep(3 * time.Second)
    
    // Graceful shutdown
    pool.Shutdown()
    
    // Wait a bit to see final results
    time.Sleep(500 * time.Millisecond)
    
    fmt.Println("🎉 Advanced worker pool demo completed!")
}
```

---

## Real-World Project: Concurrent File Processor

### **Step 4: Building a File Processing System**

Create `file_processor.go`:

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "sync"
    "time"
)

// FileTask represents a file processing task
type FileTask struct {
    FilePath string
    TaskType string
}

// FileResult represents the result of file processing
type FileResult struct {
    FilePath string
    TaskType string
    Success  bool
    Message  string
    Duration time.Duration
}

// FileProcessor handles concurrent file operations
type FileProcessor struct {
    maxWorkers int
    wg         sync.WaitGroup
}

func NewFileProcessor(maxWorkers int) *FileProcessor {
    return &FileProcessor{
        maxWorkers: maxWorkers,
    }
}

// ProcessFiles processes multiple files concurrently
func (fp *FileProcessor) ProcessFiles(tasks []FileTask) []FileResult {
    taskChan := make(chan FileTask, len(tasks))
    resultChan := make(chan FileResult, len(tasks))
    
    // Start workers
    for i := 0; i < fp.maxWorkers; i++ {
        fp.wg.Add(1)
        go fp.worker(i+1, taskChan, resultChan)
    }
    
    // Send tasks
    for _, task := range tasks {
        taskChan <- task
    }
    close(taskChan)
    
    // Wait for all workers to complete
    fp.wg.Wait()
    close(resultChan)
    
    // Collect results
    var results []FileResult
    for result := range resultChan {
        results = append(results, result)
    }
    
    return results
}

func (fp *FileProcessor) worker(id int, tasks <-chan FileTask, results chan<- FileResult) {
    defer fp.wg.Done()
    
    for task := range tasks {
        start := time.Now()
        
        fmt.Printf("👷 Worker %d processing: %s (%s)\n", 
            id, filepath.Base(task.FilePath), task.TaskType)
        
        result := fp.processFile(task)
        result.Duration = time.Since(start)
        
        results <- result
        
        fmt.Printf("✅ Worker %d completed: %s in %v\n", 
            id, filepath.Base(task.FilePath), result.Duration)
    }
}

func (fp *FileProcessor) processFile(task FileTask) FileResult {
    result := FileResult{
        FilePath: task.FilePath,
        TaskType: task.TaskType,
    }
    
    switch task.TaskType {
    case "count_lines":
        count, err := fp.countLines(task.FilePath)
        if err != nil {
            result.Success = false
            result.Message = fmt.Sprintf("Error: %v", err)
        } else {
            result.Success = true
            result.Message = fmt.Sprintf("Lines: %d", count)
        }
        
    case "get_size":
        size, err := fp.getFileSize(task.FilePath)
        if err != nil {
            result.Success = false
            result.Message = fmt.Sprintf("Error: %v", err)
        } else {
            result.Success = true
            result.Message = fmt.Sprintf("Size: %d bytes", size)
        }
        
    case "check_extension":
        ext := filepath.Ext(task.FilePath)
        result.Success = true
        result.Message = fmt.Sprintf("Extension: %s", ext)
        
    default:
        result.Success = false
        result.Message = "Unknown task type"
    }
    
    // Simulate processing time
    time.Sleep(time.Duration(100+len(task.FilePath)*2) * time.Millisecond)
    
    return result
}

func (fp *FileProcessor) countLines(filePath string) (int, error) {
    content, err := os.ReadFile(filePath)
    if err != nil {
        return 0, err
    }
    
    lines := strings.Split(string(content), "\n")
    return len(lines), nil
}

func (fp *FileProcessor) getFileSize(filePath string) (int64, error) {
    info, err := os.Stat(filePath)
    if err != nil {
        return 0, err
    }
    return info.Size(), nil
}

// createSampleFiles creates some sample files for testing
func createSampleFiles() []string {
    files := []string{
        "sample1.txt",
        "sample2.txt", 
        "sample3.txt",
        "sample4.txt",
        "sample5.txt",
    }
    
    contents := []string{
        "Line 1\nLine 2\nLine 3",
        "Hello\nWorld\nFrom\nGo\nConcurrency",
        "Single line file",
        "Multiple\nlines\nwith\ndifferent\ncontent\nhere",
        "Short\nFile",
    }
    
    for i, file := range files {
        os.WriteFile(file, []byte(contents[i]), 0644)
    }
    
    return files
}

func main() {
    fmt.Println("🗂️  Concurrent File Processor Demo")
    fmt.Println("=====================================")
    
    // Create sample files
    files := createSampleFiles()
    defer func() {
        // Cleanup
        for _, file := range files {
            os.Remove(file)
        }
    }()
    
    // Create tasks
    var tasks []FileTask
    taskTypes := []string{"count_lines", "get_size", "check_extension"}
    
    for _, file := range files {
        for _, taskType := range taskTypes {
            tasks = append(tasks, FileTask{
                FilePath: file,
                TaskType: taskType,
            })
        }
    }
    
    fmt.Printf("📋 Created %d tasks for %d files\n", len(tasks), len(files))
    
    // Process files with different worker counts
    workerCounts := []int{1, 3, 5}
    
    for _, workerCount := range workerCounts {
        fmt.Printf("\n🚀 Processing with %d workers...\n", workerCount)
        
        start := time.Now()
        processor := NewFileProcessor(workerCount)
        results := processor.ProcessFiles(tasks)
        duration := time.Since(start)
        
        // Analyze results
        successCount := 0
        for _, result := range results {
            if result.Success {
                successCount++
            }
        }
        
        fmt.Printf("📊 Results: %d/%d successful in %v\n", 
            successCount, len(results), duration)
    }
    
    fmt.Println("\n🎉 File processor demo completed!")
}
```

---

## 🎯 Practice Exercises

<div style="background: #e7f3ff; border-left: 4px solid #2196f3; padding: 15px; margin: 10px 0;">

### Exercise 1: Parallel Sum Calculator
Create a program that calculates the sum of large slices in parallel using multiple goroutines and WaitGroup.

### Exercise 2: Web URL Checker
Build a concurrent URL checker that tests multiple URLs for availability and reports results.

### Exercise 3: Image Resizer Pool
Create a worker pool that processes image files (simulate with sleep) and reports completion status.

</div>

---

## ⚠️ Common Pitfalls and Best Practices

<div style="background: #fff3cd; border-left: 4px solid #ffc107; padding: 15px; margin: 10px 0;">

### Common Mistakes:
1. **Forgetting `defer wg.Done()`** - Always use defer to ensure Done() is called
2. **Adding after starting goroutines** - Call Add() before starting goroutines
3. **Not handling panics** - Use recover() in goroutines to prevent crashes
4. **Reusing WaitGroup** - Create new WaitGroup for each batch of work

### Best Practices:
1. **Pass WaitGroup by pointer** - Always pass `*sync.WaitGroup`
2. **Use defer for Done()** - Ensures Done() is called even if panic occurs
3. **Add before Go** - Call Add() before launching goroutines
4. **Limit goroutine count** - Use worker pools for large numbers of tasks

</div>

---

## 🔑 Key Takeaways

- **WaitGroup coordinates** multiple goroutines without data exchange
- **Always use defer wg.Done()** to ensure proper cleanup
- **Worker pools** are perfect for processing many similar tasks
- **Context support** enables timeouts and cancellation
- **Graceful shutdown** is important for production applications

---

## 🚀 What's Next?

Excellent work mastering WaitGroup coordination! You now know how to:
- Coordinate multiple goroutines effectively
- Build robust worker pool patterns
- Handle timeouts and graceful shutdowns
- Process tasks concurrently with proper synchronization

**Next up:** [Lesson 6.3 - Mutex & Data Protection](./module-6-3-mutex.md) where you'll learn to protect shared data from race conditions using mutexes.

---

<div align="center">

**🎉 Congratulations on completing Lesson 6.2! 🎉**

*You're mastering Go's concurrency patterns!*

</div>
