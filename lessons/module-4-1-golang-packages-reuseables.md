# 🚀 Module 4.1: Go Packages & Reusable Code

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Building Modular and Reusable Go Applications**

*From Zero to Go Hero - Module 4.1*

</div>

---

## 📋 Module Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 2-3 hours    | Beginner       | Basic Go syntax, functions, variables |

### 🎯 Learning Objectives
By the end of this module, you will:
- ✅ Understand Go's package system and organization
- ✅ Create your own packages and make code reusable
- ✅ Master import statements and package visibility
- ✅ Learn Go's naming conventions and best practices
- ✅ Build a practical multi-package project

---

## 📦 Understanding Go Packages

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 What Are Packages?

Think of packages like **toolboxes** 🧰. Each toolbox contains related tools (functions, types, variables) that work together to solve specific problems. Just like you wouldn't mix your kitchen tools with your garage tools, Go packages help organize related code together.

</div>

### 🚄 **1. Package Basics**

```
┌─────────────────────────────────────┐
│  Go Package Structure               │
├─────────────────────────────────────┤
│  main package     ████████████ 100% │
│  custom packages  ███████████▌  95% │
│  standard library ████████████ 100% │
│  third-party      ██████▌      65%  │
└─────────────────────────────────────┘
```

**Every Go file belongs to a package!** Here's what you need to know:

#### **The `main` Package - Your Program's Entry Point**

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

> 🎯 **Key Point**: The `main` package is special - it tells Go "this is an executable program, not a library."

#### **Custom Packages - Your Reusable Code**

```go
// File: math/calculator.go
package math

// Add performs addition of two integers
func Add(a, b int) int {
    return a + b
}

// multiply is not exported (lowercase first letter)
func multiply(a, b int) int {
    return a * b
}
```

---

## 🔍 Package Visibility Rules

<div style="background: linear-gradient(135deg, #ff9a9e 0%, #fecfef 100%); padding: 15px; border-radius: 8px; margin: 20px 0;">

### 🚪 The "Door" Analogy

Think of **uppercase** names as **open doors** 🚪 - anyone can walk through them.
Think of **lowercase** names as **locked doors** 🔒 - only people inside the house (package) can use them.

</div>

### **Exported vs Unexported Identifiers**

```go
package utils

import "strings"

// Public - can be used by other packages
func FormatName(name string) string {
    return strings.Title(strings.ToLower(name))
}

// Public constant
const MaxRetries = 3

// Public type
type User struct {
    Name  string // Public field
    email string // private field
}

// private - only usable within this package
func validateEmail(email string) bool {
    return strings.Contains(email, "@")
}
```

---

## 🏗️ Creating Your First Package

Let's build a **calculator package** step by step!

### **Step 1: Project Structure**

```
my-calculator/
├── main.go
└── calculator/
    ├── basic.go
    └── advanced.go
```

### **Step 2: Create the Calculator Package**

Create `calculator/basic.go`:

```go
package calculator

import "fmt"

// Add performs addition of two numbers
func Add(a, b float64) float64 {
    return a + b
}

// Subtract performs subtraction
func Subtract(a, b float64) float64 {
    return a - b
}

// Multiply performs multiplication
func Multiply(a, b float64) float64 {
    return a * b
}

// Divide performs division with error handling
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

// Version returns the calculator version
const Version = "1.0.0"
```

Create `calculator/advanced.go`:

```go
package calculator

import "math"

// Power calculates a^b
func Power(base, exponent float64) float64 {
    return math.Pow(base, exponent)
}

// SquareRoot calculates square root
func SquareRoot(n float64) (float64, error) {
    if n < 0 {
        return 0, fmt.Errorf("cannot calculate square root of negative number")
    }
    return math.Sqrt(n), nil
}

// factorial is a private helper function
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// Factorial calculates factorial (public interface)
func Factorial(n int) (int, error) {
    if n < 0 {
        return 0, fmt.Errorf("factorial not defined for negative numbers")
    }
    return factorial(n), nil
}
```

### **Step 3: Use Your Package**

Create `main.go`:

```go
package main

import (
    "fmt"
    "log"
    
    "./calculator" // Local import
)

func main() {
    fmt.Printf("Calculator v%s\n", calculator.Version)
    fmt.Println("===================")
    
    // Basic operations
    result := calculator.Add(10, 5)
    fmt.Printf("10 + 5 = %.2f\n", result)
    
    result = calculator.Multiply(4, 7)
    fmt.Printf("4 × 7 = %.2f\n", result)
    
    // Division with error handling
    if quotient, err := calculator.Divide(20, 4); err != nil {
        log.Printf("Error: %v", err)
    } else {
        fmt.Printf("20 ÷ 4 = %.2f\n", quotient)
    }
    
    // Advanced operations
    power := calculator.Power(2, 8)
    fmt.Printf("2^8 = %.0f\n", power)
    
    if sqrt, err := calculator.SquareRoot(16); err != nil {
        log.Printf("Error: %v", err)
    } else {
        fmt.Printf("√16 = %.2f\n", sqrt)
    }
    
    if fact, err := calculator.Factorial(5); err != nil {
        log.Printf("Error: %v", err)
    } else {
        fmt.Printf("5! = %d\n", fact)
    }
}
```

---

## 📥 Import Statements Deep Dive

<div style="background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%); padding: 15px; border-radius: 8px; margin: 20px 0;">

### 🎪 Import Circus

Imports are like **inviting performers** to your circus! Each import brings specific talents (functions, types) to your program.

</div>

### **Different Import Styles**

```go
package main

import (
    // Standard library
    "fmt"
    "strings"
    "net/http"
    
    // Aliased import
    f "fmt"
    
    // Dot import (use sparingly!)
    . "math"
    
    // Blank import (for side effects)
    _ "image/png"
    
    // Local package
    "./mypackage"
    
    // Remote package (with go modules)
    "github.com/gorilla/mux"
)

func main() {
    // Standard usage
    fmt.Println("Hello")
    
    // Aliased usage
    f.Println("Hello from alias")
    
    // Dot import usage (no package prefix needed)
    result := Sqrt(16) // instead of math.Sqrt(16)
    
    fmt.Printf("Square root: %.2f\n", result)
}
```

---

## 🏆 Best Practices for Packages

### **1. Package Naming Conventions**

```go
// ✅ Good package names
package user
package http
package json
package calculator

// ❌ Avoid these
package userPackage  // redundant "Package"
package utils        // too generic
package helper       // too vague
package myAwesomePackage // too long
```

### **2. File Organization**

```
// ✅ Good structure
myproject/
├── main.go
├── user/
│   ├── user.go      // main types and functions
│   ├── auth.go      // authentication related
│   └── profile.go   // profile management
└── database/
    ├── connection.go
    └── queries.go

// ❌ Poor structure
myproject/
├── main.go
├── everything.go    // everything in one file
└── stuff/
    └── random.go    // unclear purpose
```

### **3. Documentation Comments**

```go
// Package calculator provides basic and advanced mathematical operations.
//
// This package includes functions for arithmetic operations, power calculations,
// and factorial computations with proper error handling.
package calculator

// Add performs addition of two floating-point numbers.
//
// Example:
//   result := calculator.Add(3.5, 2.1)
//   fmt.Printf("Result: %.2f", result) // Output: Result: 5.60
func Add(a, b float64) float64 {
    return a + b
}
```

---

## 🛠️ Practical Exercise: Building a Text Utilities Package

Let's create a practical package for text processing!

### **Step 1: Create the Package Structure**

```
textutils/
├── main.go
└── textprocessor/
    ├── formatter.go
    ├── analyzer.go
    └── validator.go
```

### **Step 2: Implement the Text Processor**

Create `textprocessor/formatter.go`:

```go
package textprocessor

import (
    "strings"
    "unicode"
)

// TitleCase converts text to title case
func TitleCase(text string) string {
    return strings.Title(strings.ToLower(text))
}

// CamelCase converts text to camelCase
func CamelCase(text string) string {
    words := strings.Fields(text)
    if len(words) == 0 {
        return ""
    }
    
    result := strings.ToLower(words[0])
    for i := 1; i < len(words); i++ {
        result += strings.Title(strings.ToLower(words[i]))
    }
    return result
}

// SnakeCase converts text to snake_case
func SnakeCase(text string) string {
    var result []rune
    for i, r := range text {
        if unicode.IsUpper(r) && i > 0 {
            result = append(result, '_')
        }
        result = append(result, unicode.ToLower(r))
    }
    return string(result)
}

// Reverse reverses a string
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

Create `textprocessor/analyzer.go`:

```go
package textprocessor

import (
    "strings"
    "unicode"
)

// TextStats holds statistics about text
type TextStats struct {
    Characters int
    Words      int
    Lines      int
    Paragraphs int
}

// AnalyzeText returns statistics about the given text
func AnalyzeText(text string) TextStats {
    stats := TextStats{
        Characters: len(text),
        Words:      len(strings.Fields(text)),
        Lines:      strings.Count(text, "\n") + 1,
        Paragraphs: len(strings.Split(strings.TrimSpace(text), "\n\n")),
    }
    
    if strings.TrimSpace(text) == "" {
        stats.Lines = 0
        stats.Paragraphs = 0
    }
    
    return stats
}

// CountVowels counts vowels in text
func CountVowels(text string) int {
    vowels := "aeiouAEIOU"
    count := 0
    for _, char := range text {
        if strings.ContainsRune(vowels, char) {
            count++
        }
    }
    return count
}

// IsAlphanumeric checks if text contains only letters and numbers
func IsAlphanumeric(text string) bool {
    for _, char := range text {
        if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
            return false
        }
    }
    return len(text) > 0
}
```

Create `textprocessor/validator.go`:

```go
package textprocessor

import (
    "regexp"
    "strings"
)

// IsValidEmail checks if the string is a valid email format
func IsValidEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}

// IsValidPhone checks if the string is a valid phone number
func IsValidPhone(phone string) bool {
    // Remove common separators
    cleaned := strings.ReplaceAll(phone, "-", "")
    cleaned = strings.ReplaceAll(cleaned, " ", "")
    cleaned = strings.ReplaceAll(cleaned, "(", "")
    cleaned = strings.ReplaceAll(cleaned, ")", "")
    
    phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
    return phoneRegex.MatchString(cleaned)
}

// ContainsProfanity checks for common profanity (simplified example)
func ContainsProfanity(text string) bool {
    // This is a simplified example - in production, use a proper filter
    profanityList := []string{"badword1", "badword2"} // Add actual words
    
    lowerText := strings.ToLower(text)
    for _, word := range profanityList {
        if strings.Contains(lowerText, word) {
            return true
        }
    }
    return false
}

// IsStrongPassword checks password strength
func IsStrongPassword(password string) bool {
    if len(password) < 8 {
        return false
    }
    
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
    hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
    hasDigit := regexp.MustCompile(`\d`).MatchString(password)
    hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
    
    return hasUpper && hasLower && hasDigit && hasSpecial
}
```

### **Step 3: Create the Main Application**

Create `main.go`:

```go
package main

import (
    "fmt"
    "strings"
    
    "./textprocessor"
)

func main() {
    fmt.Println("🔤 Text Utilities Demo")
    fmt.Println(strings.Repeat("=", 40))
    
    // Sample text
    text := "Hello World! This is a SAMPLE text for processing."
    
    // Formatting demonstrations
    fmt.Println("\n📝 Text Formatting:")
    fmt.Printf("Original:    %s\n", text)
    fmt.Printf("Title Case:  %s\n", textprocessor.TitleCase(text))
    fmt.Printf("Camel Case:  %s\n", textprocessor.CamelCase(text))
    fmt.Printf("Snake Case:  %s\n", textprocessor.SnakeCase(text))
    fmt.Printf("Reversed:    %s\n", textprocessor.Reverse(text))
    
    // Text analysis
    fmt.Println("\n📊 Text Analysis:")
    stats := textprocessor.AnalyzeText(text)
    fmt.Printf("Characters:  %d\n", stats.Characters)
    fmt.Printf("Words:       %d\n", stats.Words)
    fmt.Printf("Lines:       %d\n", stats.Lines)
    fmt.Printf("Vowels:      %d\n", textprocessor.CountVowels(text))
    fmt.Printf("Alphanumeric: %t\n", textprocessor.IsAlphanumeric("Hello123"))
    
    // Validation examples
    fmt.Println("\n✅ Validation Examples:")
    
    emails := []string{"user@example.com", "invalid-email", "test@domain.co.uk"}
    for _, email := range emails {
        fmt.Printf("Email '%s' is valid: %t\n", email, textprocessor.IsValidEmail(email))
    }
    
    phones := []string{"+1234567890", "123-456-7890", "invalid-phone"}
    for _, phone := range phones {
        fmt.Printf("Phone '%s' is valid: %t\n", phone, textprocessor.IsValidPhone(phone))
    }
    
    passwords := []string{"Password123!", "weak", "Strong@Pass1"}
    for _, password := range passwords {
        fmt.Printf("Password '%s' is strong: %t\n", password, textprocessor.IsStrongPassword(password))
    }
}
```

---

## 🎯 Practice Challenges

### **Challenge 1: Math Package Extension**
Extend the calculator package with:
- Trigonometric functions (Sin, Cos, Tan)
- Logarithmic functions
- Statistical functions (Mean, Median, Mode)

### **Challenge 2: File Utilities Package**
Create a `fileutils` package with functions to:
- Check if a file exists
- Get file size and modification time
- Read file contents safely
- Create backup copies

### **Challenge 3: Web Utilities Package**
Build a `webutils` package that includes:
- URL validation
- HTML tag stripping
- Query parameter parsing
- Basic HTTP status code meanings

---

## 🚀 Running Your Package

```bash
# Navigate to your project directory
cd my-calculator

# Run your program
go run main.go

# Build an executable
go build -o calculator main.go

# Run the executable
./calculator
```

**Expected Output:**
```
Calculator v1.0.0
===================
10 + 5 = 15.00
4 × 7 = 28.00
20 ÷ 4 = 5.00
2^8 = 256
√16 = 4.00
5! = 120
```

---

## 📚 Key Takeaways

<div style="background: linear-gradient(135deg, #ffecd2 0%, #fcb69f 100%); padding: 20px; border-radius: 10px; margin: 20px 0;">

### 🎯 Remember These Golden Rules:

1. **📦 Package = Toolbox**: Group related functionality together
2. **🔑 Uppercase = Public**: Capital letters make things accessible to other packages
3. **🔒 Lowercase = Private**: Lowercase keeps things internal to the package
4. **📝 Document Everything**: Good comments make your packages user-friendly
5. **🏗️ Structure Matters**: Organize files logically within packages
6. **🎯 Single Responsibility**: Each package should have one clear purpose

</div>

---

## 🔗 What's Next?

In the next module, we'll explore:
- **Go Modules**: Managing dependencies like a pro
- **Third-party packages**: Using external libraries
- **Publishing packages**: Sharing your code with the world
- **Testing packages**: Ensuring your code works correctly

---

<div align="center">

**🎉 Congratulations!** 

You've mastered Go packages and can now build reusable, modular code!

*Ready to level up? Let's dive into the next module!* 🚀

</div>
