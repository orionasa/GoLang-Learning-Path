# 22. Making API Calls in Go

Many modern applications rely on data from third-party APIs. Go's `net/http` package makes it straightforward to make API calls and consume their data. In this lesson, we'll walk through how to make GET and POST requests to a JSON API.

## Making a GET Request

A GET request is used to retrieve data from a server. Let's make a GET request to the [JSONPlaceholder](https://jsonplaceholder.typicode.com/) API, which is a fake online REST API for testing and prototyping.

We'll fetch a list of posts from the `/posts` endpoint.

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// Post represents a post from the JSONPlaceholder API
type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func main() {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    var post Post
    err = json.Unmarshal(body, &post)
    if err != nil {
        // handle error
    }

    fmt.Printf("Post Title: %s\n", post.Title)
}
```

In this example:
1.  We define a `Post` struct that matches the structure of the JSON data from the API. The `json:"..."` tags are used to map the JSON keys to the struct fields.
2.  We make a GET request to the API using `http.Get()`.
3.  We read the response body using `io.ReadAll()`.
4.  We use `json.Unmarshal()` to parse the JSON data into our `Post` struct.

## Making a POST Request

A POST request is used to send data to a server to create a new resource. Let's make a POST request to the `/posts` endpoint to create a new post.

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func main() {
    newPost := Post{
        UserID: 1,
        Title:  "My New Post",
        Body:   "This is the body of my new post.",
    }

    jsonData, err := json.Marshal(newPost)
    if err != nil {
        // handle error
    }

    resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println("Response Body:", string(body))
}
```

In this example:
1.  We create a `Post` struct with the data for our new post.
2.  We use `json.Marshal()` to convert our `Post` struct into a JSON byte slice.
3.  We use `http.Post()` to send the JSON data to the API. The second argument is the content type, and the third argument is an `io.Reader` for the request body. We use `bytes.NewBuffer()` to create a reader from our JSON byte slice.
4.  We then print the response status and body to confirm that our post was created.

Making API calls is a fundamental skill for modern software development. Go's standard library provides all the tools you need to interact with web services and build powerful networked applications.

In the next lesson, we will learn about regular expressions for pattern matching in strings.
