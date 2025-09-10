# 34. Middleware in Fiber

Middleware are functions that are executed before the main handler for a route. They are used to handle "cross-cutting concerns" such as logging, authentication, CORS, and more. In Fiber, middleware are just regular handlers that can choose to pass control to the next handler in the chain or to terminate the request-response cycle.

## Using Middleware

You can register middleware to run for all routes or for specific routes or groups.

### Application-Level Middleware

To use middleware for all routes in your application, you can use `app.Use()`.

```go
app.Use(func(c *fiber.Ctx) error {
    fmt.Println("This middleware runs for every request")
    return c.Next()
})
```
The `c.Next()` function passes control to the next handler in the chain. If a middleware does not call `c.Next()`, the request will not be passed to the main handler.

### Route-Level Middleware

You can also apply middleware to a specific route or a group of routes.

```go
// Middleware for a single route
app.Get("/secret", myAuthMiddleware, func(c *fiber.Ctx) error {
    return c.SendString("This is a secret page.")
})

// Middleware for a group of routes
api := app.Group("/api", myApiMiddleware)
api.Get("/users", ...)
```

## Built-in Middleware

Fiber comes with a rich set of official middleware that you can use in your application. These middleware are available in the `github.com/gofiber/fiber/v2/middleware` package.

### Logger

The logger middleware logs information about each request, such as the IP address, method, path, and status code.

```go
import "github.com/gofiber/fiber/v2/middleware/logger"

app.Use(logger.New())
```

### CORS

The CORS (Cross-Origin Resource Sharing) middleware enables cross-origin requests.

```go
import "github.com/gofiber/fiber/v2/middleware/cors"

app.Use(cors.New())
```

### Recover

The recover middleware recovers from panics anywhere in the stack chain and sends a `500` error. This prevents your server from crashing due to an unhandled panic.

```go
import "github.com/gofiber/fiber/v2/middleware/recover"

app.Use(recover.New())
```

## Custom Middleware

Writing your own middleware is easy. A middleware is just a function that follows the `fiber.Handler` signature.

Here's an example of a simple authentication middleware that checks for an API key in the request header.

```go
func authMiddleware(c *fiber.Ctx) error {
    apiKey := c.Get("X-API-KEY")

    // In a real application, you would check the API key against a database.
    if apiKey != "my-secret-key" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid API Key",
        })
    }

    return c.Next()
}

app.Get("/protected", authMiddleware, func(c *fiber.Ctx) error {
    return c.SendString("You have access to the protected area.")
})
```

In this example, the `authMiddleware` checks for the `X-API-KEY` header. If the key is not valid, it sends a `401 Unauthorized` response and terminates the request. If the key is valid, it calls `c.Next()` to pass control to the main handler.

Middleware is a powerful feature that helps you write clean, modular, and reusable code. By composing middleware, you can build complex and robust web applications with ease.

In the next lesson, we will explore some advanced topics in Fiber, such as error handling and validation.
