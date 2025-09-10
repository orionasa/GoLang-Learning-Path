# 4. Variables and Data Types in Go

In the previous lesson, you learned about the basic syntax of a Go program. Now, let's explore one of the most fundamental concepts in any programming language: variables and data types.

## What is a Variable?

A variable is a named storage location that holds a value. In Go, you can declare a variable and assign a value to it.

## Declaring Variables

In Go, you can declare a variable using the `var` keyword, followed by the variable name and its type.

```go
var name string = "Go"
var version int = 1
```

If you provide an initial value, Go can infer the type of the variable, so you can omit the type declaration:

```go
var name = "Go" // type string is inferred
var version = 1 // type int is inferred
```

### Short Variable Declaration

Go provides a shorter way to declare and initialize variables using the `:=` operator. This is known as the short variable declaration.

```go
name := "Go"     // equivalent to var name string = "Go"
version := 1   // equivalent to var version int = 1
```

The short variable declaration is only available inside functions. Outside of a function, every statement begins with a keyword (like `var`, `func`, etc.).

## Basic Data Types

Go has several built-in data types. Let's look at some of the most common ones.

### Integers

Integers are whole numbers. Go has several integer types, such as `int`, `int8`, `int16`, `int32`, and `int64`. The `int` type is the most common and has a size that is appropriate for the underlying platform (32 or 64 bits).

```go
var age int = 30
var temperature = -5
```

There are also unsigned integer types, such as `uint`, `uint8`, etc., which can only hold non-negative values.

### Floating-Point Numbers

Floating-point numbers are numbers with a decimal point. Go has two floating-point types: `float32` and `float64`. `float64` is the default and is generally recommended for most cases.

```go
var price float64 = 99.95
pi := 3.14159
```

### Booleans

A boolean type, declared with the `bool` keyword, can have one of two values: `true` or `false`.

```go
var isLoggedIn bool = true
var isEnabled = false
```

### Strings

A string is a sequence of characters. In Go, strings are enclosed in double quotes (`"`).

```go
var message string = "Hello, World!"
greeting := "Welcome to Go"
```

Strings in Go are immutable, which means that once a string is created, its value cannot be changed.

## Zero Values

In Go, variables that are declared without an explicit initial value are given a "zero value". The zero value is:

*   `0` for numeric types (int, float, etc.)
*   `false` for the boolean type
*   `""` (the empty string) for strings

For example:

```go
var i int
var f float64
var b bool
var s string

// i will be 0
// f will be 0.0
// b will be false
// s will be ""
```

This ensures that variables always have a well-defined value, which helps in preventing bugs.

In the next lesson, we will learn about control flow statements, such as `if`, `for`, and `switch`, which allow you to control the execution of your code.
