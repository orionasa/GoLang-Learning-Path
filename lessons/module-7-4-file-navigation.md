# 🗂️ Module 7.4: Navigating the File System

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master File System Navigation in Go**

*From Zero to Go Hero - Module 7.4*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1 hour       | Intermediate   | Modules 7.1-7.3, Basic understanding of file systems |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Navigate directory structures with `filepath.Walk()` and `os.ReadDir()`
- ✅ Gather file and directory information efficiently
- ✅ Manipulate and validate file paths safely
- ✅ Handle cross-platform path compatibility
- ✅ Build practical file system utilities

---

## 🌟 Understanding File System Navigation

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎭 The File System Analogy

Think of file system navigation like exploring a library:
- **filepath.Walk()**: Like having a guide who shows you every book in every section
- **os.ReadDir()**: Like looking at just one shelf at a time
- **filepath.Join()**: Like having a librarian who knows exactly how to write addresses

</div>

---

## 📁 Directory Listing with os.ReadDir()

### 💡 **Basic Directory Reading**

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func exploreDirectory(dirPath string) error {
    fmt.Printf("📁 Exploring directory: %s\n", dirPath)
    
    entries, err := os.ReadDir(dirPath)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }
    
    fmt.Printf("Found %d items:\n\n", len(entries))
    
    var totalSize int64
    fileCount := 0
    dirCount := 0
    
    for _, entry := range entries {
        info, err := entry.Info()
        if err != nil {
            continue
        }
        
        var icon string
        if entry.IsDir() {
            icon = "📁"
            dirCount++
        } else {
            icon = "📄"
            fileCount++
            totalSize += info.Size()
        }
        
        fmt.Printf("%s %-30s %10s %s\n", 
            icon, 
            entry.Name(), 
            formatFileSize(info.Size()),
            info.ModTime().Format("2006-01-02 15:04"))
    }
    
    fmt.Printf("\n📊 Summary:\n")
    fmt.Printf("  Directories: %d\n", dirCount)
    fmt.Printf("  Files: %d\n", fileCount)
    fmt.Printf("  Total size: %s\n", formatFileSize(totalSize))
    
    return nil
}

func formatFileSize(size int64) string {
    const unit = 1024
    if size < unit {
        return fmt.Sprintf("%d B", size)
    }
    div, exp := int64(unit), 0
    for n := size / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func main() {
    testDir := "sample_directory"
    os.MkdirAll(filepath.Join(testDir, "subdirectory"), 0755)
    
    files := []string{
        "readme.txt",
        "config.json", 
        "data.csv",
        "subdirectory/nested_file.txt",
    }
    
    for _, file := range files {
        fullPath := filepath.Join(testDir, file)
        content := fmt.Sprintf("Content for %s", file)
        os.WriteFile(fullPath, []byte(content), 0644)
    }
    
    exploreDirectory(testDir)
    os.RemoveAll(testDir)
}
```

---

## 🚶 Recursive Walking with filepath.Walk()

### 💡 **Basic Directory Walking**

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func walkDirectory(rootPath string) error {
    fmt.Printf("🚶 Walking directory tree: %s\n\n", rootPath)
    
    var totalFiles int
    var totalDirs int
    var totalSize int64
    
    err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Printf("❌ Error accessing %s: %v\n", path, err)
            return nil
        }
        
        relPath, _ := filepath.Rel(rootPath, path)
        if relPath == "." {
            relPath = "(root)"
        }
        
        depth := strings.Count(relPath, string(filepath.Separator))
        indent := strings.Repeat("  ", depth)
        
        if info.IsDir() {
            fmt.Printf("%s📁 %s/\n", indent, info.Name())
            totalDirs++
        } else {
            fmt.Printf("%s📄 %s (%s)\n", indent, info.Name(), formatFileSize(info.Size()))
            totalFiles++
            totalSize += info.Size()
        }
        
        return nil
    })
    
    if err != nil {
        return fmt.Errorf("error walking directory: %w", err)
    }
    
    fmt.Printf("\n📊 Walk Summary:\n")
    fmt.Printf("  Directories: %d\n", totalDirs)
    fmt.Printf("  Files: %d\n", totalFiles)
    fmt.Printf("  Total size: %s\n", formatFileSize(totalSize))
    
    return nil
}

func main() {
    testDir := "walk_test"
    dirs := []string{
        "walk_test/documents",
        "walk_test/documents/projects/go",
        "walk_test/media/images",
        "walk_test/config",
    }
    
    for _, dir := range dirs {
        os.MkdirAll(dir, 0755)
    }
    
    files := []string{
        "walk_test/readme.txt",
        "walk_test/documents/notes.md",
        "walk_test/documents/projects/go/main.go",
        "walk_test/media/images/photo1.jpg",
        "walk_test/config/app.json",
    }
    
    for _, file := range files {
        content := fmt.Sprintf("Content of %s", filepath.Base(file))
        os.WriteFile(file, []byte(content), 0644)
    }
    
    walkDirectory(testDir)
    os.RemoveAll(testDir)
}
```

---

## 🛠️ Path Manipulation and Validation

### 💡 **Safe Path Operations**

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

type PathUtils struct{}

func (pu *PathUtils) SafeJoin(base string, paths ...string) (string, error) {
    result := filepath.Clean(base)
    
    for _, path := range paths {
        cleanPath := filepath.Clean(path)
        
        if strings.Contains(cleanPath, "..") {
            return "", fmt.Errorf("path contains directory traversal: %s", path)
        }
        
        result = filepath.Join(result, cleanPath)
    }
    
    relPath, err := filepath.Rel(base, result)
    if err != nil {
        return "", fmt.Errorf("failed to calculate relative path: %w", err)
    }
    
    if strings.HasPrefix(relPath, "..") {
        return "", fmt.Errorf("resulting path escapes base directory")
    }
    
    return result, nil
}

func (pu *PathUtils) GetPathInfo(path string) map[string]interface{} {
    info := make(map[string]interface{})
    
    info["absolute"] = filepath.IsAbs(path)
    info["directory"] = filepath.Dir(path)
    info["filename"] = filepath.Base(path)
    info["extension"] = filepath.Ext(path)
    info["clean"] = filepath.Clean(path)
    
    if _, err := os.Stat(path); err == nil {
        info["exists"] = true
        if fileInfo, err := os.Stat(path); err == nil {
            info["is_directory"] = fileInfo.IsDir()
            info["size"] = fileInfo.Size()
            info["modified"] = fileInfo.ModTime()
        }
    } else {
        info["exists"] = false
    }
    
    return info
}

func main() {
    utils := &PathUtils{}
    
    testPaths := []string{
        "documents/file.txt",
        "/absolute/path/file.txt",
        "../../../etc/passwd",
        "normal/path/to/file.go",
    }
    
    for _, testPath := range testPaths {
        fmt.Printf("Testing path: %q\n", testPath)
        
        info := utils.GetPathInfo(testPath)
        fmt.Printf("  📁 Directory: %s\n", info["directory"])
        fmt.Printf("  📄 Filename: %s\n", info["filename"])
        fmt.Printf("  🏷️ Extension: %s\n", info["extension"])
        fmt.Printf("  📍 Absolute: %v\n", info["absolute"])
        fmt.Printf("  🔍 Exists: %v\n", info["exists"])
        
        if testPath != "" {
            safeJoined, err := utils.SafeJoin("/safe/base", testPath)
            if err != nil {
                fmt.Printf("  ⚠️ Safe join failed: %v\n", err)
            } else {
                fmt.Printf("  🔒 Safe join: %s\n", safeJoined)
            }
        }
        fmt.Println()
    }
}
```

---

## 🎯 Key Takeaways

<div style="background: linear-gradient(45deg, #f093fb 0%, #f5576c 100%); padding: 15px; border-radius: 8px; color: white;">

**🎯 File Navigation Best Practices:**
- Use `os.ReadDir()` for single directory exploration
- Use `filepath.Walk()` for recursive directory traversal
- Always validate and clean file paths for security
- Handle cross-platform path differences with `filepath` package
- Use relative paths when possible for portability

</div>

---

## 🚀 What's Next?

Great! You've learned how to navigate file systems safely and efficiently. Next, you'll discover performance optimization techniques for handling large files and implementing efficient data processing patterns.

<div align="center">

[**👈 Back to Module 7 Overview**](./module-7-file-system-io.md) | [**👉 Next: Performance Tuning →**](./module-7-5-performance-tuning.md)

</div>
