# 🚀 Module 9: Tooling, Dependencies, and Testing

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master Go's Professional Development Toolkit**

*From Zero to Go Hero - Module 9*

</div>

---

## 📋 Module Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 4-6 hours    | Intermediate   | Modules 1-8, Basic Go syntax, Functions, Structs |

### 🎯 Learning Objectives
By the end of this module, you will:
- ✅ Master Go Modules for dependency management
- ✅ Write comprehensive unit tests using the testing package
- ✅ Create benchmarks to measure code performance
- ✅ Understand Go's toolchain and development workflow
- ✅ Build production-ready Go projects with proper testing

---

## 🛠️ Go's Professional Development Ecosystem

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 Why Professional Tooling Matters

Go's strength lies not just in its language features, but in its comprehensive toolchain that makes development, testing, and deployment seamless. This module teaches you the essential tools every Go developer needs to know.

</div>

### 🚄 **1. The Go Development Workflow**
```
┌─────────────────────────────────────┐
│  Professional Go Development        │
├─────────────────────────────────────┤
│  Dependencies  ████████████████ 95% │
│  Testing       ███████████████▌ 90% │
│  Benchmarking  ████████████▌   75% │
│  Documentation ████████▌       50% │
└─────────────────────────────────────┘
```

### 🧵 **2. Go Modules: Dependency Management**
<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Go Modules** - The modern way to manage dependencies in Go, replacing the old GOPATH approach.

```go
// go.mod file example
module github.com/yourname/awesome-project

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/stretchr/testify v1.8.4
)
```

</div>

### 🎯 **3. Testing Philosophy in Go**
- **Built-in Testing** - No external frameworks needed
- **Table-Driven Tests** - Test multiple scenarios efficiently
- **Benchmarking** - Measure performance scientifically
- **Coverage Analysis** - Ensure code quality

### 🏢 **4. Essential Go Tools**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📦 go mod</strong><br>
Dependency management and module operations
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🧪 go test</strong><br>
Run tests and benchmarks
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📊 go vet</strong><br>
Static analysis and code quality
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🎨 go fmt</strong><br>
Automatic code formatting
</div>

</div>

---

## 📦 Go Modules: Step-by-Step Guide

### **Understanding Go Modules**

Think of Go Modules like a **recipe book** for your project. Just as a recipe lists all ingredients and their amounts, `go.mod` lists all dependencies and their versions.

```
🏠 Your Go Project
├── 📋 go.mod (Recipe book - lists what you need)
├── 🔒 go.sum (Security checksums - ensures authenticity)
├── 📁 main.go (Your actual code)
└── 📁 other files...
```

#### **Step 1: Creating Your First Module**

Initialize a new module:

```bash
# Create a new directory
mkdir awesome-calculator
cd awesome-calculator

# Initialize Go module
go mod init github.com/yourname/awesome-calculator
```

This creates a `go.mod` file:

```go
module github.com/yourname/awesome-calculator

go 1.21
```

#### **Step 2: Adding Dependencies**

Let's add a popular JSON library:

```bash
# Add a dependency
go get github.com/tidwall/gjson
```

Your `go.mod` now looks like:

```go
module github.com/yourname/awesome-calculator

go 1.21

require github.com/tidwall/gjson v1.17.0

indirect (
    github.com/tidwall/match v1.1.1 // indirect
    github.com/tidwall/pretty v1.2.0 // indirect
)
```

#### **Step 3: Using Dependencies in Code**

Create `main.go`:

```go
package main

import (
    "fmt"
    "github.com/tidwall/gjson"
)

func main() {
    json := `{"name":"John","age":30,"city":"New York"}`
    name := gjson.Get(json, "name")
    fmt.Printf("Hello, %s!\n", name.String())
}
```

#### **Step 4: Module Management Commands**

<div style="background: #e3f2fd; border-radius: 8px; padding: 15px; margin: 15px 0;">

```bash
# Download dependencies
go mod download

# Remove unused dependencies
go mod tidy

# View dependency graph
go mod graph

# Update all dependencies
go get -u ./...

# Update specific dependency
go get -u github.com/tidwall/gjson
```

</div>

### **🔒 Understanding go.sum**

The `go.sum` file is like a **security guard** - it contains cryptographic checksums to ensure dependencies haven't been tampered with:

```
github.com/tidwall/gjson v1.17.0 h1:wlYEnwqAHgzmhNUFfw7Xalt2JzQvsMx2Se4PcoFCT/U=
github.com/tidwall/gjson v1.17.0/go.mod h1:/wbyibRr2FHMks5tjHJ5F8dMZh3AcwJEMf5vlfC0lxk=
```

---

## 🧪 Unit Testing: Building Reliable Code

### **Testing Philosophy**

<div style="background: linear-gradient(135deg, #ff9a9e 0%, #fecfef 100%); padding: 20px; border-radius: 10px; color: #333; margin: 20px 0;">

### 💡 Testing Analogy

Think of testing like **quality control in a factory**:
- Each function is a machine that transforms input to output
- Tests verify each machine works correctly
- Edge cases are like stress-testing under extreme conditions

</div>

#### **Step 1: Basic Test Structure**

Create `calculator.go`:

```go
package main

import "errors"

// Add performs addition of two integers
func Add(a, b int) int {
    return a + b
}

// Divide performs division with error handling
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// IsEven checks if a number is even
func IsEven(n int) bool {
    return n%2 == 0
}
```

#### **Step 2: Writing Your First Tests**

Create `calculator_test.go`:

```go
package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestDivide(t *testing.T) {
    result, err := Divide(10, 2)
    if err != nil {
        t.Errorf("Divide(10, 2) returned error: %v", err)
    }
    
    expected := 5.0
    if result != expected {
        t.Errorf("Divide(10, 2) = %f; want %f", result, expected)
    }
}

func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) should return an error")
    }
}
```

#### **Step 3: Table-Driven Tests**

<div style="background: #f8f9fa; border-left: 4px solid #007bff; padding: 15px; margin: 10px 0;">

**Table-Driven Tests** - Test multiple scenarios efficiently with a single test function.

</div>

```go
func TestIsEven(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected bool
    }{
        {"positive even", 4, true},
        {"positive odd", 5, false},
        {"zero", 0, true},
        {"negative even", -4, true},
        {"negative odd", -3, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsEven(tt.input)
            if result != tt.expected {
                t.Errorf("IsEven(%d) = %v; want %v", tt.input, result, tt.expected)
            }
        })
    }
}
```

#### **Step 4: Running Tests**

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run specific test
go test -run TestAdd

# Run tests with coverage
go test -cover

# Generate detailed coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### **🎯 Advanced Testing Patterns**

#### **Setup and Teardown**

```go
func TestMain(m *testing.M) {
    // Setup code here
    fmt.Println("Setting up tests...")
    
    // Run tests
    code := m.Run()
    
    // Teardown code here
    fmt.Println("Cleaning up...")
    
    os.Exit(code)
}
```

#### **Testing with External Dependencies**

```go
// Using testify for assertions
import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestWithTestify(t *testing.T) {
    result := Add(2, 3)
    
    // More readable assertions
    assert.Equal(t, 5, result)
    require.NotNil(t, result) // Stops test if fails
}
```

---

## 📊 Benchmarking: Measuring Performance

### **Understanding Benchmarks**

<div style="background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%); padding: 20px; border-radius: 10px; color: #333; margin: 20px 0;">

### ⚡ Benchmarking Analogy

Benchmarking is like **timing a race**:
- Measure how fast your code runs
- Compare different implementations
- Find performance bottlenecks
- Make data-driven optimizations

</div>

#### **Step 1: Basic Benchmarks**

Create `benchmark_test.go`:

```go
package main

import (
    "testing"
    "strings"
)

// Two different ways to concatenate strings
func ConcatWithPlus(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s
    }
    return result
}

func ConcatWithBuilder(strs []string) string {
    var builder strings.Builder
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}

// Benchmark functions must start with "Benchmark"
func BenchmarkConcatWithPlus(b *testing.B) {
    strs := []string{"Hello", " ", "World", "!", " ", "Go", " ", "is", " ", "awesome"}
    
    for i := 0; i < b.N; i++ {
        ConcatWithPlus(strs)
    }
}

func BenchmarkConcatWithBuilder(b *testing.B) {
    strs := []string{"Hello", " ", "World", "!", " ", "Go", " ", "is", " ", "awesome"}
    
    for i := 0; i < b.N; i++ {
        ConcatWithBuilder(strs)
    }
}
```

#### **Step 2: Running Benchmarks**

```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkConcatWithPlus

# Run benchmarks with memory allocation stats
go test -bench=. -benchmem

# Run benchmarks multiple times for accuracy
go test -bench=. -count=5
```

Sample output:
```
BenchmarkConcatWithPlus-8      1000000    1234 ns/op    456 B/op    12 allocs/op
BenchmarkConcatWithBuilder-8   5000000     234 ns/op     64 B/op     1 allocs/op
```

#### **Step 3: Advanced Benchmarking**

```go
func BenchmarkWithSetup(b *testing.B) {
    // Setup code (not timed)
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    
    // Reset timer to exclude setup time
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        // Only this code is timed
        processData(data)
    }
}

func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // Code to benchmark in parallel
            expensiveOperation()
        }
    })
}
```

### **📈 Interpreting Benchmark Results**

```
┌─────────────────────────────────────────────────────────┐
│  Benchmark Result Breakdown                             │
├─────────────────────────────────────────────────────────┤
│  BenchmarkFunction-8    1000000    1234 ns/op    456 B/op│
│  │                 │         │         │            │    │
│  │                 │         │         │            └─── Bytes allocated per operation
│  │                 │         │         └──────────────── Nanoseconds per operation
│  │                 │         └────────────────────────── Number of iterations
│  │                 └──────────────────────────────────── Number of CPU cores used
│  └────────────────────────────────────────────────────── Function name
└─────────────────────────────────────────────────────────┘
```

---

## 🛠️ Practical Project: Building a Tested Library

Let's create a complete project that demonstrates all concepts:

#### **Step 1: Project Structure**

```
math-utils/
├── go.mod
├── go.sum
├── math.go
├── math_test.go
├── benchmark_test.go
└── README.md
```

#### **Step 2: Initialize Module**

```bash
mkdir math-utils
cd math-utils
go mod init github.com/yourname/math-utils
```

#### **Step 3: Core Library (math.go)**

```go
package mathutils

import (
    "errors"
    "math"
)

// Statistics holds statistical data
type Statistics struct {
    Mean   float64
    Median float64
    StdDev float64
}

// Factorial calculates factorial of n
func Factorial(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("factorial of negative number is undefined")
    }
    if n == 0 || n == 1 {
        return 1, nil
    }
    
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result, nil
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    
    sqrt := int(math.Sqrt(float64(n)))
    for i := 3; i <= sqrt; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}

// CalculateStats computes basic statistics
func CalculateStats(numbers []float64) (*Statistics, error) {
    if len(numbers) == 0 {
        return nil, errors.New("cannot calculate statistics for empty slice")
    }
    
    // Calculate mean
    sum := 0.0
    for _, num := range numbers {
        sum += num
    }
    mean := sum / float64(len(numbers))
    
    // Calculate standard deviation
    variance := 0.0
    for _, num := range numbers {
        variance += math.Pow(num-mean, 2)
    }
    stdDev := math.Sqrt(variance / float64(len(numbers)))
    
    return &Statistics{
        Mean:   mean,
        StdDev: stdDev,
    }, nil
}
```

#### **Step 4: Comprehensive Tests (math_test.go)**

```go
package mathutils

import (
    "testing"
    "math"
)

func TestFactorial(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
        hasError bool
    }{
        {"zero", 0, 1, false},
        {"one", 1, 1, false},
        {"small positive", 5, 120, false},
        {"negative", -1, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Factorial(tt.input)
            
            if tt.hasError {
                if err == nil {
                    t.Errorf("Factorial(%d) should return error", tt.input)
                }
                return
            }
            
            if err != nil {
                t.Errorf("Factorial(%d) returned unexpected error: %v", tt.input, err)
                return
            }
            
            if result != tt.expected {
                t.Errorf("Factorial(%d) = %d; want %d", tt.input, result, tt.expected)
            }
        })
    }
}

func TestIsPrime(t *testing.T) {
    tests := []struct {
        input    int
        expected bool
    }{
        {-1, false},
        {0, false},
        {1, false},
        {2, true},
        {3, true},
        {4, false},
        {17, true},
        {25, false},
        {97, true},
    }

    for _, tt := range tests {
        t.Run(fmt.Sprintf("IsPrime(%d)", tt.input), func(t *testing.T) {
            result := IsPrime(tt.input)
            if result != tt.expected {
                t.Errorf("IsPrime(%d) = %v; want %v", tt.input, result, tt.expected)
            }
        })
    }
}

func TestCalculateStats(t *testing.T) {
    t.Run("normal case", func(t *testing.T) {
        numbers := []float64{1, 2, 3, 4, 5}
        stats, err := CalculateStats(numbers)
        
        if err != nil {
            t.Errorf("CalculateStats returned error: %v", err)
            return
        }
        
        expectedMean := 3.0
        if math.Abs(stats.Mean-expectedMean) > 0.0001 {
            t.Errorf("Mean = %f; want %f", stats.Mean, expectedMean)
        }
    })
    
    t.Run("empty slice", func(t *testing.T) {
        _, err := CalculateStats([]float64{})
        if err == nil {
            t.Error("CalculateStats should return error for empty slice")
        }
    })
}
```

#### **Step 5: Performance Benchmarks (benchmark_test.go)**

```go
package mathutils

import (
    "testing"
    "math/rand"
)

func BenchmarkFactorial(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Factorial(10)
    }
}

func BenchmarkIsPrime(b *testing.B) {
    numbers := make([]int, 1000)
    for i := range numbers {
        numbers[i] = rand.Intn(10000) + 1
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, num := range numbers {
            IsPrime(num)
        }
    }
}

func BenchmarkCalculateStats(b *testing.B) {
    numbers := make([]float64, 10000)
    for i := range numbers {
        numbers[i] = rand.Float64() * 1000
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        CalculateStats(numbers)
    }
}
```

---

## 🎯 Practice Exercises

<div style="background: #fff3cd; border: 1px solid #ffeaa7; border-radius: 8px; padding: 20px; margin: 20px 0;">

### 🏋️ Exercise 1: Module Management
1. Create a new module for a "todo-app"
2. Add dependencies: `github.com/google/uuid` and `github.com/stretchr/testify`
3. Create a simple Todo struct and basic operations
4. Practice `go mod tidy` and `go mod graph`

### 🧪 Exercise 2: Comprehensive Testing
1. Write tests for your Todo operations
2. Use table-driven tests for validation functions
3. Test error conditions and edge cases
4. Achieve 90%+ test coverage

### 📊 Exercise 3: Performance Analysis
1. Implement two different sorting algorithms
2. Benchmark both implementations
3. Compare performance with different data sizes
4. Analyze memory allocation patterns

</div>

---

## 🚀 Best Practices & Tips

### **Module Management**
- ✅ Use semantic versioning (v1.2.3)
- ✅ Keep dependencies minimal and up-to-date
- ✅ Use `go mod tidy` regularly
- ✅ Commit both `go.mod` and `go.sum`

### **Testing**
- ✅ Write tests before or alongside code (TDD)
- ✅ Test both happy path and error cases
- ✅ Use descriptive test names
- ✅ Aim for high test coverage (80%+)

### **Benchmarking**
- ✅ Benchmark on target hardware
- ✅ Run multiple iterations for accuracy
- ✅ Profile memory allocations
- ✅ Compare before and after optimizations

---

## 🎉 Module Summary

Congratulations! You've mastered Go's professional development toolkit:

### **What You've Learned:**
- 📦 **Go Modules**: Modern dependency management
- 🧪 **Unit Testing**: Writing reliable, maintainable tests
- 📊 **Benchmarking**: Measuring and optimizing performance
- 🛠️ **Toolchain**: Professional development workflow

### **Key Takeaways:**
- Go's toolchain makes professional development seamless
- Testing is built into the language and culture
- Benchmarking enables data-driven optimization
- Modules provide reliable dependency management

### **Next Steps:**
You're now ready to build production-grade Go applications with confidence. The next modules will cover web development, databases, and deployment - all building on the solid foundation you've established here.

---

<div align="center">

**🎯 Ready for Module 10? Let's build web services!**

[← Previous Module](module-8-data-processing-serialization.md) | [Next Module →](module-10-client-side-networking.md)

</div>
