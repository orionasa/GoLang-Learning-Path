# 20. Networking in Go

Networking is a first-class citizen in Go. The `net/http` package provides a rich set of tools for building HTTP clients and servers. In this lesson, we'll explore the basics of creating a simple web server and making HTTP requests.

## Creating a Simple Web Server

You can create a simple web server in just a few lines of code using the `net/http` package.

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":8080", nil)
}
```

Let's break this down:
*   `http.HandleFunc("/", helloHandler)`: This registers the `helloHandler` function to handle all requests to the web root (`/`).
*   `http.ListenAndServe(":8080", nil)`: This starts an HTTP server on port 8080. The second argument is a handler, but we can leave it as `nil` to use the default server mux (which we've configured with `HandleFunc`).

The `helloHandler` function takes two arguments:
*   `http.ResponseWriter`: This is where you write your HTTP response.
*   `*http.Request`: This contains all the information about the incoming HTTP request, such as the URL, headers, and body.

## Handling Different Routes

You can register handlers for different routes to build a more complex web server.

```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to the home page!")
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is the about page.")
    })

    http.ListenAndServe(":8080", nil)
}
```

## Making HTTP Requests

The `net/http` package also makes it easy to create HTTP clients to make requests to other servers.

### GET Request

Here's how you can make a simple GET request:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    resp, err := http.Get("https://api.github.com/users/golang")
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
}
```

### POST Request

Here's how you can make a POST request with a JSON body:

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    data := map[string]string{"name": "Go"}
    jsonData, _ := json.Marshal(data)

    resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()

    // ... process the response
}
```

## The `http.Client`

For more control over the HTTP client, such as setting timeouts and custom headers, you can create an `http.Client`.

```go
client := &http.Client{
    Timeout: 10 * time.Second,
}

req, err := http.NewRequest("GET", "https://api.github.com", nil)
if err != nil {
    // handle error
}

req.Header.Add("Authorization", "Bearer your-token")

resp, err := client.Do(req)
// ...
```

The `net/http` package provides a solid foundation for building networked applications in Go. It's powerful, flexible, and easy to use.

In the next lesson, we will explore how to work with files on the local file system.
