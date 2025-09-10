# 5. Control Flow in Go

In the previous lesson, you learned about variables and data types. Now, let's look at how you can control the flow of execution in your Go programs using control flow statements.

## `if` Statements

The `if` statement is used to execute a block of code only if a certain condition is true.

```go
package main

import "fmt"

func main() {
    age := 18

    if age >= 18 {
        fmt.Println("You are an adult.")
    }
}
```

You can also include an `else` block, which will be executed if the condition is false.

```go
if age >= 18 {
    fmt.Println("You are an adult.")
} else {
    fmt.Println("You are a minor.")
}
```

You can also use `else if` to check for multiple conditions.

```go
if age >= 18 {
    fmt.Println("You are an adult.")
} else if age >= 13 {
    fmt.Println("You are a teenager.")
} else {
    fmt.Println("You are a child.")
}
```

Go also allows you to declare a variable within the `if` statement itself. This variable is only accessible within the `if` and `else` blocks.

```go
if n := 10; n%2 == 0 {
    fmt.Println("n is an even number.")
} else {
    fmt.Println("n is an odd number.")
}
```

## `for` Loops

The `for` loop is the only looping construct in Go. It can be used in several ways.

### The Basic `for` Loop

The basic `for` loop has three components separated by semicolons:
1.  The init statement: executed before the first iteration.
2.  The condition expression: evaluated before every iteration.
3.  The post statement: executed at the end of every iteration.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### The `while` Loop (using `for`)

In Go, you can create a "while" loop by using the `for` loop with only a condition.

```go
n := 0
for n < 5 {
    fmt.Println(n)
    n++
}
```

### The Infinite Loop

You can create an infinite loop by using the `for` loop without any components.

```go
for {
    fmt.Println("This will run forever!")
}
```

You can use the `break` statement to exit a loop.

### `range` with `for` Loops

The `range` keyword can be used with a `for` loop to iterate over elements in a slice or map. We will cover slices and maps in a later lesson.

## `switch` Statements

The `switch` statement is a shorter way to write a sequence of `if-else` statements. It runs the first case whose value is equal to the condition expression.

```go
day := "Sunday"

switch day {
case "Monday":
    fmt.Println("It's Monday.")
case "Tuesday":
    fmt.Println("It's Tuesday.")
case "Wednesday":
    fmt.Println("It's Wednesday.")
case "Thursday":
    fmt.Println("It's Thursday.")
case "Friday":
    fmt.Println("It's Friday.")
case "Saturday", "Sunday":
    fmt.Println("It's the weekend!")
default:
    fmt.Println("It's some other day.")
}
```

A few things to note about `switch` statements in Go:

*   You don't need a `break` statement at the end of each case. The `break` is implicit.
*   You can have multiple expressions in the same `case` statement, separated by commas.
*   The `default` case is executed if none of the other cases match.

You can also use a `switch` statement without an expression. This is equivalent to `switch true` and is a clean way to write a long `if-else-if` chain.

```go
age := 25

switch {
case age < 18:
    fmt.Println("You are a minor.")
case age >= 18 && age < 65:
    fmt.Println("You are an adult.")
default:
    fmt.Println("You are a senior.")
}
```

Congratulations! You have now learned the basics of control flow in Go. In the next set of lessons, we will dive deeper into Go's features, starting with functions.
