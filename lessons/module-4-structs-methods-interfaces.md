# 🚀 Module 4: Structs, Methods & Interfaces

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Building Complex Data Structures and Behaviors**

*From Zero to Go Hero - Module 4*

</div>

---

## 📋 Module Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 3-4 hours    | Beginner       | Modules 1-3 (Basic Syntax, Functions, Control Flow) |

### 🎯 Learning Objectives
By the end of this module, you will:
- ✅ Create and use custom data structures with structs
- ✅ Understand the difference between value and pointer receivers
- ✅ Define and implement methods on custom types
- ✅ Master Go's interface system for flexible code design
- ✅ Apply composition over inheritance principles in Go

---

## 🏗️ Understanding Structs: Building Custom Data Types

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 What Are Structs?

Structs are Go's way of creating custom data types that group related data together. Think of them as blueprints for creating objects that hold multiple pieces of information.

</div>

### 🏠 **1. Basic Struct Declaration**

```
┌─────────────────────────────────────┐
│  Struct Components                  │
├─────────────────────────────────────┤
│  Field Names    ████████████████ 100%│
│  Field Types    ███████████████▌ 97% │
│  Tags (optional)████████▌       42% │
│  Methods        ████▌           22% │
└─────────────────────────────────────┘
```

Let's start with a simple example:

```go
package main

import "fmt"

// Define a Person struct
type Person struct {
    Name    string
    Age     int
    Email   string
    IsAdmin bool
}

func main() {
    // Creating a struct instance
    var person1 Person
    person1.Name = "Alice"
    person1.Age = 30
    person1.Email = "alice@example.com"
    person1.IsAdmin = false
    
    fmt.Printf("Person: %+v\n", person1)
}
```

### 🎯 **2. Different Ways to Initialize Structs**

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Struct Initialization Methods** - Multiple ways to create struct instances

```go
// Method 1: Zero value initialization
var person1 Person

// Method 2: Struct literal with field names
person2 := Person{
    Name:    "Bob",
    Age:     25,
    Email:   "bob@example.com",
    IsAdmin: true,
}

// Method 3: Struct literal without field names (order matters!)
person3 := Person{"Charlie", 35, "charlie@example.com", false}

// Method 4: Partial initialization
person4 := Person{
    Name: "Diana",
    Age:  28,
    // Email and IsAdmin will be zero values
}
```

</div>

### 🔍 **3. Nested Structs and Composition**

Go encourages composition over inheritance. Here's how to build complex structures:

```go
type Address struct {
    Street  string
    City    string
    Country string
    ZipCode string
}

type Employee struct {
    Person  Person  // Embedded struct
    Address Address // Nested struct
    Salary  float64
    Department string
}

func main() {
    emp := Employee{
        Person: Person{
            Name:    "John Doe",
            Age:     32,
            Email:   "john@company.com",
            IsAdmin: false,
        },
        Address: Address{
            Street:  "123 Main St",
            City:    "New York",
            Country: "USA",
            ZipCode: "10001",
        },
        Salary:     75000.0,
        Department: "Engineering",
    }
    
    // Accessing nested fields
    fmt.Printf("Employee: %s lives in %s\n", emp.Person.Name, emp.Address.City)
}
```

### 🎨 **4. Anonymous Fields and Embedding**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🔗 Anonymous Fields</strong><br>
Fields without explicit names, accessed directly
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📦 Embedding</strong><br>
Including one struct inside another for composition
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🎯 Promotion</strong><br>
Embedded fields can be accessed directly
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🔄 Method Sets</strong><br>
Embedded types bring their methods along
</div>

</div>

```go
type User struct {
    Person    // Anonymous field - embedding
    Username  string
    LastLogin time.Time
}

func main() {
    user := User{
        Person: Person{
            Name:  "Alice Smith",
            Age:   28,
            Email: "alice@example.com",
        },
        Username:  "alice_smith",
        LastLogin: time.Now(),
    }
    
    // Direct access to embedded fields (promotion)
    fmt.Printf("User %s is %d years old\n", user.Name, user.Age)
    // Instead of: user.Person.Name, user.Person.Age
}
```

---

## 🔧 Methods: Adding Behavior to Structs

<div style="background: linear-gradient(135deg, #ff9a9e 0%, #fecfef 50%, #fecfef 100%); padding: 20px; border-radius: 10px; color: #333; margin: 20px 0;">

### 🌟 What Are Methods?

Methods are functions that belong to a specific type. They allow you to define behavior for your custom types, making your code more organized and object-oriented.

</div>

### 📝 **Step 1: Basic Method Declaration**

Methods have a special receiver that comes before the method name:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with value receiver
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    fmt.Printf("Area: %.2f\n", rect.Area())
    fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}
```

### 🎯 **Step 2: Value vs Pointer Receivers**

<div style="background: #fff3cd; border: 1px solid #ffeaa7; border-radius: 8px; padding: 15px; margin: 10px 0;">

**⚠️ Important Concept**: The choice between value and pointer receivers affects performance and mutability.

</div>

```go
type BankAccount struct {
    Owner   string
    Balance float64
}

// Value receiver - receives a COPY of the struct
func (b BankAccount) GetBalance() float64 {
    return b.Balance
}

// Pointer receiver - receives a POINTER to the struct
func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount // This modifies the original struct
}

// Value receiver trying to modify (WON'T WORK as expected)
func (b BankAccount) WithdrawWrong(amount float64) {
    b.Balance -= amount // This modifies only the COPY!
}

// Pointer receiver for modification (CORRECT WAY)
func (b *BankAccount) Withdraw(amount float64) bool {
    if b.Balance >= amount {
        b.Balance -= amount
        return true
    }
    return false
}

func main() {
    account := BankAccount{
        Owner:   "John Doe",
        Balance: 1000.0,
    }
    
    fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())
    
    account.Deposit(500.0)
    fmt.Printf("After deposit: $%.2f\n", account.GetBalance())
    
    if account.Withdraw(200.0) {
        fmt.Printf("After withdrawal: $%.2f\n", account.GetBalance())
    }
}
```

### 🔄 **Step 3: Method Sets and Type Conversion**

```go
type Temperature float64

// Methods on custom types
func (t Temperature) Celsius() float64 {
    return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
    return float64(t)*9/5 + 32
}

func (t Temperature) Kelvin() float64 {
    return float64(t) + 273.15
}

// Method with pointer receiver for modification
func (t *Temperature) SetCelsius(celsius float64) {
    *t = Temperature(celsius)
}

func main() {
    var temp Temperature = 25.0 // 25°C
    
    fmt.Printf("%.1f°C = %.1f°F = %.1fK\n", 
        temp.Celsius(), 
        temp.Fahrenheit(), 
        temp.Kelvin())
    
    temp.SetCelsius(30.0)
    fmt.Printf("New temperature: %.1f°C\n", temp.Celsius())
}
```

---

## 🎭 Interfaces: Defining Behavior Contracts

<div style="background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%); padding: 20px; border-radius: 10px; color: #333; margin: 20px 0;">

### 🌟 Understanding Interfaces

Interfaces in Go define a contract of behavior. They specify what methods a type must have, but not how those methods are implemented. This enables powerful polymorphism and flexible design.

</div>

### 🎯 **Step 1: Basic Interface Declaration**

```go
// Define an interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Rectangle implements Shape interface
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Circle implements Shape interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

// Function that works with any Shape
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 7}
    
    // Both Rectangle and Circle implement Shape interface
    PrintShapeInfo(rect)
    PrintShapeInfo(circle)
}
```

### 🔍 **Step 2: Empty Interface and Type Assertions**

<div style="background: #f8f9fa; border-left: 4px solid #007bff; padding: 15px; margin: 10px 0;">

**The Empty Interface** - `interface{}` can hold any type, similar to `Object` in other languages

```go
func PrintAnything(value interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func main() {
    PrintAnything(42)
    PrintAnything("Hello")
    PrintAnything([]int{1, 2, 3})
    PrintAnything(Rectangle{Width: 5, Height: 3})
}
```

</div>

**Type Assertions** help you get the concrete type back:

```go
func ProcessValue(value interface{}) {
    // Type assertion with safety check
    if str, ok := value.(string); ok {
        fmt.Printf("String value: %s (length: %d)\n", str, len(str))
    } else if num, ok := value.(int); ok {
        fmt.Printf("Integer value: %d (squared: %d)\n", num, num*num)
    } else {
        fmt.Printf("Unknown type: %T\n", value)
    }
}

func main() {
    ProcessValue("Hello, Go!")
    ProcessValue(42)
    ProcessValue(3.14)
}
```

### 🎨 **Step 3: Interface Composition**

```go
// Basic interfaces
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string)
}

// Composed interface
type ReadWriter interface {
    Reader
    Writer
}

// File struct implementing ReadWriter
type File struct {
    Name    string
    Content string
}

func (f *File) Read() string {
    return f.Content
}

func (f *File) Write(data string) {
    f.Content += data
}

// Function accepting the composed interface
func ProcessFile(rw ReadWriter) {
    content := rw.Read()
    fmt.Printf("Current content: %s\n", content)
    rw.Write(" - Updated!")
}

func main() {
    file := &File{Name: "example.txt", Content: "Hello"}
    ProcessFile(file)
    fmt.Printf("Final content: %s\n", file.Read())
}
```

---

## 🏗️ Practical Example: Building a Simple Library System

Let's combine everything we've learned to build a practical example:

### 📚 **Step 1: Define Core Structures**

```go
package main

import (
    "fmt"
    "time"
)

// Base interfaces
type Borrowable interface {
    GetID() string
    GetTitle() string
    IsAvailable() bool
    Borrow() error
    Return() error
}

type Searchable interface {
    Search(query string) bool
}

// LibraryItem combines both interfaces
type LibraryItem interface {
    Borrowable
    Searchable
}

// Base struct for common fields
type Item struct {
    ID          string
    Title       string
    Available   bool
    BorrowedBy  string
    BorrowDate  time.Time
}

// Book struct embedding Item
type Book struct {
    Item
    Author string
    ISBN   string
    Pages  int
}

// DVD struct embedding Item
type DVD struct {
    Item
    Director string
    Duration int // minutes
    Genre    string
}
```

### 🎯 **Step 2: Implement Methods**

```go
// Methods for Item (will be inherited by Book and DVD)
func (i *Item) GetID() string {
    return i.ID
}

func (i *Item) GetTitle() string {
    return i.Title
}

func (i *Item) IsAvailable() bool {
    return i.Available
}

func (i *Item) Borrow() error {
    if !i.Available {
        return fmt.Errorf("item %s is already borrowed", i.Title)
    }
    i.Available = false
    i.BorrowDate = time.Now()
    return nil
}

func (i *Item) Return() error {
    if i.Available {
        return fmt.Errorf("item %s is not currently borrowed", i.Title)
    }
    i.Available = true
    i.BorrowedBy = ""
    return nil
}

// Book-specific methods
func (b Book) Search(query string) bool {
    return contains(b.Title, query) || contains(b.Author, query) || contains(b.ISBN, query)
}

func (b Book) GetInfo() string {
    return fmt.Sprintf("Book: %s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

// DVD-specific methods
func (d DVD) Search(query string) bool {
    return contains(d.Title, query) || contains(d.Director, query) || contains(d.Genre, query)
}

func (d DVD) GetInfo() string {
    return fmt.Sprintf("DVD: %s directed by %s (%d min, %s)", 
        d.Title, d.Director, d.Duration, d.Genre)
}

// Helper function
func contains(text, query string) bool {
    return len(query) > 0 && 
           len(text) >= len(query) && 
           text[:len(query)] == query
}
```

### 🏛️ **Step 3: Library Management System**

```go
type Library struct {
    Name  string
    Items []LibraryItem
}

func (l *Library) AddItem(item LibraryItem) {
    l.Items = append(l.Items, item)
}

func (l *Library) SearchItems(query string) []LibraryItem {
    var results []LibraryItem
    for _, item := range l.Items {
        if item.Search(query) {
            results = append(results, item)
        }
    }
    return results
}

func (l *Library) BorrowItem(id string) error {
    for _, item := range l.Items {
        if item.GetID() == id {
            return item.Borrow()
        }
    }
    return fmt.Errorf("item with ID %s not found", id)
}

func (l *Library) ReturnItem(id string) error {
    for _, item := range l.Items {
        if item.GetID() == id {
            return item.Return()
        }
    }
    return fmt.Errorf("item with ID %s not found", id)
}

func (l *Library) ListAvailableItems() {
    fmt.Printf("\n=== Available Items in %s ===\n", l.Name)
    for _, item := range l.Items {
        if item.IsAvailable() {
            switch v := item.(type) {
            case Book:
                fmt.Printf("📖 %s\n", v.GetInfo())
            case DVD:
                fmt.Printf("💿 %s\n", v.GetInfo())
            }
        }
    }
}
```

### 🎮 **Step 4: Demo Usage**

```go
func main() {
    // Create library
    library := Library{Name: "City Central Library"}
    
    // Create items
    book1 := Book{
        Item:   Item{ID: "B001", Title: "The Go Programming Language", Available: true},
        Author: "Alan Donovan",
        ISBN:   "978-0134190440",
        Pages:  380,
    }
    
    book2 := Book{
        Item:   Item{ID: "B002", Title: "Clean Code", Available: true},
        Author: "Robert Martin",
        ISBN:   "978-0132350884",
        Pages:  464,
    }
    
    dvd1 := DVD{
        Item:     Item{ID: "D001", Title: "The Matrix", Available: true},
        Director: "The Wachowskis",
        Duration: 136,
        Genre:    "Sci-Fi",
    }
    
    // Add items to library
    library.AddItem(book1)
    library.AddItem(book2)
    library.AddItem(dvd1)
    
    // Demo operations
    library.ListAvailableItems()
    
    // Search for items
    fmt.Printf("\n=== Search Results for 'Go' ===\n")
    results := library.SearchItems("Go")
    for _, item := range results {
        fmt.Printf("Found: %s\n", item.GetTitle())
    }
    
    // Borrow an item
    fmt.Printf("\n=== Borrowing Operations ===\n")
    if err := library.BorrowItem("B001"); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Successfully borrowed: The Go Programming Language\n")
    }
    
    library.ListAvailableItems()
    
    // Return an item
    if err := library.ReturnItem("B001"); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Successfully returned: The Go Programming Language\n")
    }
    
    library.ListAvailableItems()
}
```

---

## 🎯 Key Takeaways & Best Practices

<div style="background: #d4edda; border: 1px solid #c3e6cb; border-radius: 8px; padding: 20px; margin: 20px 0;">

### ✅ **Structs Best Practices**
- Use meaningful field names and organize related data together
- Consider zero values when designing structs
- Use embedding for composition over inheritance
- Keep structs focused on a single responsibility

### ✅ **Methods Best Practices**
- Use pointer receivers when you need to modify the struct
- Use value receivers for small structs that don't need modification
- Be consistent with receiver types within a type's method set
- Consider performance implications of copying large structs

### ✅ **Interfaces Best Practices**
- Keep interfaces small and focused (single responsibility)
- Define interfaces where they're used, not where they're implemented
- Use composition to build larger interfaces from smaller ones
- Prefer many small interfaces over few large ones

</div>

---

## 🚀 Practice Exercises

### 🎯 **Exercise 1: Student Management System**
Create a student management system with:
- `Student` struct with name, ID, grades
- Methods to add grades, calculate average
- `Printable` interface for displaying student info

### 🎯 **Exercise 2: Shape Calculator**
Implement different shapes (Triangle, Square, Pentagon) that:
- Implement a `Geometry` interface with Area() and Perimeter() methods
- Create a slice of shapes and calculate total area
- Add a `Drawable` interface with a Draw() method

### 🎯 **Exercise 3: Vehicle Fleet**
Design a vehicle fleet system with:
- Base `Vehicle` struct with common properties
- Specific vehicle types (Car, Truck, Motorcycle)
- `Drivable` interface with Start(), Stop(), GetSpeed() methods
- Fleet management functions

---

## 📚 Additional Resources

- [Go by Example - Structs](https://gobyexample.com/structs)
- [Go by Example - Methods](https://gobyexample.com/methods)
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Effective Go - Interfaces](https://golang.org/doc/effective_go.html#interfaces)

---

<div align="center">

**🎉 Congratulations!** 

You've mastered structs, methods, and interfaces in Go. You're now ready to build complex, well-structured applications using Go's powerful type system and composition patterns.

**Next up:** Module 5 - Introduction to Concurrency

</div>
