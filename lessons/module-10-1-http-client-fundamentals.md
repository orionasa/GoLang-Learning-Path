# 🌍 Module 10.1: HTTP Client Fundamentals

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master HTTP Requests with Go's Built-in Client**

*From Zero to Go Hero - Module 10.1*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1.5-2 hours  | Beginner       | Basic Go syntax, understanding of HTTP concepts |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand the `net/http` package and its HTTP client capabilities
- ✅ Make GET and POST requests with proper headers
- ✅ Handle HTTP responses and status codes effectively
- ✅ Work with request/response bodies and JSON data
- ✅ Build a practical HTTP client for real-world scenarios

---

## 🌐 Understanding HTTP in Go

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 The net/http Package: Your Gateway to the Web

Go's `net/http` package is like having a **universal translator** for web communication. Just as a translator helps you communicate with people who speak different languages, the HTTP client helps your Go program communicate with web servers around the world!

</div>

### 🚄 **1. HTTP Client Architecture**
```
┌─────────────────────────────────────────────────────────────┐
│                    HTTP CLIENT WORKFLOW                     │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Your Go App     ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  net/http        ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  HTTP Request    ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  Web Server      ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  HTTP Response   ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  Your Go App     ████████████████████████████████████ 100% │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 🧵 **2. The Default HTTP Client**
<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Default Client** - Go provides a ready-to-use HTTP client that handles most common scenarios

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
)

func main() {
    // The simplest HTTP GET request
    response, err := http.Get("https://httpbin.org/get")
    if err != nil {
        log.Fatal("Error making request:", err)
    }
    defer response.Body.Close()

    // Read the response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal("Error reading response:", err)
    }

    fmt.Printf("Status: %s\n", response.Status)
    fmt.Printf("Body: %s\n", string(body))
}
```

</div>

### 🎯 **3. HTTP Methods Explained**
- **GET** - Retrieve data from a server (like asking for information)
- **POST** - Send data to a server (like submitting a form)
- **PUT** - Update existing data on a server
- **DELETE** - Remove data from a server
- **PATCH** - Partially update data on a server

### 🏢 **4. HTTP Status Code Categories**

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin: 20px 0;">

<div style="background: #d4edda; border: 2px solid #c3e6cb; border-radius: 8px; padding: 15px; text-align: center;">
<strong>✅ 2xx Success</strong><br>
Request was successful
</div>

<div style="background: #fff3cd; border: 2px solid #ffeaa7; border-radius: 8px; padding: 15px; text-align: center;">
<strong>🔄 3xx Redirection</strong><br>
Further action needed
</div>

<div style="background: #f8d7da; border: 2px solid #f5c6cb; border-radius: 8px; padding: 15px; text-align: center;">
<strong>❌ 4xx Client Error</strong><br>
Problem with request
</div>

<div style="background: #f8d7da; border: 2px solid #f5c6cb; border-radius: 8px; padding: 15px; text-align: center;">
<strong>💥 5xx Server Error</strong><br>
Problem with server
</div>

</div>

---

## 🛠️ Making HTTP Requests: Step-by-Step Guide

### **Step 1: Simple GET Request**

Let's start with the most basic HTTP operation - fetching data from a server:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
)

func simpleGetRequest() {
    fmt.Println("🌍 Making a simple GET request...")
    
    // Make the request
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        log.Fatal("❌ Request failed:", err)
    }
    defer resp.Body.Close() // Always close the response body!

    // Check status code
    fmt.Printf("📊 Status Code: %d\n", resp.StatusCode)
    fmt.Printf("📋 Status: %s\n", resp.Status)

    // Read response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("❌ Failed to read response:", err)
    }

    fmt.Printf("📄 Response Body:\n%s\n", string(body))
}

func main() {
    simpleGetRequest()
}
```

### **Step 2: GET Request with Custom Headers**

Real-world applications often need to send custom headers:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
)

func getWithHeaders() {
    fmt.Println("🔧 Making GET request with custom headers...")
    
    // Create a new request
    req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
    if err != nil {
        log.Fatal("❌ Failed to create request:", err)
    }

    // Add custom headers
    req.Header.Add("User-Agent", "GoLang-Tutorial-Client/1.0")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", "Bearer your-token-here")

    // Create HTTP client and make request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("❌ Request failed:", err)
    }
    defer resp.Body.Close()

    // Read and display response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("❌ Failed to read response:", err)
    }

    fmt.Printf("📊 Status: %s\n", resp.Status)
    fmt.Printf("📄 Response:\n%s\n", string(body))
}

func main() {
    getWithHeaders()
}
```

### **Step 3: POST Request with JSON Data**

Sending data to a server is equally important:

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "log"
)

// Define a struct for our data
type Post struct {
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID int    `json:"userId"`
}

func postWithJSON() {
    fmt.Println("📤 Making POST request with JSON data...")
    
    // Create data to send
    newPost := Post{
        Title:  "My First Go HTTP Post",
        Body:   "This is the body of my post created with Go!",
        UserID: 1,
    }

    // Convert struct to JSON
    jsonData, err := json.Marshal(newPost)
    if err != nil {
        log.Fatal("❌ Failed to marshal JSON:", err)
    }

    // Create POST request
    resp, err := http.Post(
        "https://jsonplaceholder.typicode.com/posts",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        log.Fatal("❌ POST request failed:", err)
    }
    defer resp.Body.Close()

    // Read response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("❌ Failed to read response:", err)
    }

    fmt.Printf("📊 Status: %s\n", resp.Status)
    fmt.Printf("📄 Response:\n%s\n", string(body))
}

func main() {
    postWithJSON()
}
```

### **Step 4: Handling Response Headers and Cookies**

Sometimes you need to inspect response metadata:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
)

func inspectResponse() {
    fmt.Println("🔍 Inspecting HTTP response details...")
    
    resp, err := http.Get("https://httpbin.org/response-headers?Content-Type=application/json&Server=httpbin/1.0")
    if err != nil {
        log.Fatal("❌ Request failed:", err)
    }
    defer resp.Body.Close()

    // Display status information
    fmt.Printf("📊 Status Code: %d\n", resp.StatusCode)
    fmt.Printf("📋 Status: %s\n", resp.Status)
    fmt.Printf("🌐 Protocol: %s\n", resp.Proto)

    // Display headers
    fmt.Println("\n📋 Response Headers:")
    for name, values := range resp.Header {
        for _, value := range values {
            fmt.Printf("  %s: %s\n", name, value)
        }
    }

    // Display cookies (if any)
    if len(resp.Cookies()) > 0 {
        fmt.Println("\n🍪 Cookies:")
        for _, cookie := range resp.Cookies() {
            fmt.Printf("  %s = %s\n", cookie.Name, cookie.Value)
        }
    }

    // Read body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("❌ Failed to read response:", err)
    }

    fmt.Printf("\n📄 Response Body:\n%s\n", string(body))
}

func main() {
    inspectResponse()
}
```

---

## 🎯 Practical Project: Weather API Client

Let's build a practical HTTP client that fetches weather data:

<div style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌤️ Project: Simple Weather Client

We'll create a client that fetches weather data from a public API and displays it in a user-friendly format.

</div>

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "log"
    "time"
)

// Weather API response structure
type WeatherResponse struct {
    Location struct {
        Name    string `json:"name"`
        Country string `json:"country"`
    } `json:"location"`
    Current struct {
        TempC     float64 `json:"temp_c"`
        TempF     float64 `json:"temp_f"`
        Condition struct {
            Text string `json:"text"`
            Icon string `json:"icon"`
        } `json:"condition"`
        Humidity int     `json:"humidity"`
        WindKph  float64 `json:"wind_kph"`
    } `json:"current"`
}

type WeatherClient struct {
    baseURL string
    apiKey  string
    client  *http.Client
}

func NewWeatherClient(apiKey string) *WeatherClient {
    return &WeatherClient{
        baseURL: "http://api.weatherapi.com/v1",
        apiKey:  apiKey,
        client: &http.Client{
            Timeout: 10 * time.Second, // 10 second timeout
        },
    }
}

func (wc *WeatherClient) GetCurrentWeather(city string) (*WeatherResponse, error) {
    // Build the URL
    url := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no", 
        wc.baseURL, wc.apiKey, city)

    fmt.Printf("🌍 Fetching weather for %s...\n", city)

    // Make the request
    resp, err := wc.client.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to make request: %w", err)
    }
    defer resp.Body.Close()

    // Check status code
    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("API returned status %d: %s", 
            resp.StatusCode, string(body))
    }

    // Read and parse response
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

func displayWeather(weather *WeatherResponse) {
    fmt.Println("\n" + "="*50)
    fmt.Printf("🌤️  WEATHER REPORT FOR %s, %s\n", 
        weather.Location.Name, weather.Location.Country)
    fmt.Println("="*50)
    
    fmt.Printf("🌡️  Temperature: %.1f°C (%.1f°F)\n", 
        weather.Current.TempC, weather.Current.TempF)
    fmt.Printf("☁️  Condition: %s\n", weather.Current.Condition.Text)
    fmt.Printf("💧 Humidity: %d%%\n", weather.Current.Humidity)
    fmt.Printf("💨 Wind Speed: %.1f km/h\n", weather.Current.WindKph)
    
    fmt.Println("="*50)
}

func main() {
    // Note: You'll need to get a free API key from weatherapi.com
    // For demo purposes, we'll use a placeholder
    apiKey := "your-api-key-here"
    
    if apiKey == "your-api-key-here" {
        fmt.Println("⚠️  Please get a free API key from https://www.weatherapi.com/")
        fmt.Println("📝 Replace 'your-api-key-here' with your actual API key")
        return
    }

    // Create weather client
    client := NewWeatherClient(apiKey)

    // Cities to check
    cities := []string{"London", "New York", "Tokyo", "Sydney"}

    for _, city := range cities {
        weather, err := client.GetCurrentWeather(city)
        if err != nil {
            fmt.Printf("❌ Error fetching weather for %s: %v\n", city, err)
            continue
        }

        displayWeather(weather)
        time.Sleep(1 * time.Second) // Be nice to the API
    }
}
```

---

## 🧪 Practice Exercises

### **Exercise 1: Basic HTTP Client**
Create a program that:
1. Makes a GET request to `https://httpbin.org/json`
2. Prints the status code and response body
3. Handles any errors gracefully

### **Exercise 2: Custom Headers**
Build a client that:
1. Sends a GET request with custom `User-Agent` and `Accept` headers
2. Makes a request to `https://httpbin.org/headers`
3. Displays the headers that the server received

### **Exercise 3: POST with Form Data**
Create a program that:
1. Sends POST data as form-encoded values
2. Uses `https://httpbin.org/post` as the endpoint
3. Sends at least 3 key-value pairs

### **Exercise 4: JSON API Consumer**
Build a client for the JSONPlaceholder API that:
1. Fetches all posts from `https://jsonplaceholder.typicode.com/posts`
2. Displays the first 5 posts in a formatted way
3. Creates a new post and displays the server response

---

## 🎯 Key Takeaways

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**🎓 What You've Learned:**

1. **HTTP Client Basics** - Using `net/http` for web requests
2. **Request Methods** - GET, POST, and custom headers
3. **Response Handling** - Status codes, headers, and body parsing
4. **JSON Processing** - Marshaling and unmarshaling data
5. **Error Handling** - Graceful failure management
6. **Real-world Application** - Building a weather API client

</div>

### 🔄 **Common Patterns to Remember:**
- Always close response bodies with `defer resp.Body.Close()`
- Check status codes before processing responses
- Use `json.Marshal` and `json.Unmarshal` for JSON data
- Handle errors at every step of the process
- Set appropriate timeouts for production applications

---

## 🚀 What's Next?

Great job! You've mastered the fundamentals of HTTP clients in Go. In the next lesson, we'll dive deeper into **REST API Consumption**, where you'll learn:

- 🔐 Authentication strategies (API keys, OAuth)
- 📄 Pagination handling for large datasets
- 🔄 Retry mechanisms and circuit breakers
- 📊 Response caching and optimization

**Ready to level up?** → [Continue to Module 10.2: REST API Consumption](./module-10-2-rest-api-consumption.md)

---

<div align="center">

**Happy HTTP Requesting! 🌐**

*Remember: The web is your oyster - Go makes it easy to crack it open!*

</div>
