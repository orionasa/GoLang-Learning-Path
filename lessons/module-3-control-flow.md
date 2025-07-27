# 🎮 Module 3: Control Flow & Decision Making

<div align="center">

![Go Control Flow](https://via.placeholder.com/400x200/FF6B6B/FFFFFF?text=Go+Control+Flow+%26+Loops)

**Master the Flow of Your Programs!**

*From Zero to Go Hero - Module 3*

</div>

---

## 📋 Module Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 4-5 hours    | Beginner       | Modules 1 & 2 Complete |

### 🎯 Learning Objectives
By the end of this module, you will:
- ✅ Master conditional statements (if, else, switch)
- ✅ Understand and use different types of loops
- ✅ Control program flow with break and continue
- ✅ Handle multiple conditions efficiently
- ✅ Write clean, readable decision-making code

---

## 🤔 What is Control Flow?

<div style="background: linear-gradient(135deg, #FF6B6B 0%, #4ECDC4 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 Think of Control Flow as Traffic Lights

Control flow determines **which code runs when** and **how many times**. Just like traffic lights control the flow of cars, control flow statements control the execution of your code!

</div>

### 🚦 Control Flow Visualization

```
┌─────────────────────────────────────┐
│  Program Execution Flow             │
├─────────────────────────────────────┤
│  Start ──→ Condition? ──→ Action A  │
│              │                      │
│              ↓                      │
│           Action B ──→ Loop? ──→ End│
│                         │           │
│                         ↓           │
│                    Repeat ──────────┘
└─────────────────────────────────────┘
```

---

## 🔀 Conditional Statements: Making Decisions

### 🎯 The `if` Statement

The most basic decision-making tool in programming!

<div style="background: #e8f5e8; border-left: 4px solid #4caf50; padding: 15px; margin: 15px 0;">
<strong>💡 Real-World Analogy:</strong><br>
"<em>If it's raining, take an umbrella</em>" - This is exactly how if statements work!
</div>

#### **Basic if Statement**

```go
package main

import "fmt"

func main() {
    age := 18
    
    if age >= 18 {
        fmt.Println("You can vote! 🗳️")
    }
}
```

#### **if-else Statement**

```go
package main

import "fmt"

func main() {
    temperature := 25
    
    if temperature > 30 {
        fmt.Println("It's hot! 🌞 Stay hydrated!")
    } else {
        fmt.Println("Nice weather! 🌤️ Perfect for a walk!")
    }
}
```

#### **if-else if-else Chain**

```go
package main

import "fmt"

func main() {
    score := 85
    
    if score >= 90 {
        fmt.Println("Grade: A+ 🌟 Excellent!")
    } else if score >= 80 {
        fmt.Println("Grade: B+ 👍 Great job!")
    } else if score >= 70 {
        fmt.Println("Grade: C+ 📚 Good work!")
    } else if score >= 60 {
        fmt.Println("Grade: D 📖 Keep studying!")
    } else {
        fmt.Println("Grade: F 💪 Don't give up!")
    }
}
```

### 🎪 Advanced if Features

#### **Multiple Conditions with Logical Operators**

```go
package main

import "fmt"

func main() {
    age := 25
    hasLicense := true
    hasInsurance := true
    
    // AND operator (&&)
    if age >= 18 && hasLicense && hasInsurance {
        fmt.Println("You can drive! 🚗")
    }
    
    // OR operator (||)
    weekend := true
    holiday := false
    
    if weekend || holiday {
        fmt.Println("Time to relax! 🏖️")
    }
    
    // NOT operator (!)
    isRaining := false
    
    if !isRaining {
        fmt.Println("Perfect day for outdoor activities! ☀️")
    }
}
```

#### **Short Variable Declaration in if**

```go
package main

import "fmt"

func main() {
    // Declare and use variable in the same if statement
    if num := 42; num > 0 {
        fmt.Printf("%d is positive! ✅\n", num)
    }
    // 'num' is only available within this if block
}
```

---

## 🔄 Switch Statements: Elegant Multiple Choice

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 Why Switch?

When you have many conditions to check, `switch` is cleaner and more readable than long `if-else` chains!

</div>

### 🎯 Basic Switch Statement

```go
package main

import "fmt"

func main() {
    day := "Monday"
    
    switch day {
    case "Monday":
        fmt.Println("Start of the work week! 💼")
    case "Tuesday":
        fmt.Println("Tuesday blues 😴")
    case "Wednesday":
        fmt.Println("Hump day! 🐪")
    case "Thursday":
        fmt.Println("Almost there! 🎯")
    case "Friday":
        fmt.Println("TGIF! 🎉")
    case "Saturday", "Sunday":
        fmt.Println("Weekend vibes! 🏖️")
    default:
        fmt.Println("Invalid day 🤔")
    }
}
```

### 🎪 Advanced Switch Features

#### **Switch with Expressions**

```go
package main

import "fmt"

func main() {
    score := 85
    
    switch {
    case score >= 90:
        fmt.Println("Outstanding! 🌟")
    case score >= 80:
        fmt.Println("Excellent! 👏")
    case score >= 70:
        fmt.Println("Good! 👍")
    case score >= 60:
        fmt.Println("Pass 📝")
    default:
        fmt.Println("Need improvement 📚")
    }
}
```

#### **Switch with Short Declaration**

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    switch hour := time.Now().Hour(); {
    case hour < 12:
        fmt.Println("Good morning! ☀️")
    case hour < 17:
        fmt.Println("Good afternoon! 🌤️")
    default:
        fmt.Println("Good evening! 🌙")
    }
}
```

---

## 🔁 Loops: Repeating Actions

<div style="background: linear-gradient(135deg, #4ECDC4 0%, #44A08D 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 The Power of Repetition

Loops let you repeat code without writing it multiple times. Imagine having to write "Happy Birthday" 100 times manually! 😱

</div>

### 🎯 The `for` Loop (Go's Only Loop!)

Go has only one loop keyword: `for`. But it's incredibly versatile!

#### **1. Traditional for Loop**

```go
package main

import "fmt"

func main() {
    fmt.Println("Countdown! 🚀")
    
    for i := 10; i >= 1; i-- {
        fmt.Printf("%d... ", i)
    }
    fmt.Println("Blast off! 🚀")
}
```

#### **2. While-style Loop**

```go
package main

import "fmt"

func main() {
    count := 0
    
    for count < 5 {
        fmt.Printf("Count: %d 📊\n", count)
        count++
    }
}
```

#### **3. Infinite Loop**

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    counter := 0
    
    for {
        fmt.Printf("Running... %d ⚡\n", counter)
        counter++
        
        if counter >= 3 {
            break // Exit the loop
        }
        
        time.Sleep(1 * time.Second)
    }
    fmt.Println("Loop finished! ✅")
}
```

### 🎪 Advanced Loop Techniques

#### **Looping Through Ranges**

```go
package main

import "fmt"

func main() {
    // Loop through numbers
    fmt.Println("Numbers 1-5:")
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // Loop through a string
    word := "Hello"
    fmt.Println("Characters in 'Hello':")
    for i := 0; i < len(word); i++ {
        fmt.Printf("%c ", word[i])
    }
    fmt.Println()
}
```

#### **Nested Loops**

```go
package main

import "fmt"

func main() {
    fmt.Println("Multiplication Table (1-5):")
    fmt.Println("========================")
    
    for i := 1; i <= 5; i++ {
        for j := 1; j <= 5; j++ {
            fmt.Printf("%2d ", i*j)
        }
        fmt.Println()
    }
}
```

---

## 🎮 Loop Control: break & continue

### 🛑 The `break` Statement

Exits the loop immediately!

```go
package main

import "fmt"

func main() {
    fmt.Println("Finding the first even number:")
    
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Printf("Found it! %d is even! 🎯\n", i)
            break // Exit the loop
        }
        fmt.Printf("%d is odd, continuing...\n", i)
    }
}
```

### ⏭️ The `continue` Statement

Skips the current iteration and continues with the next one!

```go
package main

import "fmt"

func main() {
    fmt.Println("Odd numbers from 1 to 10:")
    
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        fmt.Printf("%d ", i)
    }
    fmt.Println()
}
```

---

## 🎯 Practical Examples

### 🎲 Example 1: Number Guessing Game

```go
package main

import "fmt"

func main() {
    secretNumber := 7
    maxAttempts := 3
    
    fmt.Println("🎲 Number Guessing Game!")
    fmt.Println("I'm thinking of a number between 1 and 10")
    
    for attempt := 1; attempt <= maxAttempts; attempt++ {
        var guess int
        fmt.Printf("Attempt %d: Enter your guess: ", attempt)
        fmt.Scan(&guess)
        
        if guess == secretNumber {
            fmt.Println("🎉 Congratulations! You guessed it!")
            break
        } else if guess < secretNumber {
            fmt.Println("📈 Too low!")
        } else {
            fmt.Println("📉 Too high!")
        }
        
        if attempt == maxAttempts {
            fmt.Printf("😔 Game over! The number was %d\n", secretNumber)
        }
    }
}
```

### 📊 Example 2: Grade Calculator

```go
package main

import "fmt"

func main() {
    var totalScore, numSubjects int
    
    fmt.Println("📊 Grade Calculator")
    fmt.Print("How many subjects? ")
    fmt.Scan(&numSubjects)
    
    for i := 1; i <= numSubjects; i++ {
        var score int
        fmt.Printf("Enter score for subject %d: ", i)
        fmt.Scan(&score)
        
        // Validate score
        if score < 0 || score > 100 {
            fmt.Println("❌ Invalid score! Please enter 0-100")
            i-- // Retry this subject
            continue
        }
        
        totalScore += score
    }
    
    average := float64(totalScore) / float64(numSubjects)
    
    fmt.Printf("\n📈 Results:\n")
    fmt.Printf("Total Score: %d\n", totalScore)
    fmt.Printf("Average: %.2f\n", average)
    
    // Determine grade
    switch {
    case average >= 90:
        fmt.Println("Grade: A+ 🌟 Outstanding!")
    case average >= 80:
        fmt.Println("Grade: B+ 👏 Excellent!")
    case average >= 70:
        fmt.Println("Grade: C+ 👍 Good!")
    case average >= 60:
        fmt.Println("Grade: D 📚 Pass")
    default:
        fmt.Println("Grade: F 💪 Keep trying!")
    }
}
```

### 🎨 Example 3: Pattern Printer

```go
package main

import "fmt"

func main() {
    fmt.Println("🎨 Pattern Printer")
    
    // Triangle pattern
    fmt.Println("\n🔺 Triangle Pattern:")
    for i := 1; i <= 5; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("⭐ ")
        }
        fmt.Println()
    }
    
    // Diamond pattern
    fmt.Println("\n💎 Diamond Pattern:")
    // Upper half
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3-i; j++ {
            fmt.Print("  ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("💎 ")
        }
        fmt.Println()
    }
    // Lower half
    for i := 2; i >= 1; i-- {
        for j := 1; j <= 3-i; j++ {
            fmt.Print("  ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("💎 ")
        }
        fmt.Println()
    }
}
```

---

## 🧠 Best Practices & Tips

<div style="background: #fff3cd; border: 1px solid #ffeaa7; border-radius: 8px; padding: 15px; margin: 15px 0;">

### 💡 **Pro Tips for Clean Control Flow**

1. **Keep conditions simple and readable**
   ```go
   // Good ✅
   if age >= 18 && hasLicense {
       // code
   }
   
   // Avoid ❌
   if !(age < 18 || !hasLicense) {
       // code
   }
   ```

2. **Use early returns to reduce nesting**
   ```go
   // Good ✅
   if !isValid {
       return
   }
   // Continue with main logic
   
   // Avoid ❌
   if isValid {
       // Long nested code
   }
   ```

3. **Prefer switch over long if-else chains**
   ```go
   // Good ✅
   switch status {
   case "active", "pending":
       // handle
   }
   
   // Avoid ❌
   if status == "active" || status == "pending" {
       // handle
   }
   ```

</div>

### 🚨 Common Pitfalls to Avoid

<div style="background: #f8d7da; border: 1px solid #f5c6cb; border-radius: 8px; padding: 15px; margin: 15px 0;">

1. **Infinite loops without exit conditions**
2. **Off-by-one errors in loop conditions**
3. **Forgetting to update loop variables**
4. **Using assignment (=) instead of comparison (==)**

</div>

---

## 🎯 Practice Exercises

<div style="background: linear-gradient(135deg, #84fab0 0%, #8fd3f4 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🏋️ Challenge Yourself!

Try these exercises to master control flow:

</div>

### 🥉 **Beginner Level**

1. **Age Classifier**: Write a program that classifies people as child (0-12), teenager (13-19), adult (20-64), or senior (65+)

2. **Even/Odd Counter**: Count and display even and odd numbers from 1 to 20

3. **Simple Calculator**: Create a calculator that performs basic operations based on user choice

### 🥈 **Intermediate Level**

4. **Prime Number Checker**: Check if a given number is prime

5. **Fibonacci Sequence**: Generate the first 10 numbers in the Fibonacci sequence

6. **Password Validator**: Validate passwords based on length, special characters, etc.

### 🥇 **Advanced Level**

7. **Menu-Driven Program**: Create a program with multiple options and loop until user chooses to exit

8. **Number Pattern Generator**: Create various number patterns based on user input

9. **Simple ATM Simulator**: Simulate basic ATM operations with balance checking and transactions

---

## 🎉 Module Summary

<div style="background: #d1ecf1; border: 1px solid #bee5eb; border-radius: 8px; padding: 15px;">

### ✅ **What You've Mastered**

- **Conditional Logic**: if, else, switch statements
- **Loops**: for loops in all their forms
- **Flow Control**: break and continue statements
- **Best Practices**: Writing clean, readable control flow code
- **Real Applications**: Practical examples and exercises

</div>

### 🚀 **What's Next?**

<div style="background: #d4edda; border: 1px solid #c3e6cb; border-radius: 8px; padding: 15px;">

In **Module 4**, we'll explore:
- Functions and methods
- Parameters and return values
- Scope and visibility
- Advanced function concepts

</div>

---

<div align="center">

**🎯 Ready for the next challenge?**

*Continue your Go journey with Module 4!*

---

*Happy Coding! 🚀*

</div>
