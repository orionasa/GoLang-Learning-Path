# 8. Arrays and Slices in Go

In Go, arrays and slices are used to store ordered collections of elements of the same type. While they are related, they have some important differences.

## Arrays

An array is a fixed-size collection of elements. The size of an array is part of its type. This means that `[5]int` and `[10]int` are two different array types.

### Declaring an Array

You can declare an array by specifying the number of elements and the type of the elements.

```go
var numbers [5]int // declares an array of 5 integers
```

You can also declare and initialize an array at the same time:

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

You can let Go infer the size of the array by using `...` instead of a number:

```go
numbers := [...]int{1, 2, 3, 4, 5} // size is inferred to be 5
```

### Accessing Array Elements

You can access the elements of an array using an index, starting from 0.

```go
first := numbers[0] // get the first element
numbers[1] = 10     // set the second element to 10
```

### Arrays are Value Types

In Go, arrays are value types. This means that when you assign an array to a new variable or pass it to a function, a copy of the array is created.

```go
a := [...]int{1, 2, 3}
b := a // b is a copy of a
b[0] = 100

fmt.Println(a) // prints [1 2 3]
fmt.Println(b) // prints [100 2 3]
```

Because of their fixed size and value semantics, arrays are not used as often in Go as slices.

## Slices

A slice is a more flexible and powerful way to work with sequences of data. A slice is a lightweight data structure that provides a view into an underlying array.

A slice has three components:
*   A pointer to the underlying array
*   The length of the slice
*   The capacity of the slice

The length is the number of elements in the slice, while the capacity is the number of elements in the underlying array, starting from the first element of the slice.

### Creating a Slice

You can create a slice from an existing array or another slice.

```go
numbers := [...]int{1, 2, 3, 4, 5}
slice := numbers[1:4] // creates a slice from index 1 to 3 (exclusive)
// slice will be [2 3 4]
```

You can also create a slice literal, which creates an underlying array and returns a slice that refers to it.

```go
letters := []string{"a", "b", "c"}
```

### The `make` Function

You can also create a slice using the built-in `make` function. This is useful when you want to create a slice with a specific length and capacity.

```go
s := make([]int, 5)    // length 5, capacity 5
s = make([]int, 5, 10) // length 5, capacity 10
```

### Slices are Reference Types

Unlike arrays, slices are reference types. This means that when you assign a slice to a new variable, both variables will refer to the same underlying array.

```go
s1 := []int{1, 2, 3}
s2 := s1 // s2 refers to the same underlying array as s1
s2[0] = 100

fmt.Println(s1) // prints [100 2 3]
fmt.Println(s2) // prints [100 2 3]
```

### The `append` Function

The built-in `append` function is used to add elements to a slice. `append` returns a new slice that may or may not refer to the same underlying array.

```go
s := []int{1, 2, 3}
s = append(s, 4, 5) // s is now [1 2 3 4 5]
```

If the capacity of the original slice is not large enough, `append` will allocate a new, larger underlying array and return a slice that refers to it.

## Conclusion

Arrays and slices are fundamental data structures in Go. While arrays have a fixed size, slices provide a more flexible and powerful way to work with sequences of data. In most cases, you will find yourself using slices instead of arrays.

In the next lesson, we will explore maps, which are used to store key-value pairs.
