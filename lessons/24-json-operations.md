# 24. JSON Operations in Go

JSON (JavaScript Object Notation) is a lightweight data-interchange format that is easy for humans to read and write, and easy for machines to parse and generate. Go has excellent built-in support for working with JSON in the `encoding/json` package.

## Encoding JSON (Marshalling)

Encoding, or "marshalling," is the process of converting a Go data structure into a JSON string. The `json.Marshal()` function is used for this.

Let's take a Go struct and marshal it into JSON:

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{Name: "Alice", Age: 30}

    jsonData, err := json.Marshal(p)
    if err != nil {
        // handle error
    }

    fmt.Println(string(jsonData)) // {"name":"Alice","age":30}
}
```

In this example, the `json.Marshal()` function takes our `Person` struct and returns a byte slice containing the JSON representation.

Notice the `json:"..."` tags on the struct fields. These are called struct tags and are used to control how the struct is encoded. In this case, they specify the names of the JSON keys.

## Decoding JSON (Unmarshalling)

Decoding, or "unmarshalling," is the process of converting a JSON string into a Go data structure. The `json.Unmarshal()` function is used for this.

Let's take a JSON string and unmarshal it into a `Person` struct:

```go
func main() {
    jsonString := `{"name":"Bob","age":25}`
    var p Person

    err := json.Unmarshal([]byte(jsonString), &p)
    if err != nil {
        // handle error
    }

    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
```

The `json.Unmarshal()` function takes a byte slice of JSON data and a pointer to the Go value where you want to store the decoded data.

## Working with Arbitrary JSON Data

Sometimes, you don't know the structure of the JSON data in advance. In these cases, you can unmarshal the JSON into a `map[string]interface{}`.

```go
func main() {
    jsonString := `{"name":"Charlie","age":35,"city":"New York"}`
    var data map[string]interface{}

    err := json.Unmarshal([]byte(jsonString), &data)
    if err != nil {
        // handle error
    }

    name := data["name"].(string)
    age := data["age"].(float64) // Numbers are decoded as float64

    fmt.Printf("Name: %s, Age: %d\n", name, int(age))
}
```

When you unmarshal JSON into an `interface{}`, numeric values are decoded as `float64`. You'll need to use a type assertion to access the values.

## Encoding and Decoding from Streams

The `encoding/json` package also provides `Encoder` and `Decoder` types for working with streams of JSON data, such as reading from or writing to a file or a network connection.

### `json.Encoder`

An `Encoder` writes JSON objects to an output stream.

```go
import "os"

func main() {
    p := Person{Name: "David", Age: 40}
    encoder := json.NewEncoder(os.Stdout)
    encoder.Encode(p)
}
```

### `json.Decoder`

A `Decoder` reads and decodes JSON objects from an input stream.

```go
import "strings"

func main() {
    jsonStream := `{"name":"Eve","age":28}`
    reader := strings.NewReader(jsonStream)
    decoder := json.NewDecoder(reader)

    var p Person
    err := decoder.Decode(&p)
    if err != nil {
        // handle error
    }

    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
```

Working with JSON is a common requirement for modern applications, and Go's `encoding/json` package provides a powerful and flexible set of tools for the job.
