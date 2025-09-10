# 6. Functions in Go

In Go, a function is a block of code that performs a specific task. You've already seen the `main` function, which is the entry point of a Go program. In this lesson, we'll explore how to define and use your own functions.

## Defining a Function

You define a function using the `func` keyword, followed by the function name, a list of parameters, the return type, and the function body.

```go
func add(a int, b int) int {
    return a + b
}
```

In this example, we've defined a function named `add` that takes two integer parameters, `a` and `b`, and returns an integer.

If multiple consecutive parameters have the same type, you can omit the type from all but the last one:

```go
func add(a, b int) int {
    return a + b
}
```

## Calling a Function

To call a function, you simply use its name followed by the arguments in parentheses:

```go
result := add(5, 3) // result will be 8
```

## Multiple Return Values

A unique feature of Go is that functions can return multiple values. This is often used to return both a result and an error value.

```go
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

Here's how you would call this function and handle the return values:

```go
result, err := divide(10, 2)
if err != nil {
    // handle the error
} else {
    // use the result
}
```

## Named Return Values

Go functions can have named return values. When a function has named return values, they are treated as variables defined at the top of the function.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

While this can be useful in short functions, it's generally better to be explicit and return the values directly, as it can harm readability in longer functions.

## Variadic Functions

A function that can be called with a varying number of arguments is called a variadic function. To create a variadic function, you use an ellipsis (`...`) before the type of the last parameter.

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

You can then call this function with any number of integer arguments:

```go
sum(1, 2)       // returns 3
sum(1, 2, 3, 4) // returns 10
```

## Anonymous Functions (Closures)

Go also supports anonymous functions, which are functions without a name. These can be useful when you want to define a function inline without having to name it.

Anonymous functions are also closures, which means they can access the variables from the scope in which they are defined.

```go
func main() {
    add := func(a, b int) int {
        return a + b
    }

    fmt.Println(add(3, 4)) // prints 7
}
```

In the next lesson, we'll look at how to organize your code into packages.
