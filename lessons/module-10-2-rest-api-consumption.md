# 🔌 Module 10.2: REST API Consumption

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master Professional REST API Integration**

*From Zero to Go Hero - Module 10.2*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 2-2.5 hours  | Intermediate   | Module 10.1, Understanding of REST principles |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Understand REST API principles and best practices
- ✅ Implement authentication strategies (API keys, Bearer tokens)
- ✅ Handle pagination and large datasets efficiently
- ✅ Build robust error handling and retry mechanisms
- ✅ Create reusable API client structures

---

## 🌐 Understanding REST APIs

<div style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 REST: The Language of Modern Web Services

REST (Representational State Transfer) is like having a **universal conversation protocol** for web services. Just as humans follow social conventions when talking (greetings, turn-taking, politeness), REST APIs follow standardized patterns that make them predictable and easy to use!

</div>

### 🚄 **1. REST API Architecture**
```
┌─────────────────────────────────────────────────────────────┐
│                    REST API ECOSYSTEM                       │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Client App      ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  HTTP Methods    ████████████████████████████████████ 100% │
│  GET POST PUT    ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  REST Endpoints  ████████████████████████████████████ 100% │
│  /users /posts   ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  JSON Response   ████████████████████████████████████ 100% │
│       ↓          ████████████████████████████████████ 100% │
│  Your Go App     ████████████████████████████████████ 100% │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 🧵 **2. REST Principles**
<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**RESTful Design** - Following these principles makes APIs predictable and maintainable:

- **Stateless** - Each request contains all needed information
- **Resource-based** - URLs represent resources (nouns, not verbs)
- **HTTP Methods** - Use appropriate verbs (GET, POST, PUT, DELETE)
- **JSON Format** - Consistent data exchange format
- **Status Codes** - Meaningful HTTP response codes

</div>

### 🎯 **3. Common REST Patterns**
- **Collection**: `/api/users` (GET all users, POST new user)
- **Resource**: `/api/users/123` (GET, PUT, DELETE specific user)
- **Nested**: `/api/users/123/posts` (GET user's posts)
- **Filtering**: `/api/users?role=admin&active=true`
- **Pagination**: `/api/users?page=2&limit=10`

---

## 🛠️ Building Professional API Clients

### **Step 1: Creating a Reusable API Client Structure**

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "time"
)

// APIClient represents a reusable HTTP client for REST APIs
type APIClient struct {
    BaseURL    string
    HTTPClient *http.Client
    Headers    map[string]string
}

// NewAPIClient creates a new API client with default settings
func NewAPIClient(baseURL string) *APIClient {
    return &APIClient{
        BaseURL: baseURL,
        HTTPClient: &http.Client{
            Timeout: 30 * time.Second,
        },
        Headers: make(map[string]string),
    }
}

// SetAPIKey sets the API key for authentication
func (c *APIClient) SetAPIKey(apiKey string) {
    c.Headers["X-API-Key"] = apiKey
}

// SetBearerToken sets the Bearer token for authentication
func (c *APIClient) SetBearerToken(token string) {
    c.Headers["Authorization"] = "Bearer " + token
}

// GET makes a GET request
func (c *APIClient) GET(endpoint string, params map[string]string) (*http.Response, error) {
    u, _ := url.Parse(c.BaseURL + endpoint)
    
    if len(params) > 0 {
        q := u.Query()
        for key, value := range params {
            q.Add(key, value)
        }
        u.RawQuery = q.Encode()
    }
    
    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }
    
    for key, value := range c.Headers {
        req.Header.Set(key, value)
    }
    
    return c.HTTPClient.Do(req)
}

// POST makes a POST request
func (c *APIClient) POST(endpoint string, body interface{}) (*http.Response, error) {
    jsonBody, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    
    req, err := http.NewRequest("POST", c.BaseURL+endpoint, bytes.NewBuffer(jsonBody))
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Content-Type", "application/json")
    for key, value := range c.Headers {
        req.Header.Set(key, value)
    }
    
    return c.HTTPClient.Do(req)
}
```

### **Step 2: Authentication Examples**

```go
// User represents a user from the API
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

func demonstrateAuthentication() {
    // API Key Authentication
    weatherClient := NewAPIClient("https://api.openweathermap.org/data/2.5")
    weatherClient.SetAPIKey("your-api-key-here")
    
    // Bearer Token Authentication  
    githubClient := NewAPIClient("https://api.github.com")
    githubClient.SetBearerToken("your-github-token-here")
    
    // Public API (No Auth)
    publicClient := NewAPIClient("https://jsonplaceholder.typicode.com")
    
    resp, err := publicClient.GET("/users", nil)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    var users []User
    json.NewDecoder(resp.Body).Decode(&users)
    fmt.Printf("Retrieved %d users\n", len(users))
}
```

### **Step 3: Handling Pagination**

```go
// Post represents a blog post
type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func fetchAllPosts(client *APIClient) ([]Post, error) {
    var allPosts []Post
    page := 1
    limit := 10
    
    for {
        params := map[string]string{
            "_page":  fmt.Sprintf("%d", page),
            "_limit": fmt.Sprintf("%d", limit),
        }
        
        resp, err := client.GET("/posts", params)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()
        
        var posts []Post
        if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
            return nil, err
        }
        
        if len(posts) == 0 {
            break // No more posts
        }
        
        allPosts = append(allPosts, posts...)
        
        if len(posts) < limit {
            break // Last page
        }
        
        page++
    }
    
    return allPosts, nil
}
```

---

## 🎯 Practical Project: API Aggregator

<div style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌟 Project: Multi-Source Data Dashboard

Build a dashboard that aggregates data from multiple REST APIs to provide a unified view.

</div>

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "sync"
    "time"
)

// DataSource represents a generic data source
type DataSource struct {
    Name   string
    Client *APIClient
}

// Dashboard aggregates data from multiple sources
type Dashboard struct {
    sources []DataSource
    mu      sync.RWMutex
    data    map[string]interface{}
}

func NewDashboard() *Dashboard {
    return &Dashboard{
        sources: make([]DataSource, 0),
        data:    make(map[string]interface{}),
    }
}

func (d *Dashboard) AddSource(name, baseURL string) {
    client := NewAPIClient(baseURL)
    client.Headers["User-Agent"] = "Dashboard/1.0"
    
    d.sources = append(d.sources, DataSource{
        Name:   name,
        Client: client,
    })
}

func (d *Dashboard) FetchAllData() error {
    var wg sync.WaitGroup
    errors := make(chan error, len(d.sources))
    
    for _, source := range d.sources {
        wg.Add(1)
        go func(s DataSource) {
            defer wg.Done()
            
            resp, err := s.Client.GET("/posts", map[string]string{"_limit": "5"})
            if err != nil {
                errors <- fmt.Errorf("%s: %w", s.Name, err)
                return
            }
            defer resp.Body.Close()
            
            body, err := io.ReadAll(resp.Body)
            if err != nil {
                errors <- fmt.Errorf("%s: %w", s.Name, err)
                return
            }
            
            var posts []Post
            if err := json.Unmarshal(body, &posts); err != nil {
                errors <- fmt.Errorf("%s: %w", s.Name, err)
                return
            }
            
            d.mu.Lock()
            d.data[s.Name] = posts
            d.mu.Unlock()
            
            fmt.Printf("✅ Fetched data from %s\n", s.Name)
        }(source)
    }
    
    wg.Wait()
    close(errors)
    
    // Check for errors
    for err := range errors {
        if err != nil {
            return err
        }
    }
    
    return nil
}

func (d *Dashboard) DisplaySummary() {
    d.mu.RLock()
    defer d.mu.RUnlock()
    
    fmt.Println("\n" + "="*50)
    fmt.Println("📊 DASHBOARD SUMMARY")
    fmt.Println("="*50)
    
    for name, data := range d.data {
        if posts, ok := data.([]Post); ok {
            fmt.Printf("📝 %s: %d posts\n", name, len(posts))
            if len(posts) > 0 {
                fmt.Printf("   Latest: %s\n", posts[0].Title[:50]+"...")
            }
        }
    }
    
    fmt.Println("="*50)
}

func main() {
    dashboard := NewDashboard()
    dashboard.AddSource("JSONPlaceholder", "https://jsonplaceholder.typicode.com")
    
    fmt.Println("🚀 Starting dashboard data fetch...")
    
    start := time.Now()
    if err := dashboard.FetchAllData(); err != nil {
        fmt.Printf("❌ Error: %v\n", err)
        return
    }
    
    fmt.Printf("⏱️ Data fetched in %v\n", time.Since(start))
    dashboard.DisplaySummary()
}
```

---

## 🧪 Practice Exercises

### **Exercise 1: GitHub API Client**
Create a client that:
1. Fetches user repositories
2. Handles authentication with personal access token
3. Implements pagination for large repository lists

### **Exercise 2: Weather Service Aggregator**
Build a service that:
1. Fetches weather from multiple APIs
2. Compares results and shows differences
3. Handles API failures gracefully

### **Exercise 3: E-commerce API Consumer**
Create a client that:
1. Fetches product listings with filtering
2. Handles different response formats
3. Implements caching for frequently accessed data

---

## 🎯 Key Takeaways

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**🎓 What You've Learned:**

1. **Professional API Clients** - Reusable, configurable client structures
2. **Authentication** - API keys, Bearer tokens, and security best practices
3. **Pagination** - Handling large datasets efficiently
4. **Error Handling** - Robust error management and retry logic
5. **Concurrent Requests** - Using goroutines for parallel API calls

</div>

---

## 🚀 What's Next?

Excellent work! You've mastered REST API consumption. Next, we'll explore **Timeouts & Error Handling** to make your network code production-ready.

**Ready to continue?** → [Module 10.3: Timeouts & Error Handling](./module-10-3-timeouts-error-handling.md)

---

<div align="center">

**Happy API Consuming! 🔌**

*Remember: Great APIs are consumed, not just called!*

</div>
