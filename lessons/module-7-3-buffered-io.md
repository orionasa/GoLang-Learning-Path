# ⚡ Module 7.3: Buffered I/O with bufio

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Supercharge Your I/O Performance**

*From Zero to Go Hero - Module 7.3*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1-1.5 hours  | Intermediate   | Modules 7.1-7.2, Understanding of io.Reader/Writer |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand buffering concepts and performance benefits
- ✅ Master `bufio.Reader` and `bufio.Writer` for efficient I/O
- ✅ Use `bufio.Scanner` for structured data processing
- ✅ Optimize buffer sizes for different use cases
- ✅ Build high-performance data processing applications

---

## 🌟 Understanding Buffered I/O

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎭 The Buffering Analogy

Think of buffered I/O like different ways of moving water:
- **Unbuffered I/O**: Like using a tiny cup to move water drop by drop
- **Buffered I/O**: Like using a large bucket to move water efficiently
- **Scanner**: Like having a smart bucket that knows exactly how much to pour

Buffering reduces the number of system calls, dramatically improving performance!

</div>

### 🚄 **1. Performance Impact Visualization**

```
┌─────────────────────────────────────────────────────────┐
│  ⚡ I/O Performance Comparison (1MB file)               │
├─────────────────────────────────────────────────────────┤
│  Unbuffered     1000 syscalls    ██████████████████ 100%│
│  4KB Buffer     250 syscalls     ████▌              25% │
│  8KB Buffer     125 syscalls     ██▌                12% │
│  64KB Buffer    16 syscalls      ▌                   2% │
└─────────────────────────────────────────────────────────┘
```

---

## 📖 Understanding bufio.Reader

### 🎯 **When to Use bufio.Reader**
- Reading large files line by line
- Processing streaming data
- When you need to peek ahead in the data
- Reducing system call overhead

### 💡 **Basic bufio.Reader Usage**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func demonstrateBufferedReading() {
    // Create sample data
    sampleData := `Line 1: Introduction to buffered I/O
Line 2: Understanding performance benefits
Line 3: Working with bufio.Reader
Line 4: Scanning techniques and patterns
Line 5: Buffer size optimization strategies`
    
    // Write sample file
    err := os.WriteFile("sample.txt", []byte(sampleData), 0644)
    if err != nil {
        fmt.Printf("Error creating sample file: %v\n", err)
        return
    }
    defer os.Remove("sample.txt")
    
    // Open file for buffered reading
    file, err := os.Open("sample.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Create buffered reader with default buffer size (4096 bytes)
    reader := bufio.NewReader(file)
    
    fmt.Println("📖 Reading file line by line with bufio.Reader:")
    lineNum := 1
    
    for {
        line, err := reader.ReadLine()
        if err != nil {
            if err.Error() == "EOF" {
                break
            }
            fmt.Printf("Error reading line: %v\n", err)
            return
        }
        
        fmt.Printf("%d: %s\n", lineNum, string(line))
        lineNum++
    }
}

func main() {
    demonstrateBufferedReading()
}
```

### 🔧 **Advanced bufio.Reader Techniques**

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

type LogParser struct {
    reader *bufio.Reader
    stats  map[string]int
}

func NewLogParser(reader io.Reader) *LogParser {
    return &LogParser{
        reader: bufio.NewReader(reader),
        stats:  make(map[string]int),
    }
}

func (lp *LogParser) ParseLogs() error {
    for {
        // Read until delimiter (newline)
        line, err := lp.reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                // Process the last line if it doesn't end with newline
                if len(strings.TrimSpace(line)) > 0 {
                    lp.processLine(line)
                }
                break
            }
            return fmt.Errorf("error reading line: %w", err)
        }
        
        lp.processLine(line)
    }
    
    return nil
}

func (lp *LogParser) processLine(line string) {
    line = strings.TrimSpace(line)
    if line == "" {
        return
    }
    
    // Simple log level extraction
    if strings.Contains(line, "[INFO]") {
        lp.stats["INFO"]++
    } else if strings.Contains(line, "[WARN]") {
        lp.stats["WARN"]++
    } else if strings.Contains(line, "[ERROR]") {
        lp.stats["ERROR"]++
    } else {
        lp.stats["OTHER"]++
    }
}

func (lp *LogParser) GetStats() map[string]int {
    return lp.stats
}

func main() {
    // Create sample log data
    logData := `2023-07-28 10:00:01 [INFO] Application started
2023-07-28 10:00:02 [INFO] Database connection established
2023-07-28 10:00:15 [WARN] High memory usage detected: 85%
2023-07-28 10:00:30 [ERROR] Failed to process request: timeout
2023-07-28 10:00:31 [INFO] Retrying request processing
2023-07-28 10:00:32 [INFO] Request processed successfully
2023-07-28 10:01:00 [WARN] Disk space running low: 90% used
2023-07-28 10:01:15 [ERROR] Database connection lost
2023-07-28 10:01:16 [INFO] Attempting database reconnection
2023-07-28 10:01:17 [INFO] Database reconnected successfully`
    
    // Write log file
    err := os.WriteFile("app.log", []byte(logData), 0644)
    if err != nil {
        fmt.Printf("Error creating log file: %v\n", err)
        return
    }
    defer os.Remove("app.log")
    
    // Open and parse logs
    file, err := os.Open("app.log")
    if err != nil {
        fmt.Printf("Error opening log file: %v\n", err)
        return
    }
    defer file.Close()
    
    parser := NewLogParser(file)
    err = parser.ParseLogs()
    if err != nil {
        fmt.Printf("Error parsing logs: %v\n", err)
        return
    }
    
    // Display statistics
    fmt.Println("📊 Log Analysis Results:")
    stats := parser.GetStats()
    for level, count := range stats {
        fmt.Printf("  %s: %d entries\n", level, count)
    }
}
```

---

## ✏️ Understanding bufio.Writer

### 🎯 **When to Use bufio.Writer**
- Writing large amounts of data efficiently
- Reducing system call overhead
- Building high-performance data exporters
- Streaming data to files or network connections

### 💡 **Basic bufio.Writer Usage**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func demonstrateBufferedWriting() {
    // Create output file
    file, err := os.Create("buffered_output.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    defer os.Remove("buffered_output.txt")
    
    // Create buffered writer with custom buffer size (8KB)
    writer := bufio.NewWriterSize(file, 8192)
    
    fmt.Println("✏️ Writing data with buffered writer...")
    
    // Write multiple lines efficiently
    for i := 1; i <= 1000; i++ {
        line := fmt.Sprintf("Line %d: This is a sample line with timestamp %s\n", 
            i, time.Now().Format("15:04:05.000"))
        
        _, err := writer.WriteString(line)
        if err != nil {
            fmt.Printf("Error writing line %d: %v\n", i, err)
            return
        }
        
        // Show progress every 100 lines
        if i%100 == 0 {
            fmt.Printf("📝 Written %d lines\n", i)
        }
    }
    
    // IMPORTANT: Flush the buffer to ensure all data is written
    err = writer.Flush()
    if err != nil {
        fmt.Printf("Error flushing buffer: %v\n", err)
        return
    }
    
    fmt.Println("✅ All data written and flushed!")
    
    // Verify file size
    info, _ := file.Stat()
    fmt.Printf("📄 Final file size: %d bytes\n", info.Size())
}

func main() {
    demonstrateBufferedWriting()
}
```

### 🔧 **Performance Comparison: Buffered vs Unbuffered**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func writeUnbuffered(filename string, lines int) time.Duration {
    start := time.Now()
    
    file, err := os.Create(filename)
    if err != nil {
        fmt.Printf("Error creating unbuffered file: %v\n", err)
        return 0
    }
    defer file.Close()
    
    for i := 0; i < lines; i++ {
        line := fmt.Sprintf("Unbuffered line %d with some content\n", i)
        file.WriteString(line)
    }
    
    return time.Since(start)
}

func writeBuffered(filename string, lines int) time.Duration {
    start := time.Now()
    
    file, err := os.Create(filename)
    if err != nil {
        fmt.Printf("Error creating buffered file: %v\n", err)
        return 0
    }
    defer file.Close()
    
    writer := bufio.NewWriter(file)
    defer writer.Flush()
    
    for i := 0; i < lines; i++ {
        line := fmt.Sprintf("Buffered line %d with some content\n", i)
        writer.WriteString(line)
    }
    
    return time.Since(start)
}

func main() {
    lines := 10000
    
    fmt.Printf("🏁 Performance test: Writing %d lines\n", lines)
    
    // Test unbuffered writing
    fmt.Println("⏱️ Testing unbuffered writing...")
    unbufferedTime := writeUnbuffered("unbuffered.txt", lines)
    
    // Test buffered writing
    fmt.Println("⏱️ Testing buffered writing...")
    bufferedTime := writeBuffered("buffered.txt", lines)
    
    // Compare results
    fmt.Println("\n📊 Performance Results:")
    fmt.Printf("Unbuffered: %v\n", unbufferedTime)
    fmt.Printf("Buffered:   %v\n", bufferedTime)
    
    if unbufferedTime > bufferedTime {
        speedup := float64(unbufferedTime) / float64(bufferedTime)
        fmt.Printf("🚀 Buffered I/O is %.2fx faster!\n", speedup)
    }
    
    // Clean up
    os.Remove("unbuffered.txt")
    os.Remove("buffered.txt")
}
```

---

## 🔍 Master bufio.Scanner

### 🎯 **When to Use bufio.Scanner**
- Processing structured text data (CSV, logs, config files)
- Reading input line by line or word by word
- Custom delimiter-based parsing
- Memory-efficient processing of large files

### 💡 **Basic Scanner Usage**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func demonstrateScanner() {
    // Create sample CSV data
    csvData := `Name,Age,City,Salary
John Doe,30,New York,75000
Jane Smith,25,Los Angeles,68000
Bob Johnson,35,Chicago,82000
Alice Brown,28,Houston,71000
Charlie Wilson,32,Phoenix,79000`
    
    // Write CSV file
    err := os.WriteFile("employees.csv", []byte(csvData), 0644)
    if err != nil {
        fmt.Printf("Error creating CSV file: %v\n", err)
        return
    }
    defer os.Remove("employees.csv")
    
    // Open file for scanning
    file, err := os.Open("employees.csv")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Create scanner
    scanner := bufio.NewScanner(file)
    
    fmt.Println("📋 Processing CSV data with Scanner:")
    lineNum := 0
    var headers []string
    
    for scanner.Scan() {
        line := scanner.Text()
        lineNum++
        
        if lineNum == 1 {
            // Process header
            headers = strings.Split(line, ",")
            fmt.Printf("Headers: %v\n", headers)
            continue
        }
        
        // Process data rows
        fields := strings.Split(line, ",")
        if len(fields) == len(headers) {
            fmt.Printf("\nEmployee %d:\n", lineNum-1)
            for i, field := range fields {
                fmt.Printf("  %s: %s\n", headers[i], field)
            }
        }
    }
    
    // Check for scanning errors
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error scanning file: %v\n", err)
    }
}

func main() {
    demonstrateScanner()
}
```

### 🔧 **Custom Scanner with Different Split Functions**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Custom split function for processing key=value pairs
func scanKeyValuePairs(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }
    
    // Look for newline
    if i := strings.IndexByte(string(data), '\n'); i >= 0 {
        return i + 1, data[0:i], nil
    }
    
    // If at EOF, return remaining data
    if atEOF {
        return len(data), data, nil
    }
    
    // Need more data
    return 0, nil, nil
}

func processConfigFile() {
    // Create sample configuration
    configData := `# Application Configuration
server_port=8080
database_host=localhost
database_port=5432
debug_mode=true
max_connections=100
# End of configuration`
    
    // Write config file
    err := os.WriteFile("app.conf", []byte(configData), 0644)
    if err != nil {
        fmt.Printf("Error creating config file: %v\n", err)
        return
    }
    defer os.Remove("app.conf")
    
    // Open and scan config file
    file, err := os.Open("app.conf")
    if err != nil {
        fmt.Printf("Error opening config file: %v\n", err)
        return
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    scanner.Split(scanKeyValuePairs)
    
    config := make(map[string]string)
    
    fmt.Println("⚙️ Processing configuration file:")
    
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        
        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        
        // Parse key=value pairs
        if strings.Contains(line, "=") {
            parts := strings.SplitN(line, "=", 2)
            if len(parts) == 2 {
                key := strings.TrimSpace(parts[0])
                value := strings.TrimSpace(parts[1])
                config[key] = value
                fmt.Printf("  %s = %s\n", key, value)
            }
        }
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error scanning config: %v\n", err)
    }
    
    fmt.Printf("\n✅ Loaded %d configuration settings\n", len(config))
}

func main() {
    processConfigFile()
}
```

---

## 🚀 Advanced: Buffer Size Optimization

### 💡 **Buffer Size Impact Analysis**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func testBufferSizes(filename string, data []byte, bufferSizes []int) {
    fmt.Printf("🧪 Testing different buffer sizes with %d bytes of data\n", len(data))
    
    for _, size := range bufferSizes {
        // Write test file
        err := os.WriteFile(filename, data, 0644)
        if err != nil {
            fmt.Printf("Error creating test file: %v\n", err)
            continue
        }
        
        // Test read performance with this buffer size
        start := time.Now()
        
        file, err := os.Open(filename)
        if err != nil {
            fmt.Printf("Error opening file: %v\n", err)
            continue
        }
        
        reader := bufio.NewReaderSize(file, size)
        totalBytes := 0
        
        buffer := make([]byte, 1024) // Read in 1KB chunks
        for {
            n, err := reader.Read(buffer)
            totalBytes += n
            if err != nil {
                break
            }
        }
        
        file.Close()
        duration := time.Since(start)
        
        fmt.Printf("Buffer %6d bytes: %v (%d bytes read)\n", size, duration, totalBytes)
        
        os.Remove(filename)
    }
}

func findOptimalBufferSize() {
    // Create test data (1MB)
    testData := make([]byte, 1024*1024)
    for i := range testData {
        testData[i] = byte(i % 256)
    }
    
    // Test different buffer sizes
    bufferSizes := []int{
        1024,    // 1KB
        4096,    // 4KB (default)
        8192,    // 8KB
        16384,   // 16KB
        32768,   // 32KB
        65536,   // 64KB
        131072,  // 128KB
    }
    
    testBufferSizes("buffer_test.dat", testData, bufferSizes)
}

func main() {
    findOptimalBufferSize()
}
```

---

## 🧪 Hands-On Exercise: High-Performance Log Processor

### 🎯 **Challenge**
Build a high-performance log processor that can:
1. Process large log files efficiently using buffered I/O
2. Extract and analyze different log patterns
3. Generate summary reports
4. Handle multiple log formats
5. Provide real-time processing statistics

### 💡 **Starter Code**

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "regexp"
    "time"
)

type LogEntry struct {
    Timestamp time.Time
    Level     string
    Message   string
    Source    string
}

type LogProcessor struct {
    // TODO: Add fields for statistics and configuration
    BufferSize    int
    TotalLines    int
    ProcessedLines int
    ErrorCount    int
    StartTime     time.Time
}

func NewLogProcessor(bufferSize int) *LogProcessor {
    // TODO: Initialize processor
    return &LogProcessor{
        BufferSize: bufferSize,
        StartTime:  time.Now(),
    }
}

func (lp *LogProcessor) ProcessFile(filename string) error {
    // TODO: Implement efficient file processing
    // 1. Open file with buffered reader
    // 2. Process line by line
    // 3. Parse log entries
    // 4. Update statistics
    // 5. Handle errors gracefully
    return nil
}

func (lp *LogProcessor) ParseLogEntry(line string) (*LogEntry, error) {
    // TODO: Implement log parsing
    // Expected format: "2023-07-28 10:30:45 [INFO] [source] message"
    // Use regex for robust parsing
    return nil, nil
}

func (lp *LogProcessor) GenerateReport(writer io.Writer) error {
    // TODO: Generate processing report
    // Include: total lines, processing time, errors, statistics
    return nil
}

func (lp *LogProcessor) GetProcessingStats() map[string]interface{} {
    // TODO: Return processing statistics
    return make(map[string]interface{})
}

func main() {
    // Create large sample log file
    fmt.Println("📝 Creating sample log file...")
    
    // TODO: Generate realistic log data
    // Include different log levels, sources, and timestamps
    
    // TODO: Test the log processor
    // 1. Process the log file
    // 2. Display statistics
    // 3. Generate report
    // 4. Test with different buffer sizes
    
    fmt.Println("🚀 Log processor ready for implementation!")
}
```

### 🎯 **Solution**

<details>
<summary>Click to reveal solution</summary>

```go
type LogProcessor struct {
    BufferSize     int
    TotalLines     int
    ProcessedLines int
    ErrorCount     int
    StartTime      time.Time
    LevelCounts    map[string]int
    SourceCounts   map[string]int
    logPattern     *regexp.Regexp
}

func NewLogProcessor(bufferSize int) *LogProcessor {
    pattern := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\] \[([^\]]+)\] (.+)`)
    
    return &LogProcessor{
        BufferSize:   bufferSize,
        StartTime:    time.Now(),
        LevelCounts:  make(map[string]int),
        SourceCounts: make(map[string]int),
        logPattern:   pattern,
    }
}

func (lp *LogProcessor) ProcessFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    scanner.Buffer(make([]byte, lp.BufferSize), lp.BufferSize)
    
    for scanner.Scan() {
        lp.TotalLines++
        line := scanner.Text()
        
        entry, err := lp.ParseLogEntry(line)
        if err != nil {
            lp.ErrorCount++
            continue
        }
        
        lp.ProcessedLines++
        lp.LevelCounts[entry.Level]++
        lp.SourceCounts[entry.Source]++
        
        // Show progress every 1000 lines
        if lp.TotalLines%1000 == 0 {
            fmt.Printf("📊 Processed %d lines...\n", lp.TotalLines)
        }
    }
    
    return scanner.Err()
}

func (lp *LogProcessor) ParseLogEntry(line string) (*LogEntry, error) {
    matches := lp.logPattern.FindStringSubmatch(line)
    if len(matches) != 5 {
        return nil, fmt.Errorf("invalid log format")
    }
    
    timestamp, err := time.Parse("2006-01-02 15:04:05", matches[1])
    if err != nil {
        return nil, fmt.Errorf("invalid timestamp: %w", err)
    }
    
    return &LogEntry{
        Timestamp: timestamp,
        Level:     matches[2],
        Source:    matches[3],
        Message:   matches[4],
    }, nil
}

func (lp *LogProcessor) GenerateReport(writer io.Writer) error {
    duration := time.Since(lp.StartTime)
    
    fmt.Fprintf(writer, "📊 Log Processing Report\n")
    fmt.Fprintf(writer, "========================\n")
    fmt.Fprintf(writer, "Processing Time: %v\n", duration)
    fmt.Fprintf(writer, "Total Lines: %d\n", lp.TotalLines)
    fmt.Fprintf(writer, "Processed Lines: %d\n", lp.ProcessedLines)
    fmt.Fprintf(writer, "Error Lines: %d\n", lp.ErrorCount)
    fmt.Fprintf(writer, "Buffer Size: %d bytes\n", lp.BufferSize)
    
    if lp.TotalLines > 0 {
        linesPerSecond := float64(lp.TotalLines) / duration.Seconds()
        fmt.Fprintf(writer, "Processing Rate: %.2f lines/second\n", linesPerSecond)
    }
    
    fmt.Fprintf(writer, "\nLog Levels:\n")
    for level, count := range lp.LevelCounts {
        fmt.Fprintf(writer, "  %s: %d\n", level, count)
    }
    
    fmt.Fprintf(writer, "\nSources:\n")
    for source, count := range lp.SourceCounts {
        fmt.Fprintf(writer, "  %s: %d\n", source, count)
    }
    
    return nil
}
```

</details>

---

## 🎯 Key Takeaways

<div style="background: linear-gradient(45deg, #f093fb 0%, #f5576c 100%); padding: 15px; border-radius: 8px; color: white;">

**🎯 Buffered I/O Best Practices:**
- Use `bufio.Reader` for efficient reading of large files
- Use `bufio.Writer` for high-performance writing operations
- Always call `writer.Flush()` to ensure data is written
- Use `bufio.Scanner` for structured text processing
- Choose appropriate buffer sizes based on your data patterns
- Consider custom split functions for specialized parsing
- Monitor performance with different buffer sizes

</div>

---

## 🚀 What's Next?

Fantastic! You've mastered buffered I/O techniques that can dramatically improve your application's performance. Next, you'll learn how to navigate and manipulate file system structures programmatically.

<div align="center">

[**👈 Back to Module 7 Overview**](./module-7-file-system-io.md) | [**👉 Next: File Navigation →**](./module-7-4-file-navigation.md)

</div>
