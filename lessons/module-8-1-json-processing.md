# 🚀 Module 8.1: JSON Processing

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master JSON Encoding and Decoding in Go**

*From Zero to Go Hero - Module 8.1*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1-1.5 hours  | Intermediate   | Structs, Interfaces, Error Handling |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand JSON marshaling and unmarshaling in Go
- ✅ Use struct tags to customize JSON field names
- ✅ Handle nested JSON structures and arrays
- ✅ Implement custom JSON marshaling for complex types
- ✅ Build a real-world API client with JSON processing

---

## 🌟 What is JSON?

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 📊 **JSON (JavaScript Object Notation)**

JSON is a lightweight, text-based data interchange format that's easy for humans to read and write, and easy for machines to parse and generate. It's the standard format for web APIs and configuration files.

**Think of JSON like a universal translator** - it helps different systems speak the same language!

</div>

### 🏗️ **JSON Structure Basics**
```
┌─────────────────────────────────────┐
│         JSON Data Types             │
├─────────────────────────────────────┤
│  String    "Hello World"        ✓   │
│  Number    42, 3.14             ✓   │
│  Boolean   true, false          ✓   │
│  null      null                 ✓   │
│  Object    {"key": "value"}     ✓   │
│  Array     [1, 2, 3]            ✓   │
└─────────────────────────────────────┘
```

---

## 🔧 JSON in Go: The `encoding/json` Package

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Key Functions** - Go provides two main functions for JSON processing:

```go
// Convert Go data to JSON (Marshal)
func Marshal(v interface{}) ([]byte, error)

// Convert JSON to Go data (Unmarshal)  
func Unmarshal(data []byte, v interface{}) error
```

</div>

### 🎯 **Basic JSON Operations**

Let's start with simple examples:

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
)

// Person represents a person with basic information
type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Email   string `json:"email"`
    IsAdmin bool   `json:"is_admin"`
}

func main() {
    // Creating a Person
    person := Person{
        Name:    "Alice Johnson",
        Age:     28,
        Email:   "alice@example.com",
        IsAdmin: false,
    }

    // Marshal (Go struct → JSON)
    jsonData, err := json.Marshal(person)
    if err != nil {
        log.Fatal("Error marshaling JSON:", err)
    }
    
    fmt.Println("JSON Output:")
    fmt.Println(string(jsonData))
    
    // Unmarshal (JSON → Go struct)
    var newPerson Person
    err = json.Unmarshal(jsonData, &newPerson)
    if err != nil {
        log.Fatal("Error unmarshaling JSON:", err)
    }
    
    fmt.Printf("Parsed Person: %+v\n", newPerson)
}
```

**Output:**
```json
JSON Output:
{"name":"Alice Johnson","age":28,"email":"alice@example.com","is_admin":false}
Parsed Person: {Name:Alice Johnson Age:28 Email:alice@example.com IsAdmin:false}
```

---

## 🏷️ Struct Tags: Customizing JSON Output

<div style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎨 **Struct Tags Power**

Struct tags let you control how Go structs are converted to/from JSON. They're like instructions that tell Go exactly how to handle each field.

</div>

### 🔧 **Common Struct Tag Options**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🏷️ Field Naming</strong><br>
`json:"custom_name"`
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🚫 Omit Empty</strong><br>
`json:"field,omitempty"`
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>👻 Ignore Field</strong><br>
`json:"-"`
</div>

<div style="background: #ffffff; border: 2px solid #e9ecef; border-radius: 8px; padding: 15px; text-align: center;">
<strong>📝 String Numbers</strong><br>
`json:"id,string"`
</div>

</div>

### 🎯 **Advanced Struct Tags Example**

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
)

// User represents a user with various JSON customizations
type User struct {
    ID          int       `json:"id"`
    Username    string    `json:"username"`
    Email       string    `json:"email"`
    Password    string    `json:"-"`                    // Never include in JSON
    FirstName   string    `json:"first_name,omitempty"` // Omit if empty
    LastName    string    `json:"last_name,omitempty"`  // Omit if empty
    Age         int       `json:"age,omitempty"`        // Omit if zero
    CreatedAt   time.Time `json:"created_at"`
    IsActive    bool      `json:"is_active"`
    ProfilePic  string    `json:"profile_picture,omitempty"`
}

func main() {
    user := User{
        ID:        123,
        Username:  "johndoe",
        Email:     "john@example.com",
        Password:  "secret123", // This won't appear in JSON
        FirstName: "John",
        LastName:  "", // This will be omitted (empty)
        Age:       0,  // This will be omitted (zero value)
        CreatedAt: time.Now(),
        IsActive:  true,
    }

    jsonData, _ := json.MarshalIndent(user, "", "  ")
    fmt.Println("User JSON:")
    fmt.Println(string(jsonData))
}
```

**Output:**
```json
User JSON:
{
  "id": 123,
  "username": "johndoe",
  "email": "john@example.com",
  "first_name": "John",
  "created_at": "2024-01-15T10:30:00Z",
  "is_active": true
}
```

---

## 🌊 Working with Complex JSON Structures

### 🏢 **Nested Structures**

```go
package main

import (
    "encoding/json"
    "fmt"
)

// Address represents a physical address
type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    State   string `json:"state"`
    ZipCode string `json:"zip_code"`
    Country string `json:"country"`
}

// Company represents a company with nested address
type Company struct {
    Name        string    `json:"name"`
    Industry    string    `json:"industry"`
    Address     Address   `json:"address"`
    Employees   []string  `json:"employees"`
    Founded     int       `json:"founded"`
    IsPublic    bool      `json:"is_public"`
}

func main() {
    company := Company{
        Name:     "TechCorp Inc.",
        Industry: "Software Development",
        Address: Address{
            Street:  "123 Innovation Drive",
            City:    "San Francisco",
            State:   "CA",
            ZipCode: "94105",
            Country: "USA",
        },
        Employees: []string{"Alice", "Bob", "Charlie", "Diana"},
        Founded:   2015,
        IsPublic:  true,
    }

    // Pretty print JSON
    jsonData, _ := json.MarshalIndent(company, "", "  ")
    fmt.Println("Company JSON:")
    fmt.Println(string(jsonData))

    // Parse it back
    var parsedCompany Company
    json.Unmarshal(jsonData, &parsedCompany)
    
    fmt.Printf("\nParsed Company Name: %s\n", parsedCompany.Name)
    fmt.Printf("Located in: %s, %s\n", parsedCompany.Address.City, parsedCompany.Address.State)
    fmt.Printf("Number of employees: %d\n", len(parsedCompany.Employees))
}
```

---

## 🎯 Real-World Example: Building a Weather API Client

Let's build a practical example that fetches and processes weather data:

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

// WeatherResponse represents the API response structure
type WeatherResponse struct {
    Location Location `json:"location"`
    Current  Current  `json:"current"`
    Forecast Forecast `json:"forecast"`
}

type Location struct {
    Name      string  `json:"name"`
    Region    string  `json:"region"`
    Country   string  `json:"country"`
    Latitude  float64 `json:"lat"`
    Longitude float64 `json:"lon"`
    Timezone  string  `json:"tz_id"`
}

type Current struct {
    Temperature   float64   `json:"temp_c"`
    Condition     Condition `json:"condition"`
    WindSpeed     float64   `json:"wind_kph"`
    Humidity      int       `json:"humidity"`
    FeelsLike     float64   `json:"feelslike_c"`
    LastUpdated   string    `json:"last_updated"`
}

type Condition struct {
    Text string `json:"text"`
    Icon string `json:"icon"`
    Code int    `json:"code"`
}

type Forecast struct {
    ForecastDays []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
    Date string `json:"date"`
    Day  Day    `json:"day"`
}

type Day struct {
    MaxTemp   float64   `json:"maxtemp_c"`
    MinTemp   float64   `json:"mintemp_c"`
    AvgTemp   float64   `json:"avgtemp_c"`
    Condition Condition `json:"condition"`
    Humidity  int       `json:"avghumidity"`
}

// WeatherClient handles weather API operations
type WeatherClient struct {
    baseURL string
    apiKey  string
}

// NewWeatherClient creates a new weather client
func NewWeatherClient(apiKey string) *WeatherClient {
    return &WeatherClient{
        baseURL: "http://api.weatherapi.com/v1",
        apiKey:  apiKey,
    }
}

// GetWeather fetches weather data for a location
func (wc *WeatherClient) GetWeather(location string) (*WeatherResponse, error) {
    url := fmt.Sprintf("%s/forecast.json?key=%s&q=%s&days=3", 
        wc.baseURL, wc.apiKey, location)
    
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to make request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response: %w", err)
    }

    var weather WeatherResponse
    if err := json.Unmarshal(body, &weather); err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %w", err)
    }

    return &weather, nil
}

// DisplayWeather prints weather information in a nice format
func DisplayWeather(weather *WeatherResponse) {
    fmt.Println("🌤️  WEATHER REPORT")
    fmt.Println("==================")
    
    loc := weather.Location
    fmt.Printf("📍 Location: %s, %s, %s\n", loc.Name, loc.Region, loc.Country)
    fmt.Printf("🌐 Coordinates: %.2f, %.2f\n", loc.Latitude, loc.Longitude)
    fmt.Println()

    current := weather.Current
    fmt.Println("🌡️  CURRENT CONDITIONS")
    fmt.Printf("Temperature: %.1f°C (feels like %.1f°C)\n", 
        current.Temperature, current.FeelsLike)
    fmt.Printf("Condition: %s\n", current.Condition.Text)
    fmt.Printf("Wind Speed: %.1f km/h\n", current.WindSpeed)
    fmt.Printf("Humidity: %d%%\n", current.Humidity)
    fmt.Printf("Last Updated: %s\n", current.LastUpdated)
    fmt.Println()

    fmt.Println("📅 3-DAY FORECAST")
    for i, day := range weather.Forecast.ForecastDays {
        fmt.Printf("Day %d (%s):\n", i+1, day.Date)
        fmt.Printf("  🌡️  High: %.1f°C, Low: %.1f°C\n", 
            day.Day.MaxTemp, day.Day.MinTemp)
        fmt.Printf("  ☁️  Condition: %s\n", day.Day.Condition.Text)
        fmt.Printf("  💧 Humidity: %d%%\n", day.Day.Humidity)
        fmt.Println()
    }
}

func main() {
    // Note: You'll need a real API key from weatherapi.com
    client := NewWeatherClient("your-api-key-here")
    
    // For demo purposes, let's work with mock JSON data
    mockJSON := `{
        "location": {
            "name": "San Francisco",
            "region": "California",
            "country": "USA",
            "lat": 37.77,
            "lon": -122.42,
            "tz_id": "America/Los_Angeles"
        },
        "current": {
            "temp_c": 18.5,
            "condition": {
                "text": "Partly cloudy",
                "icon": "//cdn.weatherapi.com/weather/64x64/day/116.png",
                "code": 1003
            },
            "wind_kph": 15.8,
            "humidity": 65,
            "feelslike_c": 19.2,
            "last_updated": "2024-01-15 14:30"
        },
        "forecast": {
            "forecastday": [
                {
                    "date": "2024-01-15",
                    "day": {
                        "maxtemp_c": 22.0,
                        "mintemp_c": 12.0,
                        "avgtemp_c": 17.0,
                        "condition": {
                            "text": "Partly cloudy",
                            "code": 1003
                        },
                        "avghumidity": 68
                    }
                },
                {
                    "date": "2024-01-16",
                    "day": {
                        "maxtemp_c": 20.5,
                        "mintemp_c": 11.5,
                        "avgtemp_c": 16.0,
                        "condition": {
                            "text": "Sunny",
                            "code": 1000
                        },
                        "avghumidity": 62
                    }
                }
            ]
        }
    }`

    var weather WeatherResponse
    if err := json.Unmarshal([]byte(mockJSON), &weather); err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
        return
    }

    DisplayWeather(&weather)
}
```

---

## 🛠️ Handling JSON Errors and Edge Cases

### 🚨 **Common JSON Pitfalls and Solutions**

<div style="background: #fff3cd; border: 1px solid #ffeaa7; border-radius: 8px; padding: 15px; margin: 15px 0;">

**⚠️ Warning: Common Mistakes**
- Forgetting to use pointers when unmarshaling
- Not handling missing or null fields properly
- Ignoring JSON parsing errors
- Using wrong struct tag syntax

</div>

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
)

// Product with optional fields and error handling
type Product struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    Price       *float64 `json:"price,omitempty"`        // Pointer for optional field
    Description *string  `json:"description,omitempty"`  // Pointer for optional field
    Tags        []string `json:"tags"`
    InStock     bool     `json:"in_stock"`
}

func main() {
    // JSON with missing optional fields
    jsonData := `{
        "id": 1,
        "name": "Laptop",
        "tags": ["electronics", "computers"],
        "in_stock": true
    }`

    var product Product
    if err := json.Unmarshal([]byte(jsonData), &product); err != nil {
        log.Fatal("Error unmarshaling:", err)
    }

    fmt.Printf("Product: %+v\n", product)
    
    // Check optional fields safely
    if product.Price != nil {
        fmt.Printf("Price: $%.2f\n", *product.Price)
    } else {
        fmt.Println("Price: Not specified")
    }

    if product.Description != nil {
        fmt.Printf("Description: %s\n", *product.Description)
    } else {
        fmt.Println("Description: Not provided")
    }

    // Marshal back to JSON
    output, err := json.MarshalIndent(product, "", "  ")
    if err != nil {
        log.Fatal("Error marshaling:", err)
    }
    
    fmt.Println("\nJSON Output:")
    fmt.Println(string(output))
}
```

---

## 🎯 Practice Exercises

### 🏋️ **Exercise 1: User Profile API**
Create a user profile system that can:
1. Parse user data from JSON
2. Handle optional fields (profile picture, bio, social links)
3. Generate JSON output with proper formatting

### 🏋️ **Exercise 2: Configuration File Parser**
Build a configuration file parser that:
1. Reads JSON configuration files
2. Validates required fields
3. Provides default values for missing options
4. Handles nested configuration sections

### 🏋️ **Exercise 3: API Response Handler**
Create an API client that:
1. Makes HTTP requests to a REST API
2. Parses JSON responses with error handling
3. Handles different response formats (success/error)
4. Implements retry logic for failed requests

---

## 🎉 Summary

Congratulations! You've mastered JSON processing in Go. Here's what you've learned:

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎯 **Key Takeaways**
- ✅ **JSON Marshaling/Unmarshaling** - Convert between Go structs and JSON
- ✅ **Struct Tags** - Customize JSON field names and behavior
- ✅ **Complex Structures** - Handle nested objects and arrays
- ✅ **Error Handling** - Properly handle JSON parsing errors
- ✅ **Real-world Applications** - Build API clients and data processors

</div>

---

## 🚀 What's Next?

Ready to explore more data formats? Continue with:

<div style="text-align: center; margin: 30px 0;">
<a href="./module-8-2-csv-processing.md" style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); color: white; padding: 15px 30px; border-radius: 25px; text-decoration: none; font-weight: bold; font-size: 18px;">📊 Next: CSV Processing →</a>
</div>

---

*Happy coding! 🎉*
