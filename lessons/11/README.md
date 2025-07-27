# Module 11: Web & API Foundations with Fiber

In this module, we'll introduce Fiber, a popular web framework for Go, and learn how to build a basic web API.

## Introduction to Web Frameworks in Go and why Fiber

While Go's standard `net/http` package is powerful, web frameworks can provide a more productive and structured way to build web applications. Frameworks often include features like routing, middleware, and data binding.

Fiber is an Express.js-inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to be fast, flexible, and easy to use.

## Setting up a new Fiber project: Routing, Handlers, and Context

To get started with Fiber, you'll need to install it:

```bash
go get -u github.com/gofiber/fiber/v2
```

Here's a simple "Hello, World" example with Fiber:

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":3000")
}
```

*   **`fiber.New()`**: Creates a new Fiber app.
*   **`app.Get("/", ...)`**: Defines a route for GET requests to the root path (`/`).
*   **`func(c *fiber.Ctx) error`**: This is a handler function. It takes a `fiber.Ctx` (context) object, which provides access to the request and response.
*   **`app.Listen(":3000")`**: Starts the server on port 3000.

## Handling Requests and Responses (JSON, query params, etc.)

Fiber makes it easy to work with different types of requests and responses.

### JSON Response

```go
app.Get("/api/user", func(c *fiber.Ctx) error {
    user := struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }{
        Name: "John Doe",
        Age:  30,
    }
    return c.JSON(user)
})
```

### Query Parameters

```go
// GET /search?query=fiber
app.Get("/search", func(c *fiber.Ctx) error {
    query := c.Query("query")
    return c.SendString("Search query: " + query)
})
```

### Route Parameters

```go
// GET /users/johndoe
app.Get("/users/:name", func(c *fiber.Ctx) error {
    name := c.Params("name")
    return c.SendString("Hello, " + name)
})
```

### POST Request with JSON Body

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

app.Post("/api/users", func(c *fiber.Ctx) error {
    req := new(CreateUserRequest)
    if err := c.BodyParser(req); err != nil {
        return err
    }
    // ... create user in database ...
    return c.Status(fiber.StatusCreated).JSON(req)
})
```

This module provides a solid foundation for building web APIs with Fiber. In the next module, we'll explore middleware and databases.

**Previous:** [Module 10: Client-Side Networking](../10/README.md)
**Next:** [Module 12: Middleware and Databases](../12/README.md)
