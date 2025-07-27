# 📖 Module 7.1: Reading Files with os & io

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master File Reading in Go**

*From Zero to Go Hero - Module 7.1*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1-1.5 hours  | Intermediate   | Basic Go syntax, error handling, structs |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand different approaches to reading files in Go
- ✅ Use `os.Open()` and `os.ReadFile()` effectively
- ✅ Work with the `io.Reader` interface for flexible file handling
- ✅ Implement proper error handling and resource management
- ✅ Build practical file reading applications

---

## 🌟 Understanding File Reading in Go

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎭 The File Reading Analogy

Think of reading files like reading a book in a library:
- **os.ReadFile()**: Like photocopying the entire book at once
- **os.Open() + Read()**: Like checking out the book and reading page by page
- **io.Reader**: Like having different reading strategies (speed reading, detailed study, etc.)

</div>

### 🚄 **1. File Reading Approaches Comparison**

```
┌─────────────────────────────────────────────────────────┐
│  📚 File Reading Methods Performance & Use Cases        │
├─────────────────────────────────────────────────────────┤
│  Small Files    os.ReadFile()     ████████████████ 95% │
│  Large Files    os.Open()+Read()  ███████████████▌ 88% │
│  Streaming      io.Reader         ████████████▌   75% │
│  Line-by-line   bufio.Scanner     ██████████████▌  85% │
└─────────────────────────────────────────────────────────┘
```

---

## 📖 Method 1: Simple File Reading with os.ReadFile()

### 🎯 **When to Use**
- Small to medium files (< 100MB)
- When you need the entire file content at once
- Simple, one-shot operations

### 💡 **Basic Implementation**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Read entire file into memory
    content, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("File content:\n%s\n", string(content))
    fmt.Printf("File size: %d bytes\n", len(content))
}
```

### 🔧 **Enhanced Example: Configuration File Reader**

```go
package main

import (
    "fmt"
    "os"
    "strings"
)

type Config struct {
    Settings map[string]string
}

func readConfig(filename string) (*Config, error) {
    // Read the entire configuration file
    content, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    
    config := &Config{
        Settings: make(map[string]string),
    }
    
    // Parse simple key=value format
    lines := strings.Split(string(content), "\n")
    for lineNum, line := range lines {
        line = strings.TrimSpace(line)
        
        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        
        // Parse key=value pairs
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            return nil, fmt.Errorf("invalid config format at line %d: %s", lineNum+1, line)
        }
        
        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])
        config.Settings[key] = value
    }
    
    return config, nil
}

func main() {
    // Create a sample config file
    sampleConfig := `# Application Configuration
server_port=8080
database_url=localhost:5432
debug_mode=true
max_connections=100
`
    
    // Write sample config
    err := os.WriteFile("app.config", []byte(sampleConfig), 0644)
    if err != nil {
        fmt.Printf("Error creating config file: %v\n", err)
        return
    }
    
    // Read and parse config
    config, err := readConfig("app.config")
    if err != nil {
        fmt.Printf("Error reading config: %v\n", err)
        return
    }
    
    // Display configuration
    fmt.Println("📋 Application Configuration:")
    for key, value := range config.Settings {
        fmt.Printf("  %s = %s\n", key, value)
    }
    
    // Clean up
    os.Remove("app.config")
}
```

---

## 📂 Method 2: Controlled Reading with os.Open()

### 🎯 **When to Use**
- Large files that don't fit in memory
- When you need to read in chunks
- More control over the reading process

### 💡 **Basic Implementation**

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func readFileInChunks(filename string, chunkSize int) error {
    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close() // Always close the file!
    
    // Create a buffer for reading chunks
    buffer := make([]byte, chunkSize)
    totalBytes := 0
    chunkCount := 0
    
    for {
        // Read a chunk
        bytesRead, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break // End of file reached
            }
            return fmt.Errorf("error reading file: %w", err)
        }
        
        chunkCount++
        totalBytes += bytesRead
        
        fmt.Printf("📦 Chunk %d: Read %d bytes\n", chunkCount, bytesRead)
        
        // Process the chunk (first 50 characters for demo)
        preview := string(buffer[:bytesRead])
        if len(preview) > 50 {
            preview = preview[:50] + "..."
        }
        fmt.Printf("   Content preview: %q\n", preview)
    }
    
    fmt.Printf("\n✅ Total: %d bytes read in %d chunks\n", totalBytes, chunkCount)
    return nil
}

func main() {
    // Create a sample large file
    largeContent := ""
    for i := 0; i < 1000; i++ {
        largeContent += fmt.Sprintf("This is line %d with some content to make it longer.\n", i+1)
    }
    
    err := os.WriteFile("large_file.txt", []byte(largeContent), 0644)
    if err != nil {
        fmt.Printf("Error creating large file: %v\n", err)
        return
    }
    
    fmt.Println("🚀 Reading large file in chunks...")
    err = readFileInChunks("large_file.txt", 1024) // 1KB chunks
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Clean up
    os.Remove("large_file.txt")
}
```

---

## 🔄 Method 3: Flexible Reading with io.Reader Interface

### 🎯 **When to Use**
- Building reusable, testable code
- Working with different input sources (files, network, memory)
- Creating flexible data processing pipelines

### 💡 **Interface-Based Implementation**

```go
package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

// DataProcessor can work with any io.Reader
type DataProcessor struct {
    ChunkSize int
}

func (dp *DataProcessor) ProcessData(reader io.Reader, processor func([]byte) error) error {
    buffer := make([]byte, dp.ChunkSize)
    
    for {
        bytesRead, err := reader.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break
            }
            return fmt.Errorf("reading error: %w", err)
        }
        
        // Process the chunk
        if err := processor(buffer[:bytesRead]); err != nil {
            return fmt.Errorf("processing error: %w", err)
        }
    }
    
    return nil
}

// Word counter processor
func createWordCounter() (func([]byte) error, func() int) {
    wordCount := 0
    
    processor := func(data []byte) error {
        content := string(data)
        words := strings.Fields(content)
        wordCount += len(words)
        return nil
    }
    
    getCount := func() int {
        return wordCount
    }
    
    return processor, getCount
}

// Character frequency analyzer
func createCharFrequencyAnalyzer() (func([]byte) error, func() map[rune]int) {
    frequency := make(map[rune]int)
    
    processor := func(data []byte) error {
        for _, char := range string(data) {
            if char != ' ' && char != '\n' && char != '\t' {
                frequency[char]++
            }
        }
        return nil
    }
    
    getFrequency := func() map[rune]int {
        return frequency
    }
    
    return processor, getFrequency
}

func main() {
    // Create sample data
    sampleText := `
    The Go programming language is an open source project to make programmers more productive.
    Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to
    write programs that get the most out of multicore and networked machines, while its novel
    type system enables flexible and modular program construction.
    `
    
    // Write to file
    err := os.WriteFile("sample.txt", []byte(sampleText), 0644)
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    
    // Open file for reading
    file, err := os.Open("sample.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    processor := &DataProcessor{ChunkSize: 64}
    
    fmt.Println("📊 Processing file with multiple analyzers...")
    
    // Word count analysis
    wordProcessor, getWordCount := createWordCounter()
    err = processor.ProcessData(file, wordProcessor)
    if err != nil {
        fmt.Printf("Error processing words: %v\n", err)
        return
    }
    
    fmt.Printf("📝 Total words: %d\n", getWordCount())
    
    // Reset file position for next analysis
    file.Seek(0, 0)
    
    // Character frequency analysis
    charProcessor, getCharFreq := createCharFrequencyAnalyzer()
    err = processor.ProcessData(file, charProcessor)
    if err != nil {
        fmt.Printf("Error processing characters: %v\n", err)
        return
    }
    
    fmt.Println("\n📈 Character frequency (top 10):")
    freq := getCharFreq()
    count := 0
    for char, frequency := range freq {
        if count >= 10 {
            break
        }
        fmt.Printf("  '%c': %d times\n", char, frequency)
        count++
    }
    
    // Clean up
    os.Remove("sample.txt")
}
```

---

## 🛡️ Advanced Error Handling & Resource Management

### 💡 **Robust File Reading Pattern**

```go
package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

type FileReader struct {
    MaxFileSize int64 // Maximum file size to read (in bytes)
}

func NewFileReader(maxSize int64) *FileReader {
    return &FileReader{
        MaxFileSize: maxSize,
    }
}

func (fr *FileReader) ReadFileSecurely(filename string) ([]byte, error) {
    // Validate file path
    cleanPath := filepath.Clean(filename)
    if cleanPath != filename {
        return nil, fmt.Errorf("invalid file path: %s", filename)
    }
    
    // Get file info before opening
    info, err := os.Stat(filename)
    if err != nil {
        return nil, fmt.Errorf("file stat error: %w", err)
    }
    
    // Check if it's a regular file
    if !info.Mode().IsRegular() {
        return nil, fmt.Errorf("not a regular file: %s", filename)
    }
    
    // Check file size
    if info.Size() > fr.MaxFileSize {
        return nil, fmt.Errorf("file too large: %d bytes (max: %d)", info.Size(), fr.MaxFileSize)
    }
    
    // Open file
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    defer func() {
        if closeErr := file.Close(); closeErr != nil {
            fmt.Printf("Warning: failed to close file: %v\n", closeErr)
        }
    }()
    
    // Read with size limit
    content := make([]byte, info.Size())
    bytesRead, err := io.ReadFull(file, content)
    if err != nil {
        return nil, fmt.Errorf("failed to read file completely: %w", err)
    }
    
    fmt.Printf("✅ Successfully read %d bytes from %s\n", bytesRead, filename)
    return content, nil
}

func main() {
    reader := NewFileReader(1024 * 1024) // 1MB limit
    
    // Create test file
    testContent := "This is a test file for secure reading demonstration."
    err := os.WriteFile("test_secure.txt", []byte(testContent), 0644)
    if err != nil {
        fmt.Printf("Error creating test file: %v\n", err)
        return
    }
    
    // Test secure reading
    content, err := reader.ReadFileSecurely("test_secure.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("📄 File content: %s\n", string(content))
    
    // Test with non-existent file
    _, err = reader.ReadFileSecurely("nonexistent.txt")
    if err != nil {
        fmt.Printf("🚫 Expected error for non-existent file: %v\n", err)
    }
    
    // Clean up
    os.Remove("test_secure.txt")
}
```

---

## 🧪 Hands-On Exercise: Log File Analyzer

### 🎯 **Challenge**
Build a log file analyzer that can:
1. Read log files of any size
2. Count different log levels (INFO, WARN, ERROR)
3. Extract timestamps and find the time range
4. Handle malformed lines gracefully

### 💡 **Starter Code**

```go
package main

import (
    "fmt"
    "os"
    "regexp"
    "strings"
    "time"
)

type LogAnalyzer struct {
    InfoCount    int
    WarnCount    int
    ErrorCount   int
    FirstTime    *time.Time
    LastTime     *time.Time
    TotalLines   int
    MalformedLines int
}

func (la *LogAnalyzer) AnalyzeLine(line string) {
    la.TotalLines++
    
    // TODO: Implement log line analysis
    // Expected format: "2023-07-28 10:30:45 [INFO] Message content"
    // Extract timestamp and log level
    // Update counters and time range
    
    // Hint: Use regex for parsing
    // Pattern: `(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\]`
}

func (la *LogAnalyzer) PrintSummary() {
    fmt.Println("\n📊 Log Analysis Summary")
    fmt.Println("=" * 30)
    fmt.Printf("Total Lines: %d\n", la.TotalLines)
    fmt.Printf("INFO: %d\n", la.InfoCount)
    fmt.Printf("WARN: %d\n", la.WarnCount)
    fmt.Printf("ERROR: %d\n", la.ErrorCount)
    fmt.Printf("Malformed: %d\n", la.MalformedLines)
    
    if la.FirstTime != nil && la.LastTime != nil {
        duration := la.LastTime.Sub(*la.FirstTime)
        fmt.Printf("Time Range: %s to %s (Duration: %v)\n", 
            la.FirstTime.Format("2006-01-02 15:04:05"),
            la.LastTime.Format("2006-01-02 15:04:05"),
            duration)
    }
}

func main() {
    // Create sample log file
    sampleLog := `2023-07-28 10:30:45 [INFO] Application started
2023-07-28 10:30:46 [INFO] Database connection established
2023-07-28 10:31:15 [WARN] High memory usage detected
2023-07-28 10:31:45 [ERROR] Failed to process request
Invalid log line without proper format
2023-07-28 10:32:00 [INFO] Request processed successfully
2023-07-28 10:32:30 [ERROR] Database connection lost
2023-07-28 10:33:00 [INFO] Application shutdown initiated`
    
    err := os.WriteFile("app.log", []byte(sampleLog), 0644)
    if err != nil {
        fmt.Printf("Error creating log file: %v\n", err)
        return
    }
    
    // TODO: Implement the log analyzer
    // 1. Read the log file
    // 2. Process each line
    // 3. Display the analysis results
    
    // Clean up
    os.Remove("app.log")
}
```

### 🎯 **Solution**

<details>
<summary>Click to reveal solution</summary>

```go
func (la *LogAnalyzer) AnalyzeLine(line string) {
    la.TotalLines++
    
    // Regex pattern for log format: "YYYY-MM-DD HH:MM:SS [LEVEL] Message"
    pattern := `(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\]`
    re := regexp.MustCompile(pattern)
    
    matches := re.FindStringSubmatch(line)
    if len(matches) < 3 {
        la.MalformedLines++
        return
    }
    
    // Parse timestamp
    timeStr := matches[1]
    timestamp, err := time.Parse("2006-01-02 15:04:05", timeStr)
    if err != nil {
        la.MalformedLines++
        return
    }
    
    // Update time range
    if la.FirstTime == nil || timestamp.Before(*la.FirstTime) {
        la.FirstTime = &timestamp
    }
    if la.LastTime == nil || timestamp.After(*la.LastTime) {
        la.LastTime = &timestamp
    }
    
    // Count log levels
    level := strings.ToUpper(matches[2])
    switch level {
    case "INFO":
        la.InfoCount++
    case "WARN":
        la.WarnCount++
    case "ERROR":
        la.ErrorCount++
    }
}

// Complete main function
func main() {
    // ... (sample log creation code remains the same)
    
    analyzer := &LogAnalyzer{}
    
    // Read and analyze the log file
    content, err := os.ReadFile("app.log")
    if err != nil {
        fmt.Printf("Error reading log file: %v\n", err)
        return
    }
    
    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line != "" {
            analyzer.AnalyzeLine(line)
        }
    }
    
    analyzer.PrintSummary()
    
    // Clean up
    os.Remove("app.log")
}
```

</details>

---

## 🎯 Key Takeaways

<div style="background: linear-gradient(45deg, #f093fb 0%, #f5576c 100%); padding: 15px; border-radius: 8px; color: white;">

**🎯 File Reading Best Practices:**
- Use `os.ReadFile()` for small files that fit in memory
- Use `os.Open()` + chunked reading for large files
- Always use `defer file.Close()` for resource cleanup
- Handle errors explicitly and provide meaningful messages
- Consider using `io.Reader` interface for flexible, testable code
- Validate file paths and sizes for security

</div>

---

## 🚀 What's Next?

Now that you've mastered reading files, you're ready to learn about writing data to files! In the next lesson, you'll discover:

- Different file writing strategies
- Atomic writes and temporary files
- Working with file permissions and flags
- Building robust data persistence solutions

<div align="center">

[**👈 Back to Module 7 Overview**](./module-7-file-system-io.md) | [**👉 Next: Writing Files →**](./module-7-2-writing-files.md)

</div>
