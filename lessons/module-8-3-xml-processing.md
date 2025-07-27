# 🚀 Module 8.3: XML Processing

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master XML Parsing and Generation in Go**

*From Zero to Go Hero - Module 8.3*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1-1.5 hours  | Intermediate   | Structs, Interfaces, JSON Processing |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Parse XML documents with Go's encoding/xml package
- ✅ Use XML struct tags for custom field mapping
- ✅ Handle XML attributes, namespaces, and nested elements
- ✅ Generate XML documents from Go structs
- ✅ Build configuration parsers and API clients

---

## 🌟 What is XML?

<div style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🏷️ **XML (eXtensible Markup Language)**

XML is a markup language that defines rules for encoding documents in a format that is both human-readable and machine-readable. It's widely used for configuration files, web services, and data exchange.

**Think of XML like a structured document** - it's like HTML but for data, with custom tags that describe the content!

</div>

### 🏗️ **XML Structure Example**
```
┌─────────────────────────────────────┐
│            XML Structure            │
├─────────────────────────────────────┤
│  <?xml version="1.0"?>              │
│  <root>                             │
│    <person id="1">                  │
│      <name>John Doe</name>          │
│      <age>25</age>                  │
│    </person>                        │
│  </root>                            │
└─────────────────────────────────────┘
```

---

## 🔧 XML in Go: The `encoding/xml` Package

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Key Functions** - Go's XML package provides:

```go
// Parse XML into Go structs (Unmarshal)
func Unmarshal(data []byte, v interface{}) error

// Convert Go structs to XML (Marshal)
func Marshal(v interface{}) ([]byte, error)

// Pretty-print XML with indentation
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```

</div>

### 🎯 **Basic XML Operations**

```go
package main

import (
    "encoding/xml"
    "fmt"
    "log"
)

// Book represents a book with XML tags
type Book struct {
    XMLName     xml.Name `xml:"book"`
    ID          int      `xml:"id,attr"`
    Title       string   `xml:"title"`
    Author      string   `xml:"author"`
    Genre       string   `xml:"genre"`
    Price       float64  `xml:"price"`
    PublishDate string   `xml:"publish_date"`
    Description string   `xml:"description"`
}

// Library represents a collection of books
type Library struct {
    XMLName xml.Name `xml:"library"`
    Name    string   `xml:"name,attr"`
    Books   []Book   `xml:"book"`
}

func main() {
    // Sample XML data
    xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<library name="City Library">
    <book id="1">
        <title>The Go Programming Language</title>
        <author>Alan Donovan</author>
        <genre>Programming</genre>
        <price>45.99</price>
        <publish_date>2015-11-16</publish_date>
        <description>A comprehensive guide to Go programming</description>
    </book>
    <book id="2">
        <title>Clean Code</title>
        <author>Robert Martin</author>
        <genre>Programming</genre>
        <price>42.99</price>
        <publish_date>2008-08-01</publish_date>
        <description>A handbook of agile software craftsmanship</description>
    </book>
</library>`

    // Parse XML into Go struct
    var library Library
    err := xml.Unmarshal([]byte(xmlData), &library)
    if err != nil {
        log.Fatal("Error parsing XML:", err)
    }

    // Display parsed data
    fmt.Println("📚 LIBRARY INFORMATION")
    fmt.Println("=====================")
    fmt.Printf("Library Name: %s\n", library.Name)
    fmt.Printf("Number of Books: %d\n\n", len(library.Books))

    for i, book := range library.Books {
        fmt.Printf("📖 Book %d:\n", i+1)
        fmt.Printf("  ID: %d\n", book.ID)
        fmt.Printf("  Title: %s\n", book.Title)
        fmt.Printf("  Author: %s\n", book.Author)
        fmt.Printf("  Price: $%.2f\n", book.Price)
        fmt.Println()
    }

    // Marshal back to XML
    output, err := xml.MarshalIndent(library, "", "  ")
    if err != nil {
        log.Fatal("Error marshaling XML:", err)
    }

    fmt.Println("🔄 GENERATED XML:")
    fmt.Println("================")
    fmt.Println(xml.Header + string(output))
}
```

---

## 🏷️ XML Struct Tags: Advanced Mapping

<div style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎨 **XML Struct Tags Power**

XML struct tags give you fine-grained control over how Go structs map to XML elements and attributes. They're essential for working with existing XML schemas.

</div>

### 🔧 **XML Tag Options**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🏷️ Element Name</strong><br>
`xml:"elementname"`
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📝 Attribute</strong><br>
`xml:"name,attr"`
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📄 Character Data</strong><br>
`xml:",chardata"`
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🚫 Omit Empty</strong><br>
`xml:"name,omitempty"`
</div>

</div>

---

## 🎯 Practice Exercises

### 🏋️ **Exercise 1: Configuration Parser**
Create a system that:
1. Parses XML configuration files
2. Validates required settings
3. Generates default configurations
4. Handles nested configuration sections

### 🏋️ **Exercise 2: RSS Feed Reader**
Build an RSS reader that:
1. Fetches RSS feeds from URLs
2. Parses XML feed data
3. Displays articles in a formatted way
4. Supports different RSS versions

### 🏋️ **Exercise 3: SOAP Client**
Create a SOAP web service client that:
1. Generates SOAP request envelopes
2. Handles XML namespaces properly
3. Parses SOAP responses
4. Implements error handling

---

## 🎉 Summary

Congratulations! You've mastered XML processing in Go. Here's what you've learned:

<div style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎯 **Key Takeaways**
- ✅ **XML Parsing** - Parse XML documents into Go structs
- ✅ **Struct Tags** - Control XML mapping with struct tags
- ✅ **Attributes & Elements** - Handle XML attributes and nested elements
- ✅ **Namespaces** - Work with XML namespaces in web services
- ✅ **Generation** - Create XML documents from Go data

</div>

---

## 🚀 What's Next?

Ready to explore data streams and transformations? Continue with:

<div style="text-align: center; margin: 30px 0;">
<a href="./module-8-4-data-streams-transformations.md" style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%); color: white; padding: 15px 30px; border-radius: 25px; text-decoration: none; font-weight: bold; font-size: 18px;">🌊 Next: Data Streams →</a>
</div>

---

*Happy coding! 🎉*
