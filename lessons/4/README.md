# Module 4: Structs, Methods & Interfaces

In this module, we'll explore Go's support for object-oriented programming, including structs, methods, and interfaces.

## Composite Types: Structs and Pointers

### Structs

A struct is a composite data type that groups together zero or more values of different types. Structs are a way to create more complex data types.

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name) // prints "Alice"
}
```

### Pointers

A pointer is a variable that stores the memory address of another variable. Pointers are useful for passing large data structures to functions without copying them.

```go
p := Person{Name: "Alice", Age: 30}
pPtr := &p // pPtr is a pointer to p

pPtr.Age = 31 // changes the Age field of p
fmt.Println(p.Age) // prints 31
```

## Defining Methods on Structs

A method is a function that is associated with a particular type. You can define a method on a struct by specifying a "receiver" for the function.

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    r := Rectangle{Width: 10, Height: 5}
    fmt.Println(r.Area()) // prints 50
}
```

### Pointer vs. Value Receivers

When you define a method, you can use either a value receiver or a pointer receiver.

*   A **value receiver** creates a copy of the value, so the method cannot modify the original value.
*   A **pointer receiver** passes a pointer to the value, so the method can modify the original value.

```go
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    r := Rectangle{Width: 10, Height: 5}
    r.Scale(2)
    fmt.Println(r.Width, r.Height) // prints 20 10
}
```

## Interfaces: Defining and Implementing Behavior

An interface is a collection of method signatures. A type implements an interface by implementing all of the methods in the interface.

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    r := Rectangle{Width: 10, Height: 5}
    c := Circle{Radius: 5}

    PrintArea(r) // prints "Area: 50"
    PrintArea(c) // prints "Area: 78.539..."
}
```

## The Empty Interface (`interface{}`) and Type Assertions

The empty interface, `interface{}`, is an interface that has no methods. Since every type has zero or more methods, every type implements the empty interface. This means you can store a value of any type in a variable of type `interface{}`.

A type assertion is used to extract the underlying value from an interface value.

```go
var i interface{} = "hello"

s := i.(string)
fmt.Println(s) // prints "hello"

s, ok := i.(string)
fmt.Println(s, ok) // prints "hello true"

f, ok := i.(float64)
fmt.Println(f, ok) // prints "0 false"
```

## Go's philosophy: Composition over Inheritance

Go does not support inheritance in the traditional sense. Instead, it encourages composition, where you build complex types by combining simpler types. This is achieved through struct embedding.

```go
type Engine struct {
    Horsepower int
}

type Car struct {
    Engine
    Make  string
    Model string
}

func main() {
    c := Car{
        Engine: Engine{Horsepower: 300},
        Make:   "Ford",
        Model:  "Mustang",
    }
    fmt.Println(c.Horsepower) // prints 300
}
```

This module has introduced you to the key concepts of object-oriented programming in Go. You can now create your own custom data types and define their behavior using structs, methods, and interfaces.

**Previous:** [Module 3: Collections and Functions](../3/README.md)
**Next:** [Module 5: Introduction to Concurrency](../5/README.md)
