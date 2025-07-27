# Module 3: Collections and Functions

In this module, we'll learn about collections (arrays, slices, and maps) and functions in Go.

## Arrays, Slices, and their relationship

### Arrays

An array is a fixed-size collection of elements of the same type. The size of an array is part of its type.

```go
var numbers [5]int // declares an array of 5 integers
numbers[0] = 1
numbers[1] = 2
```

### Slices

A slice is a dynamically-sized, flexible view into the elements of an array. A slice is more common than an array in Go.

```go
numbers := []int{1, 2, 3, 4, 5} // creates a slice of integers
```

Slices have a length and a capacity. The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

You can create a new slice from an existing slice using the `[low:high]` syntax.

```go
numbers := []int{1, 2, 3, 4, 5}
subSlice := numbers[1:3] // contains [2, 3]
```

### Relationship

A slice does not store any data, it just describes a section of an underlying array. When you create a slice, you are creating a new slice value that points to the original array.

## Maps for key-value storage

A map is a collection of key-value pairs. The keys in a map must be unique.

```go
// creates a map with string keys and int values
ages := make(map[string]int)

ages["Alice"] = 30
ages["Bob"] = 25

fmt.Println(ages["Alice"]) // prints 30
```

You can also create a map using a map literal:

```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}
```

To delete a key-value pair from a map, use the `delete` function:

```go
delete(ages, "Alice")
```

## Defining and Calling Functions

A function is a block of code that performs a specific task. You can define a function using the `func` keyword.

```go
func add(x int, y int) int {
    return x + y
}
```

This function takes two integer arguments and returns their sum. You can call this function like this:

```go
result := add(2, 3) // result is 5
```

## Multiple Return Values and Variadic Functions

### Multiple Return Values

A function can return multiple values. This is often used in Go to return both a result and an error value.

```go
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return x / y, nil
}
```

### Variadic Functions

A variadic function is a function that can take a variable number of arguments. The `fmt.Println` function is a common example of a variadic function.

```go
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
```

You can call this function with any number of integer arguments:

```go
sum(1, 2, 3) // returns 6
sum(1, 2)    // returns 3
```

## The `error` type for basic error handling

In Go, it is idiomatic to communicate errors via an explicit, separate return value. The `error` type is a built-in interface.

```go
import (
    "fmt"
    "errors"
)

func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return x / y, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

This module has covered the essentials of collections and functions in Go. You are now equipped to write more complex and structured programs.

**Previous:** [Module 2: Basic Syntax and Data Types](../2/README.md)
**Next:** [Module 4: Structs, Methods & Interfaces](../4/README.md)
