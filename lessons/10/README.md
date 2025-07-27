# Module 10: Client-Side Networking

In this module, we'll learn how to make HTTP requests to external services and consume REST APIs in Go.

## The `net/http` Client: Making HTTP Requests (GET, POST)

The `net/http` package provides a powerful and easy-to-use HTTP client.

### Making a GET Request

To make a GET request, you can use the `http.Get` function.

```go
import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("https://api.github.com/users/github")
    if err != nil {
        fmt.Println("Error making GET request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }
    fmt.Println(string(body))
}
```

### Making a POST Request

To make a POST request, you can use the `http.Post` function.

```go
import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    jsonData := []byte(`{"name":"test"}`)
    resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error making POST request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }
    fmt.Println(string(body))
}
```

## Consuming REST APIs from external services

When consuming a REST API, you'll often need to decode the JSON response into a Go struct.

```go
import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Login    string `json:"login"`
    Name     string `json:"name"`
    PublicRepos int `json:"public_repos"`
}

func main() {
    resp, err := http.Get("https://api.github.com/users/github")
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()

    var user User
    err = json.NewDecoder(resp.Body).Decode(&user)
    if err != nil {
        // handle error
    }
    fmt.Printf("%+v\n", user)
}
```

## Handling HTTP client timeouts and errors

It's important to handle timeouts and errors when making HTTP requests. You can create a custom `http.Client` to set a timeout.

```go
import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    resp, err := client.Get("https://api.github.com/users/github")
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    // ...
}
```

## Introduction to low-level TCP with the `net` package

The `net` package provides a lower-level networking interface. You can use it to create TCP clients and servers.

Here's an example of a simple TCP client:

```go
import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        // handle error
    }
    defer conn.Close()

    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    status, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        // handle error
    }
    fmt.Println(status)
}
```

This module has provided an overview of client-side networking in Go. You can now make HTTP requests, consume REST APIs, and work with low-level TCP connections.

**Previous:** [Module 9: Tooling, Dependencies, and Testing](../9/README.md)
**Next:** [Module 11: Web & API Foundations with Fiber](../11/README.md)
