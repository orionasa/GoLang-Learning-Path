# Module 2: Basic Syntax and Data Types

In this module, we'll explore the fundamental building blocks of Go programming: packages, imports, variables, constants, data types, and control flow.

## Packages and Imports

In Go, code is organized into packages. A package is a collection of source files in the same directory that are compiled together. The `main` package is a special package that defines a standalone executable program.

To use code from another package, you need to import it using the `import` keyword. For example, to use the `fmt` package for printing to the console, you would write:

```go
import "fmt"
```

You can also import multiple packages at once:

```go
import (
    "fmt"
    "math"
)
```

## Variables, Constants, and Zero Values

### Variables

A variable is a named storage location for a value. In Go, you can declare a variable using the `var` keyword, followed by the variable name and type:

```go
var age int // declares a variable named age of type int
```

You can also initialize a variable at the time of declaration:

```go
var name string = "Alice" // declares and initializes a variable
```

Go also supports a short variable declaration syntax, which can be used within functions:

```go
name := "Alice" // declares and initializes a variable (type is inferred)
```

### Constants

A constant is a variable whose value cannot be changed. In Go, you can declare a constant using the `const` keyword:

```go
const pi = 3.14159
```

### Zero Values

If a variable is declared without an initial value, it is given a "zero value". The zero value for a numeric type is `0`, for a boolean type is `false`, and for a string type is `""` (an empty string).

## Basic Data Types

Go has a number of built-in data types, including:

*   **`int`:** A signed integer type. The size of an `int` is platform-dependent (either 32 or 64 bits).
*   **`float64`:** A 64-bit floating-point number.
*   **`string`:** A sequence of characters.
*   **`bool`:** A boolean value, either `true` or `false`.

Here are some examples of how to use these data types:

```go
var age int = 30
var temperature float64 = 98.6
var name string = "Bob"
var isStudent bool = true
```

## Control Flow

Control flow statements are used to control the order in which code is executed. Go has three control flow statements: `if/else`, `for`, and `switch`.

### `if/else`

The `if/else` statement is used to execute a block of code if a certain condition is true.

```go
if age >= 18 {
    fmt.Println("You are an adult.")
} else {
    fmt.Println("You are a minor.")
}
```

### `for`

The `for` loop is used to execute a block of code multiple times. Go has only one looping construct, the `for` loop.

```go
// A classic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// A "while" loop
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// An infinite loop
for {
    fmt.Println("This will run forever!")
}
```

### `switch`

The `switch` statement is used to select one of many code blocks to be executed.

```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("It's Monday!")
case "Tuesday":
    fmt.Println("It's Tuesday!")
default:
    fmt.Println("It's some other day.")
}
```

That's a wrap for Module 2! You now have a solid understanding of Go's basic syntax and data types.

**Previous:** [Module 1: Introduction & Setup](../1/README.md)
**Next:** [Module 3: Collections and Functions](../3/README.md)
