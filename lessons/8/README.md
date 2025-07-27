# Module 8: Data Processing and Serialization

In this module, we'll learn how to work with common data formats like JSON, CSV, and XML in Go. We'll also cover how to handle data streams and transformations.

## Working with JSON: Encoding and Decoding

The `encoding/json` package provides support for encoding and decoding JSON data.

### Encoding (Marshalling)

To encode a Go value as JSON, you can use the `json.Marshal` function.

```go
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
        fmt.Println("Error marshalling JSON:", err)
        return
    }
    fmt.Println(string(jsonData)) // prints {"name":"Alice","age":30}
}
```

### Decoding (Unmarshalling)

To decode JSON data into a Go value, you can use the `json.Unmarshal` function.

```go
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    jsonData := []byte(`{"name":"Bob","age":25}`)
    var p Person
    err := json.Unmarshal(jsonData, &p)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }
    fmt.Println(p.Name, p.Age) // prints "Bob 25"
}
```

## Processing other common formats like CSV and XML

### CSV

The `encoding/csv` package provides support for reading and writing CSV files.

```go
import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    // Writing to a CSV file
    file, err := os.Create("test.csv")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    data := [][]string{
        {"Name", "Age"},
        {"Alice", "30"},
        {"Bob", "25"},
    }
    for _, record := range data {
        err := writer.Write(record)
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    }
}
```

### XML

The `encoding/xml` package provides support for parsing XML data.

```go
import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
}

func main() {
    xmlData := []byte(`<person><name>Alice</name><age>30</age></person>`)
    var p Person
    err := xml.Unmarshal(xmlData, &p)
    if err != nil {
        fmt.Println("Error unmarshalling XML:", err)
        return
    }
    fmt.Println(p.Name, p.Age) // prints "Alice 30"
}
```

## Handling data streams and transformations

The `io` package provides interfaces that are useful for working with data streams. The `io.Reader` and `io.Writer` interfaces are particularly important.

You can use these interfaces to create pipelines of data transformations. For example, you could read data from a file, compress it, and then write it to another file.

```go
import (
    "compress/gzip"
    "io"
    "os"
)

func main() {
    // Open the input file
    inputFile, err := os.Open("input.txt")
    if err != nil {
        // handle error
    }
    defer inputFile.Close()

    // Create the output file
    outputFile, err := os.Create("output.txt.gz")
    if err != nil {
        // handle error
    }
    defer outputFile.Close()

    // Create a gzip writer
    gzipWriter := gzip.NewWriter(outputFile)
    defer gzipWriter.Close()

    // Copy the data from the input file to the gzip writer
    _, err = io.Copy(gzipWriter, inputFile)
    if err != nil {
        // handle error
    }
}
```

In this module, you've learned how to work with common data formats and how to handle data streams in Go. This is a crucial skill for building real-world applications.

**Previous:** [Module 7: File System I/O](../7/README.md)
**Next:** [Module 9: Tooling, Dependencies, and Testing](../9/README.md)
