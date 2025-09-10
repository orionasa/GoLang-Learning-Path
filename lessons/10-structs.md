# 10. Structs in Go

In Go, a struct is a composite data type that groups together zero or more fields of different types into a single unit. Structs are useful for creating your own custom data types.

## Defining a Struct

You can define a struct using the `type` and `struct` keywords.

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
```

In this example, we've defined a `Person` struct with three fields: `FirstName`, `LastName`, and `Age`.

## Creating a Struct Instance

You can create an instance of a struct in several ways.

### Using a Struct Literal

You can create a struct instance using a struct literal:

```go
p1 := Person{
    FirstName: "Alice",
    LastName:  "Smith",
    Age:       30,
}
```

You can also omit the field names, but you must provide the values in the order they are defined in the struct. This is less readable and generally not recommended.

```go
p2 := Person{"Bob", "Johnson", 25}
```

### Using the `new` Keyword

You can also create a pointer to a new struct instance using the `new` keyword. This will allocate memory for the struct and return a pointer to it. The fields will be initialized to their zero values.

```go
p3 := new(Person) // p3 is a pointer to a Person struct
```

## Accessing Struct Fields

You can access the fields of a struct using the dot (`.`) operator.

```go
fmt.Println(p1.FirstName) // prints "Alice"

p1.Age = 31 // update the Age field
```

If you have a pointer to a struct, you can still use the dot operator to access its fields. Go automatically dereferences the pointer for you.

```go
p3.FirstName = "Charlie"
```

## Methods on Structs

In Go, you can define methods on structs. A method is a function with a special "receiver" argument. The receiver appears in its own argument list between the `func` keyword and the method name.

```go
// A method on the Person struct
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}
```

You can call the method on an instance of the `Person` struct:

```go
fmt.Println(p1.FullName()) // prints "Alice Smith"
```

### Pointer Receivers

You can also define methods with a pointer receiver. This is useful when you want to modify the struct within the method.

```go
func (p *Person) SetAge(age int) {
    p.Age = age
}
```

When you call a method with a pointer receiver, Go will automatically take the address of the struct instance for you.

```go
p1.SetAge(32) // p1.Age is now 32
```

The general rule is to use a pointer receiver if the method needs to modify the receiver. If not, a value receiver is fine.

## Embedded Structs (Composition)

Go does not have inheritance in the traditional sense. Instead, it uses composition through embedded structs. You can embed one struct within another to "inherit" its fields and methods.

```go
type Employee struct {
    Person // embedded Person struct
    Salary int
}

e := Employee{
    Person: Person{FirstName: "David", LastName: "Jones", Age: 40},
    Salary: 50000,
}

fmt.Println(e.FirstName)    // prints "David" (promoted field)
fmt.Println(e.FullName()) // prints "David Jones" (promoted method)
```

In the next lesson, we will learn about interfaces, which provide a way to specify the behavior of an object.
