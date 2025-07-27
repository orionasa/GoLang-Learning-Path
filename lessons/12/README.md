# Module 12: Middleware and Databases

In this module, we'll learn about middleware in Fiber and how to connect our Go application to a database.

## Using Middleware in Fiber (Logging, Recovery, CORS)

Middleware functions are functions that have access to the request and response objects. They can be used to perform tasks like logging, authentication, and error handling.

Fiber comes with a set of built-in middleware. Here's how to use the logger middleware:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()

    // Use the logger middleware
    app.Use(logger.New())

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
```

Other useful middleware include:

*   **`recover`**: Recovers from panics anywhere in the stack.
*   **`cors`**: Enables Cross-Origin Resource Sharing (CORS).

You can also create your own custom middleware.

## Introduction to the `database/sql` package

The `database/sql` package provides a generic SQL interface. It should be used in conjunction with a database driver for the specific database you are using.

We'll be using PostgreSQL in this guide, so you'll need to install the `pgx` driver:

```bash
go get -u github.com/jackc/pgx/v4/stdlib
```

## Setting up PostgreSQL (or MySQL) and connecting from Go

First, you'll need to have a PostgreSQL server running. You can use Docker to easily set one up.

Here's how to connect to a PostgreSQL database from a Go application:

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/jackc/pgx/v4/stdlib" // The pgx driver
)

func main() {
    // Replace with your actual connection string
    connStr := "user=user password=password dbname=mydb sslmode=disable"
    db, err := sql.Open("pgx", connStr)
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }
    defer db.Close()

    // Ping the database to check if the connection is working
    err = db.Ping()
    if err != nil {
        fmt.Println("Error pinging the database:", err)
        return
    }

    fmt.Println("Successfully connected to the database!")

    // Now you can execute queries
    rows, err := db.Query("SELECT version()")
    if err != nil {
        // handle error
    }
    defer rows.Close()

    var version string
    for rows.Next() {
        err := rows.Scan(&version)
        if err != nil {
            // handle error
        }
        fmt.Println(version)
    }
}
```

In this module, we've learned how to use middleware in Fiber and how to connect to a PostgreSQL database. In the next module, we'll look at a more powerful way to interact with our database using `sqlc`.

**Previous:** [Module 11: Web & API Foundations with Fiber](../11/README.md)
**Next:** [Module 13: Database Mastery with sqlc](../13/README.md)
