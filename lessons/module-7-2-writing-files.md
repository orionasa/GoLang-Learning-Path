# ✏️ Module 7.2: Writing Data to Files

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master File Writing in Go**

*From Zero to Go Hero - Module 7.2*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1-1.5 hours  | Intermediate   | Module 7.1 (Reading Files), Basic Go syntax |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Create and write files using different Go approaches
- ✅ Understand file creation flags and permissions
- ✅ Implement atomic writes and temporary files for data safety
- ✅ Work with the `io.Writer` interface for flexible output
- ✅ Build robust data persistence solutions

---

## 🌟 Understanding File Writing in Go

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎭 The File Writing Analogy

Think of writing files like different ways of writing in a notebook:
- **os.WriteFile()**: Like writing the entire page at once with a marker
- **os.Create() + Write()**: Like opening a fresh notebook and writing line by line
- **os.OpenFile()**: Like choosing exactly how to open your notebook (append, overwrite, etc.)
- **Atomic writes**: Like writing on a separate paper first, then replacing the original

</div>

### 🚄 **1. File Writing Methods Comparison**

```
┌─────────────────────────────────────────────────────────┐
│  ✏️ File Writing Methods Performance & Safety           │
├─────────────────────────────────────────────────────────┤
│  Simple Write   os.WriteFile()    ████████████████ 95% │
│  Controlled     os.Create()       ███████████████▌ 88% │
│  Flexible       os.OpenFile()     ████████████▌   82% │
│  Atomic         Temp+Rename       ██████████████▌  85% │
└─────────────────────────────────────────────────────────┘
```

---

## 📝 Method 1: Simple Writing with os.WriteFile()

### 🎯 **When to Use**
- Small to medium files
- One-shot write operations
- When you have all data ready in memory

### 💡 **Basic Implementation**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Simple data to write
    data := []byte("Hello, Go file writing!\nThis is line 2.\nThis is line 3.")
    
    // Write data to file (creates file if it doesn't exist)
    err := os.WriteFile("simple_output.txt", data, 0644)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }
    
    fmt.Println("✅ File written successfully!")
    
    // Verify by reading back
    content, err := os.ReadFile("simple_output.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("📄 File content:\n%s\n", string(content))
    
    // Clean up
    os.Remove("simple_output.txt")
}
```

### 🔧 **Enhanced Example: Configuration Writer**

```go
package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

type AppConfig struct {
    ServerPort    int
    DatabaseURL   string
    DebugMode     bool
    MaxConnections int
    CreatedAt     time.Time
}

func (config *AppConfig) ToConfigFile() string {
    var builder strings.Builder
    
    // Add header comment
    builder.WriteString("# Application Configuration\n")
    builder.WriteString(fmt.Sprintf("# Generated on: %s\n\n", config.CreatedAt.Format("2006-01-02 15:04:05")))
    
    // Add configuration values
    builder.WriteString(fmt.Sprintf("server_port=%d\n", config.ServerPort))
    builder.WriteString(fmt.Sprintf("database_url=%s\n", config.DatabaseURL))
    builder.WriteString(fmt.Sprintf("debug_mode=%t\n", config.DebugMode))
    builder.WriteString(fmt.Sprintf("max_connections=%d\n", config.MaxConnections))
    
    return builder.String()
}

func saveConfig(config *AppConfig, filename string) error {
    configData := config.ToConfigFile()
    
    // Write configuration to file with appropriate permissions
    err := os.WriteFile(filename, []byte(configData), 0644)
    if err != nil {
        return fmt.Errorf("failed to save config: %w", err)
    }
    
    fmt.Printf("✅ Configuration saved to %s\n", filename)
    return nil
}

func main() {
    // Create sample configuration
    config := &AppConfig{
        ServerPort:     8080,
        DatabaseURL:    "postgresql://localhost:5432/myapp",
        DebugMode:      true,
        MaxConnections: 100,
        CreatedAt:      time.Now(),
    }
    
    // Save configuration
    err := saveConfig(config, "app.config")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    // Display the generated file
    content, _ := os.ReadFile("app.config")
    fmt.Printf("\n📋 Generated configuration:\n%s", string(content))
    
    // Clean up
    os.Remove("app.config")
}
```

---

## 🔧 Method 2: Controlled Writing with os.Create() and os.OpenFile()

### 🎯 **When to Use**
- Need more control over the writing process
- Writing large amounts of data in chunks
- Appending to existing files
- Setting specific file permissions

### 💡 **File Creation and Writing**

```go
package main

import (
    "fmt"
    "os"
)

func writeInChunks(filename string, data []string) error {
    // Create a new file (truncates if exists)
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()
    
    // Write data in chunks
    for i, chunk := range data {
        bytesWritten, err := file.WriteString(chunk + "\n")
        if err != nil {
            return fmt.Errorf("failed to write chunk %d: %w", i, err)
        }
        fmt.Printf("📦 Chunk %d: Wrote %d bytes\n", i+1, bytesWritten)
    }
    
    // Ensure data is written to disk
    err = file.Sync()
    if err != nil {
        return fmt.Errorf("failed to sync file: %w", err)
    }
    
    return nil
}

func main() {
    // Sample data chunks
    dataChunks := []string{
        "First chunk of data",
        "Second chunk with more information",
        "Third chunk containing final details",
        "Fourth chunk to complete the file",
    }
    
    fmt.Println("🚀 Writing file in chunks...")
    err := writeInChunks("chunked_output.txt", dataChunks)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Println("✅ File written successfully!")
    
    // Verify the result
    content, _ := os.ReadFile("chunked_output.txt")
    fmt.Printf("\n📄 Final file content:\n%s", string(content))
    
    // Clean up
    os.Remove("chunked_output.txt")
}
```

### 🔧 **Advanced: File Flags and Permissions**

```go
package main

import (
    "fmt"
    "os"
    "time"
)

type LogLevel int

const (
    INFO LogLevel = iota
    WARN
    ERROR
)

func (l LogLevel) String() string {
    switch l {
    case INFO:
        return "INFO"
    case WARN:
        return "WARN"
    case ERROR:
        return "ERROR"
    default:
        return "UNKNOWN"
    }
}

type Logger struct {
    filename string
    file     *os.File
}

func NewLogger(filename string) (*Logger, error) {
    // Open file with specific flags:
    // O_CREATE: Create if doesn't exist
    // O_WRONLY: Write-only
    // O_APPEND: Append to end of file
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return nil, fmt.Errorf("failed to open log file: %w", err)
    }
    
    return &Logger{
        filename: filename,
        file:     file,
    }, nil
}

func (l *Logger) Log(level LogLevel, message string) error {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    logEntry := fmt.Sprintf("%s [%s] %s\n", timestamp, level.String(), message)
    
    _, err := l.file.WriteString(logEntry)
    if err != nil {
        return fmt.Errorf("failed to write log entry: %w", err)
    }
    
    // Force write to disk for important logs
    if level == ERROR {
        l.file.Sync()
    }
    
    return nil
}

func (l *Logger) Close() error {
    if l.file != nil {
        return l.file.Close()
    }
    return nil
}

func main() {
    // Create logger
    logger, err := NewLogger("app.log")
    if err != nil {
        fmt.Printf("Error creating logger: %v\n", err)
        return
    }
    defer logger.Close()
    
    // Log different types of messages
    logger.Log(INFO, "Application started")
    logger.Log(INFO, "Database connection established")
    logger.Log(WARN, "High memory usage detected")
    logger.Log(ERROR, "Failed to process request")
    logger.Log(INFO, "Application shutdown")
    
    fmt.Println("✅ Logs written successfully!")
    
    // Display the log file
    content, _ := os.ReadFile("app.log")
    fmt.Printf("\n📋 Log file content:\n%s", string(content))
    
    // Clean up
    os.Remove("app.log")
}
```

---

## 🛡️ Method 3: Atomic Writes for Data Safety

### 🎯 **When to Use**
- Critical data that must not be corrupted
- Configuration files that other processes read
- When you need all-or-nothing write semantics

### 💡 **Atomic Write Implementation**

```go
package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

type AtomicWriter struct {
    targetPath string
    tempPath   string
    tempFile   *os.File
}

func NewAtomicWriter(targetPath string) (*AtomicWriter, error) {
    // Create temporary file in the same directory
    dir := filepath.Dir(targetPath)
    tempFile, err := os.CreateTemp(dir, "atomic_write_*.tmp")
    if err != nil {
        return nil, fmt.Errorf("failed to create temp file: %w", err)
    }
    
    return &AtomicWriter{
        targetPath: targetPath,
        tempPath:   tempFile.Name(),
        tempFile:   tempFile,
    }, nil
}

func (aw *AtomicWriter) Write(data []byte) (int, error) {
    return aw.tempFile.Write(data)
}

func (aw *AtomicWriter) WriteString(s string) (int, error) {
    return aw.tempFile.WriteString(s)
}

func (aw *AtomicWriter) Commit() error {
    // Close the temporary file
    if err := aw.tempFile.Close(); err != nil {
        return fmt.Errorf("failed to close temp file: %w", err)
    }
    
    // Atomically replace the target file
    if err := os.Rename(aw.tempPath, aw.targetPath); err != nil {
        // Clean up temp file on failure
        os.Remove(aw.tempPath)
        return fmt.Errorf("failed to commit atomic write: %w", err)
    }
    
    return nil
}

func (aw *AtomicWriter) Abort() error {
    // Close and remove the temporary file
    if aw.tempFile != nil {
        aw.tempFile.Close()
    }
    return os.Remove(aw.tempPath)
}

// Safe configuration updater
func updateConfigSafely(filename string, newConfig map[string]string) error {
    writer, err := NewAtomicWriter(filename)
    if err != nil {
        return err
    }
    
    // Write header
    writer.WriteString("# Application Configuration (Updated)\n")
    writer.WriteString(fmt.Sprintf("# Last updated: %s\n\n", 
        time.Now().Format("2006-01-02 15:04:05")))
    
    // Write configuration
    for key, value := range newConfig {
        _, err := writer.WriteString(fmt.Sprintf("%s=%s\n", key, value))
        if err != nil {
            writer.Abort() // Clean up on error
            return fmt.Errorf("failed to write config entry: %w", err)
        }
    }
    
    // Commit the changes atomically
    return writer.Commit()
}

func main() {
    // Create initial config file
    initialConfig := "# Initial Configuration\nserver_port=8080\ndebug=false\n"
    os.WriteFile("config.txt", []byte(initialConfig), 0644)
    
    fmt.Println("📋 Initial config:")
    content, _ := os.ReadFile("config.txt")
    fmt.Printf("%s\n", string(content))
    
    // Update configuration safely
    newConfig := map[string]string{
        "server_port":     "9090",
        "debug":          "true",
        "max_connections": "200",
        "database_url":   "postgresql://localhost:5432/newdb",
    }
    
    fmt.Println("🔄 Updating configuration atomically...")
    err := updateConfigSafely("config.txt", newConfig)
    if err != nil {
        fmt.Printf("Error updating config: %v\n", err)
        return
    }
    
    fmt.Println("✅ Configuration updated successfully!")
    fmt.Println("\n📋 Updated config:")
    content, _ = os.ReadFile("config.txt")
    fmt.Printf("%s", string(content))
    
    // Clean up
    os.Remove("config.txt")
}
```

---

## 🔄 Method 4: Working with io.Writer Interface

### 🎯 **When to Use**
- Building flexible, testable code
- Supporting multiple output destinations
- Creating reusable writing components

### 💡 **Interface-Based Writing**

```go
package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "time"
)

// Report generator that can write to any io.Writer
type ReportGenerator struct {
    Title     string
    Timestamp time.Time
}

func NewReportGenerator(title string) *ReportGenerator {
    return &ReportGenerator{
        Title:     title,
        Timestamp: time.Now(),
    }
}

func (rg *ReportGenerator) GenerateReport(writer io.Writer, data map[string]interface{}) error {
    // Write report header
    _, err := fmt.Fprintf(writer, "# %s\n", rg.Title)
    if err != nil {
        return err
    }
    
    _, err = fmt.Fprintf(writer, "Generated: %s\n\n", rg.Timestamp.Format("2006-01-02 15:04:05"))
    if err != nil {
        return err
    }
    
    // Write report sections
    _, err = fmt.Fprintf(writer, "## Summary\n")
    if err != nil {
        return err
    }
    
    for key, value := range data {
        _, err = fmt.Fprintf(writer, "- **%s**: %v\n", key, value)
        if err != nil {
            return err
        }
    }
    
    // Write footer
    _, err = fmt.Fprintf(writer, "\n---\n*Report generated by Go Application*\n")
    return err
}

// Multi-destination writer
type MultiWriter struct {
    writers []io.Writer
}

func NewMultiWriter(writers ...io.Writer) *MultiWriter {
    return &MultiWriter{writers: writers}
}

func (mw *MultiWriter) Write(p []byte) (n int, err error) {
    for _, writer := range mw.writers {
        n, err = writer.Write(p)
        if err != nil {
            return n, err
        }
    }
    return len(p), nil
}

func main() {
    // Create sample data
    reportData := map[string]interface{}{
        "Total Users":        1250,
        "Active Sessions":    89,
        "Server Uptime":     "72 hours",
        "Memory Usage":      "45%",
        "Disk Space":       "23.5 GB available",
    }
    
    // Create report generator
    generator := NewReportGenerator("System Status Report")
    
    // Example 1: Write to file
    fmt.Println("📄 Writing report to file...")
    file, err := os.Create("system_report.md")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    err = generator.GenerateReport(file, reportData)
    if err != nil {
        fmt.Printf("Error writing to file: %v\n", err)
        return
    }
    
    // Example 2: Write to string buffer (for testing)
    fmt.Println("📝 Writing report to string buffer...")
    var buffer strings.Builder
    err = generator.GenerateReport(&buffer, reportData)
    if err != nil {
        fmt.Printf("Error writing to buffer: %v\n", err)
        return
    }
    
    fmt.Println("✅ Report generated in memory!")
    
    // Example 3: Write to multiple destinations simultaneously
    fmt.Println("📊 Writing report to multiple destinations...")
    
    logFile, err := os.Create("report.log")
    if err != nil {
        fmt.Printf("Error creating log file: %v\n", err)
        return
    }
    defer logFile.Close()
    
    // Create multi-writer that writes to both stdout and log file
    multiWriter := NewMultiWriter(os.Stdout, logFile)
    
    fmt.Println("\n--- Report Output ---")
    err = generator.GenerateReport(multiWriter, reportData)
    if err != nil {
        fmt.Printf("Error writing to multiple destinations: %v\n", err)
        return
    }
    
    // Display file contents
    fmt.Println("\n📋 File content verification:")
    content, _ := os.ReadFile("system_report.md")
    fmt.Printf("File size: %d bytes\n", len(content))
    
    // Clean up
    os.Remove("system_report.md")
    os.Remove("report.log")
}
```

---

## 🧪 Hands-On Exercise: Data Export System

### 🎯 **Challenge**
Build a flexible data export system that can:
1. Export data in multiple formats (CSV, JSON, plain text)
2. Support both file and string output
3. Handle large datasets efficiently
4. Provide atomic write guarantees

### 💡 **Starter Code**

```go
package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "os"
    "strconv"
    "time"
)

type User struct {
    ID       int       `json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    JoinDate time.Time `json:"join_date"`
    Active   bool      `json:"active"`
}

type DataExporter struct {
    // TODO: Add fields for configuration
}

func NewDataExporter() *DataExporter {
    // TODO: Initialize exporter
    return &DataExporter{}
}

func (de *DataExporter) ExportCSV(writer io.Writer, users []User) error {
    // TODO: Implement CSV export
    // Hint: Use encoding/csv package
    // Include headers: ID, Name, Email, Join Date, Active
    return nil
}

func (de *DataExporter) ExportJSON(writer io.Writer, users []User) error {
    // TODO: Implement JSON export
    // Hint: Use encoding/json package with proper formatting
    return nil
}

func (de *DataExporter) ExportText(writer io.Writer, users []User) error {
    // TODO: Implement plain text export
    // Format: "ID: 1, Name: John Doe, Email: john@example.com, ..."
    return nil
}

func (de *DataExporter) ExportToFileAtomic(filename string, format string, users []User) error {
    // TODO: Implement atomic file export
    // Use temporary file + rename pattern
    return nil
}

func main() {
    // Sample data
    users := []User{
        {1, "John Doe", "john@example.com", time.Now().AddDate(-1, 0, 0), true},
        {2, "Jane Smith", "jane@example.com", time.Now().AddDate(0, -6, 0), true},
        {3, "Bob Johnson", "bob@example.com", time.Now().AddDate(0, -3, 0), false},
        {4, "Alice Brown", "alice@example.com", time.Now().AddDate(0, -1, 0), true},
    }
    
    exporter := NewDataExporter()
    
    // TODO: Test all export formats
    // 1. Export to CSV file
    // 2. Export to JSON file  
    // 3. Export to text file
    // 4. Test atomic writes
    // 5. Verify file contents
    
    fmt.Println("🚀 Data export system ready for implementation!")
}
```

### 🎯 **Solution**

<details>
<summary>Click to reveal solution</summary>

```go
func (de *DataExporter) ExportCSV(writer io.Writer, users []User) error {
    csvWriter := csv.NewWriter(writer)
    defer csvWriter.Flush()
    
    // Write header
    header := []string{"ID", "Name", "Email", "Join Date", "Active"}
    if err := csvWriter.Write(header); err != nil {
        return err
    }
    
    // Write data
    for _, user := range users {
        record := []string{
            strconv.Itoa(user.ID),
            user.Name,
            user.Email,
            user.JoinDate.Format("2006-01-02"),
            strconv.FormatBool(user.Active),
        }
        if err := csvWriter.Write(record); err != nil {
            return err
        }
    }
    
    return nil
}

func (de *DataExporter) ExportJSON(writer io.Writer, users []User) error {
    encoder := json.NewEncoder(writer)
    encoder.SetIndent("", "  ")
    return encoder.Encode(users)
}

func (de *DataExporter) ExportText(writer io.Writer, users []User) error {
    for _, user := range users {
        line := fmt.Sprintf("ID: %d, Name: %s, Email: %s, Join Date: %s, Active: %t\n",
            user.ID, user.Name, user.Email, 
            user.JoinDate.Format("2006-01-02"), user.Active)
        
        if _, err := writer.Write([]byte(line)); err != nil {
            return err
        }
    }
    return nil
}

func (de *DataExporter) ExportToFileAtomic(filename string, format string, users []User) error {
    atomicWriter, err := NewAtomicWriter(filename)
    if err != nil {
        return err
    }
    
    var exportErr error
    switch format {
    case "csv":
        exportErr = de.ExportCSV(atomicWriter, users)
    case "json":
        exportErr = de.ExportJSON(atomicWriter, users)
    case "text":
        exportErr = de.ExportText(atomicWriter, users)
    default:
        exportErr = fmt.Errorf("unsupported format: %s", format)
    }
    
    if exportErr != nil {
        atomicWriter.Abort()
        return exportErr
    }
    
    return atomicWriter.Commit()
}
```

</details>

---

## 🎯 Key Takeaways

<div style="background: linear-gradient(45deg, #f093fb 0%, #f5576c 100%); padding: 15px; border-radius: 8px; color: white;">

**🎯 File Writing Best Practices:**
- Use `os.WriteFile()` for simple, small file operations
- Use `os.Create()` or `os.OpenFile()` for more control
- Always handle errors and use `defer file.Close()`
- Consider atomic writes for critical data
- Use `io.Writer` interface for flexible, testable code
- Set appropriate file permissions (0644 for regular files)
- Use `file.Sync()` to ensure data reaches disk

</div>

---

## 🚀 What's Next?

Excellent! You've mastered file writing techniques. Next, you'll learn about buffered I/O operations that can dramatically improve performance when working with large files or frequent I/O operations.

<div align="center">

[**👈 Back to Module 7 Overview**](./module-7-file-system-io.md) | [**👉 Next: Buffered I/O →**](./module-7-3-buffered-io.md)

</div>
