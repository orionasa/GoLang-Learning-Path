# 32. Routing in Fiber

Routing is the process of determining how an application responds to a client request to a particular endpoint, which is a URI (or path) and a specific HTTP request method (GET, POST, and so on). In this lesson, we'll explore how to handle routing in Fiber.

## Basic Routing

At its core, routing in Fiber involves binding a handler function to a specific path and HTTP method.

```go
app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("GET request")
})

app.Post("/", func(c *fiber.Ctx) error {
    return c.SendString("POST request")
})
```

Fiber provides methods for all the standard HTTP verbs: `Get`, `Post`, `Put`, `Delete`, `Patch`, `Head`, `Options`, etc.

## Route Parameters

Route parameters are named URL segments that are used to capture the values specified at their position in the URL. The captured values are populated in the `c.Params()` object.

```go
// Route with a required parameter
app.Get("/users/:id", func(c *fiber.Ctx) error {
    id := c.Params("id")
    return c.SendString("User ID: " + id)
})

// Route with an optional parameter
app.Get("/posts/:id?", func(c *fiber.Ctx) error {
    id := c.Params("id") // id will be empty if not provided
    if id == "" {
        return c.SendString("All posts")
    }
    return c.SendString("Post ID: " + id)
})
```

## Route Groups

Route groups are a way to organize related routes under a common prefix. This is useful for structuring larger applications. You can also apply middleware to an entire group of routes.

```go
api := app.Group("/api/v1") // Create a new group

api.Get("/users", func(c *fiber.Ctx) error {
    return c.SendString("All users")
})

api.Get("/users/:id", func(c *fiber.Ctx) error {
    // ...
})
```
In this example, the routes will be `/api/v1/users` and `/api/v1/users/:id`.

## Organizing Your Routes

For larger applications, it's a good practice to organize your routes into separate files or packages. Here's a common pattern for structuring your routes:

1.  Create a `routes` package.
2.  Inside the `routes` package, create a file for each major feature of your application (e.g., `user_routes.go`, `post_routes.go`).
3.  Each file will contain a function that registers the routes for that feature.
4.  In your `main.go` file, you call these functions to set up all your routes.

Here's an example of what `user_routes.go` might look like:

```go
// routes/user_routes.go
package routes

import (
    "github.com/gofiber/fiber/v2"
    "your-app/handlers" // assuming you have a handlers package
)

func UserRoutes(app *fiber.App) {
    api := app.Group("/api/v1/users")

    api.Get("/", handlers.GetUsers)
    api.Get("/:id", handlers.GetUser)
    api.Post("/", handlers.CreateUser)
}
```

And in your `main.go`:
```go
// main.go
package main

import (
    "github.com/gofiber/fiber/v2"
    "your-app/routes"
)

func main() {
    app := fiber.New()

    routes.UserRoutes(app)
    // ... other route setups

    app.Listen(":3000")
}
```
This approach helps in keeping your `main.go` file clean and your routes organized and maintainable.

Routing is a fundamental aspect of any web framework. Fiber's routing system is simple, powerful, and flexible, giving you all the tools you need to build well-structured web applications.

In the next lesson, we will take a closer look at handlers and the `fiber.Ctx` object.
