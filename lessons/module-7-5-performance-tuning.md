# 🚀 Module 7.5: Performance Tuning & Large Files

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Optimize File Operations for Maximum Performance**

*From Zero to Go Hero - Module 7.5*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1-1.5 hours  | Advanced       | Modules 7.1-7.4, Understanding of concurrency basics |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Implement chunking strategies for large file processing
- ✅ Use memory-efficient streaming techniques
- ✅ Apply concurrent file processing patterns
- ✅ Benchmark and optimize file operations
- ✅ Build high-performance data processing systems

---

## 🌟 Understanding Performance Optimization

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎭 The Performance Analogy

Think of file performance optimization like organizing a massive warehouse:
- **Chunking**: Like moving boxes in optimal-sized loads
- **Streaming**: Like having a conveyor belt instead of carrying everything at once
- **Concurrency**: Like having multiple workers processing different sections
- **Benchmarking**: Like timing different approaches to find the fastest

</div>

### 🚄 **1. Performance Optimization Strategies**

```
┌─────────────────────────────────────────────────────────┐
│  🚀 Performance Optimization Impact                     │
├─────────────────────────────────────────────────────────┤
│  Chunking       50x faster      ████████████████████ 95%│
│  Streaming      20x faster      ████████████████▌    80%│
│  Concurrency    10x faster      ██████████████▌      70%│
│  Buffer Tuning  5x faster       ████████▌            40%│
└─────────────────────────────────────────────────────────┘
```

---

## 📦 Chunking Strategies for Large Files

### 🎯 **When to Use Chunking**
- Processing files larger than available RAM
- Streaming data processing
- Network file transfers
- Memory-constrained environments

### 💡 **Basic Chunking Implementation**

```go
package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

type ChunkProcessor struct {
    ChunkSize int
    Progress  chan int64
}

func NewChunkProcessor(chunkSize int) *ChunkProcessor {
    return &ChunkProcessor{
        ChunkSize: chunkSize,
        Progress:  make(chan int64, 100),
    }
}

func (cp *ChunkProcessor) ProcessFileInChunks(filename string, processor func([]byte) error) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    // Get file size for progress tracking
    info, err := file.Stat()
    if err != nil {
        return fmt.Errorf("failed to get file info: %w", err)
    }
    totalSize := info.Size()
    
    buffer := make([]byte, cp.ChunkSize)
    var totalProcessed int64
    chunkCount := 0
    
    fmt.Printf("📦 Processing %s in %d-byte chunks\n", filename, cp.ChunkSize)
    startTime := time.Now()
    
    for {
        bytesRead, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break
            }
            return fmt.Errorf("error reading chunk: %w", err)
        }
        
        // Process the chunk
        if err := processor(buffer[:bytesRead]); err != nil {
            return fmt.Errorf("error processing chunk %d: %w", chunkCount, err)
        }
        
        chunkCount++
        totalProcessed += int64(bytesRead)
        
        // Report progress
        progress := float64(totalProcessed) / float64(totalSize) * 100
        if chunkCount%100 == 0 || progress >= 100 {
            fmt.Printf("  📊 Progress: %.1f%% (%d chunks processed)\n", progress, chunkCount)
        }
    }
    
    duration := time.Since(startTime)
    throughput := float64(totalProcessed) / duration.Seconds() / 1024 / 1024 // MB/s
    
    fmt.Printf("✅ Processing complete!\n")
    fmt.Printf("  📈 Processed: %d bytes in %d chunks\n", totalProcessed, chunkCount)
    fmt.Printf("  ⏱️ Duration: %v\n", duration)
    fmt.Printf("  🚀 Throughput: %.2f MB/s\n", throughput)
    
    return nil
}

// Example: Word counting processor
func createWordCounter() func([]byte) error {
    totalWords := 0
    
    return func(chunk []byte) error {
        content := string(chunk)
        words := len(strings.Fields(content))
        totalWords += words
        
        // Store result somewhere accessible (in real app, use channels or shared state)
        if totalWords%10000 == 0 {
            fmt.Printf("    📝 Words counted so far: %d\n", totalWords)
        }
        
        return nil
    }
}

func main() {
    // Create a large test file
    testFile := "large_test.txt"
    fmt.Println("📝 Creating large test file...")
    
    file, err := os.Create(testFile)
    if err != nil {
        fmt.Printf("Error creating test file: %v\n", err)
        return
    }
    
    // Write 10MB of text data
    for i := 0; i < 100000; i++ {
        line := fmt.Sprintf("This is line %d with some sample text for processing. ", i)
        file.WriteString(line)
    }
    file.Close()
    
    // Test different chunk sizes
    chunkSizes := []int{1024, 4096, 16384, 65536} // 1KB, 4KB, 16KB, 64KB
    
    for _, chunkSize := range chunkSizes {
        fmt.Printf("\n🧪 Testing with %d-byte chunks:\n", chunkSize)
        
        processor := NewChunkProcessor(chunkSize)
        wordCounter := createWordCounter()
        
        err := processor.ProcessFileInChunks(testFile, wordCounter)
        if err != nil {
            fmt.Printf("Error processing file: %v\n", err)
        }
    }
    
    // Clean up
    os.Remove(testFile)
}
```

### 🔧 **Advanced Chunking with Memory Pool**

```go
package main

import (
    "fmt"
    "io"
    "os"
    "sync"
    "time"
)

type MemoryPool struct {
    pool     sync.Pool
    chunkSize int
}

func NewMemoryPool(chunkSize int) *MemoryPool {
    return &MemoryPool{
        pool: sync.Pool{
            New: func() interface{} {
                return make([]byte, chunkSize)
            },
        },
        chunkSize: chunkSize,
    }
}

func (mp *MemoryPool) GetBuffer() []byte {
    return mp.pool.Get().([]byte)
}

func (mp *MemoryPool) PutBuffer(buf []byte) {
    if len(buf) == mp.chunkSize {
        mp.pool.Put(buf)
    }
}

type OptimizedProcessor struct {
    memPool   *MemoryPool
    chunkSize int
}

func NewOptimizedProcessor(chunkSize int) *OptimizedProcessor {
    return &OptimizedProcessor{
        memPool:   NewMemoryPool(chunkSize),
        chunkSize: chunkSize,
    }
}

func (op *OptimizedProcessor) ProcessFile(filename string, processor func([]byte) error) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    var totalProcessed int64
    startTime := time.Now()
    
    for {
        // Get buffer from pool
        buffer := op.memPool.GetBuffer()
        
        bytesRead, err := file.Read(buffer)
        if err != nil {
            op.memPool.PutBuffer(buffer)
            if err == io.EOF {
                break
            }
            return fmt.Errorf("error reading file: %w", err)
        }
        
        // Process chunk
        if err := processor(buffer[:bytesRead]); err != nil {
            op.memPool.PutBuffer(buffer)
            return fmt.Errorf("error processing chunk: %w", err)
        }
        
        // Return buffer to pool
        op.memPool.PutBuffer(buffer)
        
        totalProcessed += int64(bytesRead)
    }
    
    duration := time.Since(startTime)
    throughput := float64(totalProcessed) / duration.Seconds() / 1024 / 1024
    
    fmt.Printf("✅ Optimized processing complete: %.2f MB/s\n", throughput)
    return nil
}

func main() {
    // Create test file
    testFile := "memory_pool_test.txt"
    content := make([]byte, 5*1024*1024) // 5MB
    for i := range content {
        content[i] = byte('A' + (i % 26))
    }
    os.WriteFile(testFile, content, 0644)
    defer os.Remove(testFile)
    
    // Test with and without memory pool
    processor := NewOptimizedProcessor(64 * 1024) // 64KB chunks
    
    simpleCounter := func(chunk []byte) error {
        // Simple processing - count bytes
        return nil
    }
    
    fmt.Println("🚀 Testing optimized processor with memory pool:")
    err := processor.ProcessFile(testFile, simpleCounter)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

---

## 🌊 Memory-Efficient Streaming

### 💡 **Streaming Data Pipeline**

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "time"
)

type DataPoint struct {
    Timestamp time.Time
    Value     float64
    Category  string
}

type StreamProcessor struct {
    input  io.Reader
    output io.Writer
    stats  map[string]float64
}

func NewStreamProcessor(input io.Reader, output io.Writer) *StreamProcessor {
    return &StreamProcessor{
        input:  input,
        output: output,
        stats:  make(map[string]float64),
    }
}

func (sp *StreamProcessor) ProcessStream() error {
    scanner := bufio.NewScanner(sp.input)
    writer := bufio.NewWriter(sp.output)
    defer writer.Flush()
    
    lineCount := 0
    processedCount := 0
    
    fmt.Println("🌊 Starting stream processing...")
    startTime := time.Now()
    
    for scanner.Scan() {
        lineCount++
        line := scanner.Text()
        
        // Parse data point
        dataPoint, err := sp.parseLine(line)
        if err != nil {
            continue // Skip invalid lines
        }
        
        // Process data point
        sp.updateStats(dataPoint)
        
        // Write processed result
        result := sp.formatOutput(dataPoint)
        writer.WriteString(result + "\n")
        
        processedCount++
        
        // Progress reporting
        if lineCount%10000 == 0 {
            fmt.Printf("  📊 Processed %d lines (%.1f%% valid)\n", 
                lineCount, float64(processedCount)/float64(lineCount)*100)
        }
    }
    
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("scanning error: %w", err)
    }
    
    duration := time.Since(startTime)
    linesPerSecond := float64(lineCount) / duration.Seconds()
    
    fmt.Printf("✅ Stream processing complete!\n")
    fmt.Printf("  📈 Lines processed: %d/%d (%.1f%% valid)\n", 
        processedCount, lineCount, float64(processedCount)/float64(lineCount)*100)
    fmt.Printf("  ⏱️ Duration: %v\n", duration)
    fmt.Printf("  🚀 Rate: %.0f lines/second\n", linesPerSecond)
    
    return nil
}

func (sp *StreamProcessor) parseLine(line string) (*DataPoint, error) {
    // Expected format: "2023-07-28T10:30:45,123.45,category_name"
    parts := strings.Split(line, ",")
    if len(parts) != 3 {
        return nil, fmt.Errorf("invalid format")
    }
    
    timestamp, err := time.Parse("2006-01-02T15:04:05", parts[0])
    if err != nil {
        return nil, fmt.Errorf("invalid timestamp: %w", err)
    }
    
    value, err := strconv.ParseFloat(parts[1], 64)
    if err != nil {
        return nil, fmt.Errorf("invalid value: %w", err)
    }
    
    return &DataPoint{
        Timestamp: timestamp,
        Value:     value,
        Category:  strings.TrimSpace(parts[2]),
    }, nil
}

func (sp *StreamProcessor) updateStats(dp *DataPoint) {
    sp.stats[dp.Category] += dp.Value
    sp.stats["total_count"]++
    sp.stats["total_value"] += dp.Value
}

func (sp *StreamProcessor) formatOutput(dp *DataPoint) string {
    return fmt.Sprintf("%s|%.2f|%s|processed", 
        dp.Timestamp.Format("2006-01-02 15:04:05"), 
        dp.Value, 
        dp.Category)
}

func (sp *StreamProcessor) PrintStats() {
    fmt.Println("\n📊 Processing Statistics:")
    for category, value := range sp.stats {
        fmt.Printf("  %s: %.2f\n", category, value)
    }
}

func main() {
    // Create sample data file
    inputFile := "stream_input.csv"
    outputFile := "stream_output.txt"
    
    fmt.Println("📝 Creating sample data...")
    file, err := os.Create(inputFile)
    if err != nil {
        fmt.Printf("Error creating input file: %v\n", err)
        return
    }
    
    // Generate 100,000 data points
    categories := []string{"sales", "marketing", "support", "development"}
    for i := 0; i < 100000; i++ {
        timestamp := time.Now().Add(-time.Duration(i) * time.Minute)
        value := float64(i%1000) + 0.5
        category := categories[i%len(categories)]
        
        line := fmt.Sprintf("%s,%.2f,%s\n", 
            timestamp.Format("2006-01-02T15:04:05"), 
            value, 
            category)
        file.WriteString(line)
    }
    file.Close()
    
    // Process stream
    input, err := os.Open(inputFile)
    if err != nil {
        fmt.Printf("Error opening input: %v\n", err)
        return
    }
    defer input.Close()
    
    output, err := os.Create(outputFile)
    if err != nil {
        fmt.Printf("Error creating output: %v\n", err)
        return
    }
    defer output.Close()
    
    processor := NewStreamProcessor(input, output)
    err = processor.ProcessStream()
    if err != nil {
        fmt.Printf("Error processing stream: %v\n", err)
        return
    }
    
    processor.PrintStats()
    
    // Clean up
    os.Remove(inputFile)
    os.Remove(outputFile)
}
```

---

## ⚡ Concurrent File Processing

### 💡 **Worker Pool Pattern for File Processing**

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "runtime"
    "sync"
    "time"
)

type FileJob struct {
    FilePath string
    FileSize int64
}

type ProcessingResult struct {
    FilePath    string
    Success     bool
    Error       error
    Duration    time.Duration
    LinesCount  int
}

type ConcurrentProcessor struct {
    WorkerCount int
    jobs        chan FileJob
    results     chan ProcessingResult
    wg          sync.WaitGroup
}

func NewConcurrentProcessor(workerCount int) *ConcurrentProcessor {
    if workerCount <= 0 {
        workerCount = runtime.NumCPU()
    }
    
    return &ConcurrentProcessor{
        WorkerCount: workerCount,
        jobs:        make(chan FileJob, workerCount*2),
        results:     make(chan ProcessingResult, workerCount*2),
    }
}

func (cp *ConcurrentProcessor) ProcessFiles(filePaths []string) []ProcessingResult {
    fmt.Printf("🚀 Starting concurrent processing with %d workers\n", cp.WorkerCount)
    
    // Start workers
    for i := 0; i < cp.WorkerCount; i++ {
        cp.wg.Add(1)
        go cp.worker(i)
    }
    
    // Start result collector
    results := make([]ProcessingResult, 0, len(filePaths))
    resultsDone := make(chan bool)
    go func() {
        for result := range cp.results {
            results = append(results, result)
            fmt.Printf("  ✅ Worker completed: %s (%.2fs)\n", 
                filepath.Base(result.FilePath), result.Duration.Seconds())
        }
        resultsDone <- true
    }()
    
    // Send jobs
    go func() {
        for _, filePath := range filePaths {
            info, err := os.Stat(filePath)
            if err != nil {
                continue
            }
            
            job := FileJob{
                FilePath: filePath,
                FileSize: info.Size(),
            }
            cp.jobs <- job
        }
        close(cp.jobs)
    }()
    
    // Wait for workers to complete
    cp.wg.Wait()
    close(cp.results)
    
    // Wait for result collection
    <-resultsDone
    
    return results
}

func (cp *ConcurrentProcessor) worker(id int) {
    defer cp.wg.Done()
    
    for job := range cp.jobs {
        startTime := time.Now()
        
        // Simulate file processing
        linesCount, err := cp.processFile(job.FilePath)
        
        result := ProcessingResult{
            FilePath:   job.FilePath,
            Success:    err == nil,
            Error:      err,
            Duration:   time.Since(startTime),
            LinesCount: linesCount,
        }
        
        cp.results <- result
    }
}

func (cp *ConcurrentProcessor) processFile(filePath string) (int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return 0, fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    // Count lines (simple processing example)
    scanner := bufio.NewScanner(file)
    lineCount := 0
    
    for scanner.Scan() {
        lineCount++
        // Simulate some processing time
        if lineCount%1000 == 0 {
            time.Sleep(1 * time.Millisecond)
        }
    }
    
    if err := scanner.Err(); err != nil {
        return lineCount, fmt.Errorf("scanning error: %w", err)
    }
    
    return lineCount, nil
}

func main() {
    // Create test files
    testDir := "concurrent_test"
    os.MkdirAll(testDir, 0755)
    defer os.RemoveAll(testDir)
    
    fmt.Println("📝 Creating test files...")
    testFiles := make([]string, 10)
    
    for i := 0; i < 10; i++ {
        filename := filepath.Join(testDir, fmt.Sprintf("file_%d.txt", i))
        file, err := os.Create(filename)
        if err != nil {
            continue
        }
        
        // Create files with different sizes
        lineCount := (i + 1) * 1000
        for j := 0; j < lineCount; j++ {
            file.WriteString(fmt.Sprintf("Line %d in file %d\n", j, i))
        }
        file.Close()
        
        testFiles[i] = filename
    }
    
    // Test sequential processing
    fmt.Println("\n⏱️ Sequential processing:")
    startTime := time.Now()
    sequentialResults := processSequentially(testFiles)
    sequentialDuration := time.Since(startTime)
    
    fmt.Printf("Sequential completed in: %v\n", sequentialDuration)
    
    // Test concurrent processing
    fmt.Println("\n⚡ Concurrent processing:")
    startTime = time.Now()
    processor := NewConcurrentProcessor(4)
    concurrentResults := processor.ProcessFiles(testFiles)
    concurrentDuration := time.Since(startTime)
    
    fmt.Printf("Concurrent completed in: %v\n", concurrentDuration)
    
    // Compare results
    if sequentialDuration > concurrentDuration {
        speedup := float64(sequentialDuration) / float64(concurrentDuration)
        fmt.Printf("🚀 Concurrent processing is %.2fx faster!\n", speedup)
    }
    
    // Summary
    fmt.Printf("\n📊 Processing Summary:\n")
    fmt.Printf("Files processed: %d\n", len(concurrentResults))
    totalLines := 0
    for _, result := range concurrentResults {
        totalLines += result.LinesCount
    }
    fmt.Printf("Total lines processed: %d\n", totalLines)
}

func processSequentially(filePaths []string) []ProcessingResult {
    results := make([]ProcessingResult, 0, len(filePaths))
    
    for _, filePath := range filePaths {
        startTime := time.Now()
        
        // Simple line counting
        file, err := os.Open(filePath)
        if err != nil {
            continue
        }
        
        scanner := bufio.NewScanner(file)
        lineCount := 0
        for scanner.Scan() {
            lineCount++
            if lineCount%1000 == 0 {
                time.Sleep(1 * time.Millisecond)
            }
        }
        file.Close()
        
        result := ProcessingResult{
            FilePath:   filePath,
            Success:    true,
            Duration:   time.Since(startTime),
            LinesCount: lineCount,
        }
        results = append(results, result)
    }
    
    return results
}
```

---

## 🎯 Key Takeaways

<div style="background: linear-gradient(45deg, #f093fb 0%, #f5576c 100%); padding: 15px; border-radius: 8px; color: white;">

**🎯 Performance Optimization Best Practices:**
- Use appropriate chunk sizes based on available memory and data patterns
- Implement memory pools for high-frequency operations
- Apply streaming for memory-efficient processing of large datasets
- Use concurrent processing for I/O-bound operations
- Always benchmark different approaches with realistic data
- Monitor memory usage and garbage collection impact

</div>

---

## 🚀 Module 7 Complete!

Congratulations! You've mastered all aspects of file system I/O in Go. You now have the skills to build efficient, scalable file processing applications that can handle everything from simple configuration files to massive data processing pipelines.

<div align="center">

[**👈 Back to Module 7 Overview**](./module-7-file-system-io.md) | [**👉 Continue to Module 8 →**](../module-8-data-processing-serialization.md)

</div>
