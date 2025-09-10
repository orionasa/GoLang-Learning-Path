# 33. Handlers and Context in Fiber

In Fiber, a handler is a function that receives a `fiber.Ctx` object, processes the request, and sends a response. The `fiber.Ctx` (or Context) object is the heart of a Fiber application. It holds all the information about the incoming HTTP request and provides methods for sending the response.

## The `fiber.Ctx` Object

The `fiber.Ctx` object provides a rich API for working with requests and responses. Let's explore some of its most common uses.

### Handling Incoming Requests

#### Accessing Route Parameters
As we saw in the previous lesson, you can access route parameters using `c.Params()`.

```go
app.Get("/users/:id", func(c *fiber.Ctx) error {
    id := c.Params("id")
    // ...
})
```

#### Accessing Query Parameters
You can get query parameters from the URL using `c.Query()`.

```go
// GET /search?q=fiber
app.Get("/search", func(c *fiber.Ctx) error {
    q := c.Query("q") // "fiber"
    // ...
})
```

#### Accessing the Request Body
For POST, PUT, and PATCH requests, you'll often need to process the request body. The `c.Body()` method returns the raw request body as a byte slice.

For parsing structured data like JSON or form data, Fiber provides the `c.BodyParser()` method.

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

app.Post("/users", func(c *fiber.Ctx) error {
    p := new(Person)

    if err := c.BodyParser(p); err != nil {
        return err
    }

    // ... work with the person object
})
```
`BodyParser()` can automatically decode JSON, form data, and XML based on the `Content-Type` header.

#### Accessing Headers
You can get request headers using `c.Get()`.

```go
authHeader := c.Get("Authorization")
```

### Sending Responses

The `fiber.Ctx` object also provides a variety of methods for sending a response to the client.

#### Sending a String
The `c.SendString()` method sends a plain text response.

```go
return c.SendString("Hello, World!")
```

#### Sending JSON
The `c.JSON()` method serializes a Go struct or map to JSON and sends it as the response. It also sets the `Content-Type` header to `application/json`.

```go
p := Person{Name: "Alice", Age: 30}
return c.JSON(p)
```

#### Setting the Status Code
You can set the HTTP status code using the `c.Status()` method. This method is chainable.

```go
return c.Status(fiber.StatusNotFound).SendString("Not Found")
```

#### Sending Files
The `c.SendFile()` method sends a file as the response.

```go
return c.SendFile("./public/index.html")
```

## The Handler Signature

A Fiber handler is a function that takes a `*fiber.Ctx` and returns an `error`.

```go
func(c *fiber.Ctx) error
```

If a handler returns `nil`, it means that the request was handled successfully. If it returns an error, Fiber's default error handler will catch it and send a `500 Internal Server Error` response. You can also create a custom error handler to control how errors are handled.

The `fiber.Ctx` object is a powerful and flexible tool that provides everything you need to build robust and efficient web applications. Mastering the context object is key to becoming proficient with Fiber.

In the next lesson, we will explore middleware in Fiber.
