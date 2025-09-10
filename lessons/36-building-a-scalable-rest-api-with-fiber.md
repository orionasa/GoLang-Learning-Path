# 36. Building a Scalable REST API with Fiber

In this final lesson of our Fiber series, we'll put everything we've learned together to build a complete REST API. We'll focus on best practices for project structure, configuration, and database integration to ensure that our application is scalable and maintainable.

## The Goal

We will build a simple API for a blog. It will have two main resources: `users` and `posts`.

Our API will support the following endpoints:
*   `GET /api/v1/users`: Get all users
*   `GET /api/v1/users/:id`: Get a single user
*   `POST /api/v1/users`: Create a new user
*   `GET /api/v1/posts`: Get all posts
*   `GET /api/v1/posts/:id`: Get a single post
*   `POST /api/v1/posts`: Create a new post

## Project Structure

A well-organized project structure is key to building a scalable application. Here's a common structure for a Go web application:

```
my-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── database.go
│   ├── handlers/
│   │   ├── user_handler.go
│   │   └── post_handler.go
│   ├── models/
│   │   ├── user.go
│   │   └── post.go
│   └── routes/
│       └── routes.go
├── go.mod
└── go.sum
```

Let's break down this structure:
*   `cmd/main.go`: The entry point of our application. It initializes the database, router, and starts the server.
*   `internal/`: This directory contains all the core application code. The `internal` name is a special feature of Go that prevents other projects from importing these packages.
*   `config/`: For loading and managing application configuration.
*   `database/`: For database connection and setup.
*   `handlers/`: Contains the handler functions for our routes.
*   `models/`: Defines the data structures for our application (e.g., `User`, `Post`).
*   `routes/`: For setting up all the application routes.

## Implementation Sketch

Let's sketch out what the code in some of these files might look like.

### `models/user.go`
```go
package models

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

### `database/database.go`
```go
package database

import (
    "your-app/models"
    // imagine a database driver here
)

// For this example, we'll use an in-memory "database"
var Users = []models.User{}
var Posts = []models.Post{}
```
In a real application, this package would handle connecting to a database like PostgreSQL or MySQL.

### `handlers/user_handler.go`
```go
package handlers

import (
    "github.com/gofiber/fiber/v2"
    "your-app/database"
    "your-app/models"
)

func GetUsers(c *fiber.Ctx) error {
    return c.JSON(database.Users)
}

func CreateUser(c *fiber.Ctx) error {
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return err
    }
    database.Users = append(database.Users, *user)
    return c.JSON(user)
}
```

### `routes/routes.go`
```go
package routes

import (
    "github.com/gofiber/fiber/v2"
    "your-app/handlers"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api/v1")

    user := api.Group("/users")
    user.Get("/", handlers.GetUsers)
    user.Post("/", handlers.CreateUser)

    // ... routes for posts
}
```

### `cmd/main.go`
```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "your-app/routes"
)

func main() {
    app := fiber.New()

    // Setup routes
    routes.SetupRoutes(app)

    app.Listen(":3000")
}
```

## Best Practices for Scalability

*   **Configuration Management:** Externalize your configuration (e.g., database connection strings, port numbers) and load it from environment variables or a config file.
*   **Structured Logging:** Use a structured logging library to produce logs that are easy to parse and search.
*   **Graceful Shutdown:** Implement a graceful shutdown mechanism to allow your server to finish processing existing requests before it shuts down.
*   **Containerization:** Package your application in a container (e.g., Docker) for easy deployment and scaling.
*   **Database Connection Pooling:** Use a connection pool to manage database connections efficiently.

This lesson has provided a high-level overview of how to structure and build a scalable REST API with Fiber. By following these patterns and best practices, you can build robust and maintainable web applications that are ready to grow.

This concludes our series on Fiber. You now have the knowledge to build everything from simple web pages to complex, high-performance REST APIs. Happy coding!
