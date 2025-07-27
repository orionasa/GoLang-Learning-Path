# 🎯 Module 2: Variables & Data Types

<div align="center">

![Go Variables](https://via.placeholder.com/400x200/4A90E2/FFFFFF?text=Go+Variables+%26+Data+Types)

**Master the Building Blocks of Go!**

*From Zero to Go Hero - Module 2*

</div>

---

## 📋 Module Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 3-4 hours    | Beginner       | Module 1 Complete |

### 🎯 Learning Objectives
By the end of this module, you will:
- ✅ Understand Go's type system and variable declarations
- ✅ Master all basic data types in Go
- ✅ Know how to work with constants and variables
- ✅ Understand type conversion and type inference
- ✅ Write programs using different data types effectively

---

## 🧱 Variables: The Foundation of Programming

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 What Are Variables?

Variables are like **labeled boxes** that store data in your program. Think of them as containers that hold different types of information!

</div>

### 📦 Variable Declaration Patterns

```
┌─────────────────────────────────────┐
│  Go Variable Declaration Methods    │
├─────────────────────────────────────┤
│  var name type = value             │ ← Explicit
│  var name = value                  │ ← Type inference  
│  name := value                     │ ← Short declaration
│  var name type                     │ ← Zero value
└─────────────────────────────────────┘
```

### 🎨 **Method 1: Explicit Declaration**

```go
package main

import "fmt"

func main() {
    // Explicit type declaration
    var message string = "Hello, Go!"
    var age int = 25
    var height float64 = 5.9
    var isStudent bool = true
    
    fmt.Println("Message:", message)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
}
```

<div style="background: #e8f5e8; border-left: 4px solid #4caf50; padding: 15px; margin: 15px 0;">
<strong>💡 Pro Tip:</strong> Explicit declarations are great when you want to be crystal clear about the data type!
</div>

### 🚀 **Method 2: Type Inference**

```go
package main

import "fmt"

func main() {
    // Go automatically figures out the type!
    var name = "Alice"        // string
    var score = 95           // int
    var average = 87.5       // float64
    var passed = true        // bool
    
    fmt.Printf("Name: %s (Type: %T)\n", name, name)
    fmt.Printf("Score: %d (Type: %T)\n", score, score)
    fmt.Printf("Average: %.1f (Type: %T)\n", average, average)
    fmt.Printf("Passed: %t (Type: %T)\n", passed, passed)
}
```

### ⚡ **Method 3: Short Declaration (Most Popular!)**

```go
package main

import "fmt"

func main() {
    // The := operator is Go's magic wand!
    username := "gopher"
    points := 1000
    multiplier := 2.5
    active := true
    
    // You can declare multiple variables at once
    x, y, z := 10, 20, 30
    
    fmt.Println("Username:", username)
    fmt.Println("Points:", points)
    fmt.Println("Multiplier:", multiplier)
    fmt.Println("Active:", active)
    fmt.Println("Coordinates:", x, y, z)
}
```

<div style="background: #fff3cd; border-left: 4px solid #ffc107; padding: 15px; margin: 15px 0;">
<strong>⚠️ Important:</strong> The <code>:=</code> operator can only be used inside functions, not at the package level!
</div>

---

## 🎭 Go's Data Types: The Cast of Characters

<div style="background: linear-gradient(135deg, #ff9a9e 0%, #fecfef 50%, #fecfef 100%); padding: 20px; border-radius: 10px; color: #333; margin: 20px 0;">

### 🌈 Data Types Family Tree

Go has a rich family of data types, each with its own special purpose!

</div>

```
                    Go Data Types
                         │
        ┌────────────────┼────────────────┐
        │                │                │
    Basic Types      Composite Types   Interface Types
        │                │                │
    ┌───┼───┐        ┌───┼───┐           │
    │   │   │        │   │   │           │
  Numbers │ Others  Arrays │ Others    interface{}
    │     │   │        │   │
  ┌─┼─┐   │   │        │   │
  │ │ │   │   │        │   │
 int│float│  bool    slice map
    │     │  string   struct
  uint    │  rune     channel
         byte        pointer
```

### 🔢 **Numeric Types: The Number Crunchers**

#### **Integers (Whole Numbers)**

```go
package main

import "fmt"

func main() {
    // Signed integers (can be positive or negative)
    var age int8 = 25          // -128 to 127
    var population int16 = 5000 // -32,768 to 32,767
    var distance int32 = 1000000 // -2 billion to 2 billion
    var bigNumber int64 = 9223372036854775807 // Very big range!
    
    // Unsigned integers (only positive)
    var level uint8 = 255      // 0 to 255
    var score uint16 = 65535   // 0 to 65,535
    var id uint32 = 4294967295 // 0 to 4+ billion
    
    // Platform-dependent sizes
    var count int = 42         // 32 or 64 bits depending on system
    var index uint = 100       // 32 or 64 bits depending on system
    
    fmt.Println("=== Integer Showcase ===")
    fmt.Printf("Age (int8): %d\n", age)
    fmt.Printf("Population (int16): %d\n", population)
    fmt.Printf("Distance (int32): %d\n", distance)
    fmt.Printf("Big Number (int64): %d\n", bigNumber)
    fmt.Printf("Level (uint8): %d\n", level)
    fmt.Printf("Score (uint16): %d\n", score)
    fmt.Printf("ID (uint32): %d\n", id)
    fmt.Printf("Count (int): %d\n", count)
    fmt.Printf("Index (uint): %d\n", index)
}
```

#### **Floating-Point Numbers (Decimals)**

```go
package main

import "fmt"

func main() {
    // Single precision (32-bit)
    var temperature float32 = 98.6
    var ratio float32 = 3.14159
    
    // Double precision (64-bit) - More accurate!
    var pi float64 = 3.141592653589793
    var e float64 = 2.718281828459045
    
    // Scientific notation
    var lightSpeed float64 = 2.998e8  // 299,800,000
    var atomSize float64 = 1.23e-10   // 0.000000000123
    
    fmt.Println("=== Float Showcase ===")
    fmt.Printf("Temperature (float32): %.1f°F\n", temperature)
    fmt.Printf("Ratio (float32): %.5f\n", ratio)
    fmt.Printf("Pi (float64): %.15f\n", pi)
    fmt.Printf("Euler's number (float64): %.15f\n", e)
    fmt.Printf("Light Speed: %.0f m/s\n", lightSpeed)
    fmt.Printf("Atom Size: %.2e meters\n", atomSize)
}
```

<div style="background: #e3f2fd; border-left: 4px solid #2196f3; padding: 15px; margin: 15px 0;">
<strong>🎯 Best Practice:</strong> Use <code>float64</code> for most decimal calculations - it's more precise and is Go's default floating-point type!
</div>

### 📝 **Text Types: The Communicators**

#### **Strings: Text Data**

```go
package main

import "fmt"

func main() {
    // Basic string declarations
    greeting := "Hello, World!"
    name := "Go Programmer"
    
    // Multi-line strings using backticks
    poem := `Roses are red,
Violets are blue,
Go is awesome,
And so are you!`
    
    // String with escape characters
    message := "She said, \"Go is amazing!\"\nI totally agree."
    
    // Raw strings (backticks) ignore escape characters
    path := `C:\Users\GoLang\Projects\awesome-app`
    
    fmt.Println("=== String Showcase ===")
    fmt.Println("Greeting:", greeting)
    fmt.Println("Name:", name)
    fmt.Println("\nPoem:")
    fmt.Println(poem)
    fmt.Println("\nMessage:")
    fmt.Println(message)
    fmt.Println("\nPath:", path)
    
    // String operations
    fmt.Println("\n=== String Operations ===")
    fmt.Printf("Length of greeting: %d characters\n", len(greeting))
    fmt.Printf("First character: %c\n", greeting[0])
    fmt.Printf("Concatenation: %s\n", greeting+" "+name)
}
```

#### **Runes: Unicode Characters**

```go
package main

import "fmt"

func main() {
    // Runes represent Unicode code points
    var heart rune = '❤'
    var smiley rune = '😊'
    var letter rune = 'A'
    var chinese rune = '中'
    
    // You can also use Unicode code points
    var star rune = '\u2605'  // ★
    var music rune = '\U0001F3B5'  // 🎵
    
    fmt.Println("=== Rune Showcase ===")
    fmt.Printf("Heart: %c (Unicode: %U)\n", heart, heart)
    fmt.Printf("Smiley: %c (Unicode: %U)\n", smiley, smiley)
    fmt.Printf("Letter: %c (Unicode: %U)\n", letter, letter)
    fmt.Printf("Chinese: %c (Unicode: %U)\n", chinese, chinese)
    fmt.Printf("Star: %c (Unicode: %U)\n", star, star)
    fmt.Printf("Music: %c (Unicode: %U)\n", music, music)
    
    // Working with strings and runes
    text := "Hello, 世界! 🌍"
    fmt.Printf("\nText: %s\n", text)
    fmt.Printf("Byte length: %d\n", len(text))
    fmt.Printf("Rune count: %d\n", len([]rune(text)))
}
```

### ✅ **Boolean: True or False**

```go
package main

import "fmt"

func main() {
    // Boolean values
    var isOnline bool = true
    var isComplete bool = false
    
    // Boolean from expressions
    age := 20
    canVote := age >= 18
    
    score := 85
    passed := score >= 60
    excellent := score >= 90
    
    fmt.Println("=== Boolean Showcase ===")
    fmt.Printf("Is Online: %t\n", isOnline)
    fmt.Printf("Is Complete: %t\n", isComplete)
    fmt.Printf("Can Vote (age %d): %t\n", age, canVote)
    fmt.Printf("Passed (score %d): %t\n", score, passed)
    fmt.Printf("Excellent (score %d): %t\n", score, excellent)
    
    // Boolean operations
    fmt.Println("\n=== Boolean Operations ===")
    fmt.Printf("true && false = %t\n", true && false)
    fmt.Printf("true || false = %t\n", true || false)
    fmt.Printf("!true = %t\n", !true)
    fmt.Printf("!false = %t\n", !false)
}
```

---

## 🔒 Constants: The Unchangeable Values

<div style="background: linear-gradient(135deg, #ffecd2 0%, #fcb69f 100%); padding: 20px; border-radius: 10px; color: #333; margin: 20px 0;">

### 🛡️ Constants: Set in Stone

Constants are like **permanent markers** - once you write them, they can't be erased or changed!

</div>

### 📌 **Basic Constants**

```go
package main

import "fmt"

// Package-level constants
const CompanyName = "Go Corp"
const MaxUsers = 1000
const Pi = 3.14159265359

func main() {
    // Function-level constants
    const greeting = "Welcome!"
    const version = "1.0.0"
    const debugMode = true
    
    fmt.Println("=== Constants Showcase ===")
    fmt.Printf("Company: %s\n", CompanyName)
    fmt.Printf("Max Users: %d\n", MaxUsers)
    fmt.Printf("Pi: %.10f\n", Pi)
    fmt.Printf("Greeting: %s\n", greeting)
    fmt.Printf("Version: %s\n", version)
    fmt.Printf("Debug Mode: %t\n", debugMode)
}
```

---

## 🔄 Type Conversion: Changing Forms

<div style="background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%); padding: 20px; border-radius: 10px; color: #333; margin: 20px 0;">

### 🎭 Type Conversion Magic

Sometimes you need to transform data from one type to another. Go makes this explicit and safe!

</div>

### 🔢 **Numeric Conversions**

```go
package main

import "fmt"

func main() {
    // Original values
    var i int = 42
    var f float64 = 3.14159
    var b byte = 255
    
    fmt.Println("=== Original Values ===")
    fmt.Printf("Integer: %d (Type: %T)\n", i, i)
    fmt.Printf("Float: %.5f (Type: %T)\n", f, f)
    fmt.Printf("Byte: %d (Type: %T)\n", b, b)
    
    // Converting between numeric types
    fmt.Println("\n=== Conversions ===")
    
    // int to float64
    floatFromInt := float64(i)
    fmt.Printf("int to float64: %.2f (Type: %T)\n", floatFromInt, floatFromInt)
    
    // float64 to int (truncates decimal part)
    intFromFloat := int(f)
    fmt.Printf("float64 to int: %d (Type: %T)\n", intFromFloat, intFromFloat)
    
    // byte to int
    intFromByte := int(b)
    fmt.Printf("byte to int: %d (Type: %T)\n", intFromByte, intFromByte)
    
    // Demonstrating precision loss
    bigFloat := 123.789
    truncated := int(bigFloat)
    fmt.Printf("\nPrecision loss: %.3f -> %d\n", bigFloat, truncated)
}
```

---

## 🎪 Practical Examples: Putting It All Together

### 🏪 **Example 1: Online Store Calculator**

```go
package main

import "fmt"

func main() {
    // Store information
    const storeName = "Go Gadgets"
    const taxRate = 0.08 // 8% tax
    
    // Product information
    productName := "Go Programming Book"
    price := 29.99
    quantity := 2
    isPremiumMember := true
    
    // Calculations
    subtotal := price * float64(quantity)
    tax := subtotal * taxRate
    
    // Premium member discount (10%)
    discount := 0.0
    if isPremiumMember {
        discount = subtotal * 0.10
    }
    
    total := subtotal + tax - discount
    
    // Display receipt
    fmt.Printf("=== %s RECEIPT ===\n", storeName)
    fmt.Printf("Product: %s\n", productName)
    fmt.Printf("Price: $%.2f x %d\n", price, quantity)
    fmt.Printf("Premium Member: %t\n", isPremiumMember)
    fmt.Println("------------------------")
    fmt.Printf("Subtotal: $%.2f\n", subtotal)
    fmt.Printf("Tax (%.0f%%): $%.2f\n", taxRate*100, tax)
    
    if discount > 0 {
        fmt.Printf("Premium Discount: -$%.2f\n", discount)
    }
    
    fmt.Println("------------------------")
    fmt.Printf("TOTAL: $%.2f\n", total)
}
```

---

## 🎯 Module 2 Summary

<div style="background: #d1ecf1; border: 1px solid #bee5eb; border-radius: 8px; padding: 15px;">

### ✅ **What You've Learned**

- **Variables**: Four ways to declare variables in Go
- **Data Types**: Numbers, strings, booleans, and more
- **Constants**: Unchangeable values with `const`
- **Type Conversion**: Safe transformation between types
- **Practical Applications**: Real-world examples

### 🚀 **Next Steps**

Ready for Module 3? You'll learn about:
- Control flow (if/else, loops)
- Functions and methods
- Error handling
- More advanced programming concepts

</div>

---

## 🏆 Practice Exercises

### 🎯 **Exercise 1: Personal Information**
Create a program that stores and displays your personal information using different data types.

### 🎯 **Exercise 2: Temperature Converter**
Build a program that converts temperatures between Fahrenheit and Celsius.

### 🎯 **Exercise 3: Shopping Cart**
Create a shopping cart calculator with multiple items, tax, and discounts.

---

<div align="center">

**🎉 Congratulations! You've mastered Go variables and data types! 🎉**

*Ready to continue your Go journey? Let's move on to Module 3!*

</div>
