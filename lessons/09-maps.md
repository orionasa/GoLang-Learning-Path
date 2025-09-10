# 9. Maps in Go

In the previous lesson, you learned about arrays and slices. Now, let's explore another powerful built-in data structure in Go: the map. A map is an unordered collection of key-value pairs.

## What is a Map?

A map is used to store data in a key-value format. Each key in a map is unique. Maps are also sometimes called hash maps or dictionaries in other languages.

## Creating a Map

You can create a map using the `make` function or by using a map literal.

### Using `make`

To create a map with `make`, you specify the key type and the value type.

```go
// Create a map with string keys and int values
ages := make(map[string]int)
```

### Using a Map Literal

You can also create a map using a map literal, which is a more concise way to create and initialize a map.

```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}
```

## Working with Maps

### Adding and Updating Elements

You can add or update elements in a map using the following syntax:

```go
ages["Charlie"] = 35 // Add a new element
ages["Alice"] = 31   // Update an existing element
```

### Accessing Elements

You can access the value associated with a key using the key as an index:

```go
aliceAge := ages["Alice"]
```

If you try to access a key that doesn't exist, you will get the zero value for the value type. For example, if the value type is `int`, you will get `0`.

### Checking for a Key's Existence

To distinguish between a key that has a zero value and a key that doesn't exist, you can use a two-value assignment when accessing a map element. The second value is a boolean that is `true` if the key exists and `false` otherwise.

```go
age, ok := ages["David"]
if ok {
    fmt.Println("David's age is", age)
} else {
    fmt.Println("David's age is not in the map.")
}
```

This is a common idiom in Go for working with maps.

### Deleting Elements

To delete an element from a map, you can use the built-in `delete` function.

```go
delete(ages, "Bob")
```

If the key doesn't exist, `delete` does nothing.

## Iterating Over a Map

You can iterate over the key-value pairs in a map using a `for` loop with the `range` keyword.

```go
for name, age := range ages {
    fmt.Printf("%s is %d years old.\n", name, age)
}
```

Note that the order of iteration over a map is not specified and is not guaranteed to be the same every time you run the program.

## Maps are Reference Types

Like slices, maps are reference types. This means that when you assign a map to a new variable, both variables refer to the same underlying data structure.

```go
m1 := map[string]int{"a": 1}
m2 := m1
m2["a"] = 100

fmt.Println(m1["a"]) // prints 100
```

In the next lesson, we will learn about structs, which allow you to create your own custom data types by grouping together fields of different types.
