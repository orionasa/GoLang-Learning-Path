# 25. Data Processing in Go

Besides JSON, Go provides support for other common data formats in its standard library. In this lesson, we'll look at how to process CSV and XML data.

## Working with CSV Data

CSV (Comma-Separated Values) is a common format for storing tabular data. The `encoding/csv` package provides tools for reading and writing CSV data.

### Reading a CSV File

The `csv.Reader` type is used to read CSV records from an input stream.

```go
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "strings"
)

func main() {
    csvString := `name,age,city
Alice,30,New York
Bob,25,London`

    reader := csv.NewReader(strings.NewReader(csvString))

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            // handle error
        }
        fmt.Println(record)
    }
}
```
This program will output:
```
[name age city]
[Alice 30 New York]
[Bob 25 London]
```

### Writing a CSV File

The `csv.Writer` type is used to write CSV records to an output stream.

```go
package main

import (
    "encoding/csv"
    "os"
)

func main() {
    records := [][]string{
        {"name", "age", "city"},
        {"Alice", "30", "New York"},
        {"Bob", "25", "London"},
    }

    writer := csv.NewWriter(os.Stdout)

    for _, record := range records {
        if err := writer.Write(record); err != nil {
            // handle error
        }
    }

    // Write any buffered data to the underlying writer.
    writer.Flush()
}
```

## Working with XML Data

XML (eXtensible Markup Language) is another widely used format for data exchange. The `encoding/xml` package provides tools for working with XML data.

### Decoding XML (Unmarshalling)

Similar to JSON, you can unmarshal XML data into a Go struct.

```go
package main

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
    xmlString := `<person><name>Alice</name><age>30</age></person>`
    var p Person

    err := xml.Unmarshal([]byte(xmlString), &p)
    if err != nil {
        // handle error
    }

    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
```

### Encoding XML (Marshalling)

You can also marshal a Go struct into XML data.

```go
func main() {
    p := Person{Name: "Bob", Age: 25}

    xmlData, err := xml.MarshalIndent(p, "", "  ")
    if err != nil {
        // handle error
    }

    fmt.Println(string(xmlData))
}
```
The `xml.MarshalIndent()` function is used to produce nicely formatted XML.

Go's standard library provides robust support for various data formats, making it a versatile language for a wide range of data processing tasks. Whether you're working with JSON, CSV, XML, or other formats, you'll likely find the tools you need in the standard library.
