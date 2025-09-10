# 11. Interfaces in Go

In the previous lesson, you learned about structs, which allow you to create custom data types. Now, we will explore interfaces, which are a powerful feature in Go for achieving polymorphism and writing flexible, decoupled code.

## What is an Interface?

An interface is a type that defines a set of methods. Any type that implements all the methods of an interface is said to satisfy that interface.

In Go, interfaces are implemented implicitly. There is no `implements` keyword. If a type has all the methods of an interface, it automatically satisfies that interface.

## Defining an Interface

You can define an interface using the `type` and `interface` keywords.

```go
// Shape is an interface for types that have an Area method.
type Shape interface {
    Area() float64
}
```

In this example, we've defined a `Shape` interface with a single method, `Area`, which takes no arguments and returns a `float64`.

## Implementing an Interface

Now, let's create a couple of types that implement the `Shape` interface.

```go
import "math"

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

Both `Rectangle` and `Circle` have an `Area()` method with the same signature as the one defined in the `Shape` interface. Therefore, both types implicitly satisfy the `Shape` interface.

## Using Interfaces

Interfaces allow you to write functions that can work with any type that satisfies the interface. This is a form of polymorphism.

Let's create a function that takes a `Shape` as an argument and prints its area.

```go
func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}
```

Now, you can call this function with both a `Rectangle` and a `Circle`:

```go
r := Rectangle{Width: 10, Height: 5}
c := Circle{Radius: 3}

PrintArea(r) // prints "Area: 50.00"
PrintArea(c) // prints "Area: 28.27"
```

## The Empty Interface

The empty interface, `interface{}`, is a special interface that has no methods. Since every type has zero or more methods, every type satisfies the empty interface.

The empty interface is used when you need to work with values of unknown type. For example, the `fmt.Println` function takes arguments of type `interface{}`.

```go
var i interface{}
i = 42
i = "hello"
```

To work with a value of type `interface{}`, you need to use a type assertion or a type switch to determine its underlying type.

### Type Assertions

A type assertion provides access to an interface value's underlying concrete value.

```go
var i interface{} = "hello"

s := i.(string)
fmt.Println(s) // prints "hello"
```

If the assertion fails (i.e., the underlying type is not what you asserted), the program will panic. To handle this gracefully, you can use a two-value assignment:

```go
s, ok := i.(string)
if ok {
    // ...
}
```

### Type Switches

A type switch is like a regular `switch` statement, but it operates on the type of an interface value.

```go
func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}
```

Interfaces are a cornerstone of Go's design and are used extensively in the standard library and idiomatic Go code. They enable you to write flexible, reusable, and decoupled code.

This concludes our intermediate-level lessons. In the next section, we will dive into more advanced topics, starting with concurrency.
