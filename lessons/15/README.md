# Module 15: Advanced Security Practices

In this module, we'll discuss some advanced security practices that you should follow when building web applications in Go.

## Input Validation and Sanitization to prevent attacks

You should never trust user input. Always validate and sanitize any data that you receive from users.

For validation, you can use a library like `go-playground/validator`.

```bash
go get -u github.com/go-playground/validator/v10
```

Here's an example of how to use it to validate a struct:

```go
import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type User struct {
    Username string `validate:"required,alphanum"`
    Email    string `validate:"required,email"`
}

func main() {
    validate := validator.New()

    user := User{Username: "johndoe", Email: "invalid-email"}
    err := validate.Struct(user)
    if err != nil {
        fmt.Println(err)
    }
}
```

For sanitization, you can use a library like `microcosm-cc/bluemonday` to prevent Cross-Site Scripting (XSS) attacks.

## Managing secrets with Environment Variables

You should never hardcode secrets like API keys and database passwords in your code. Instead, you should store them in environment variables.

You can use a library like `joho/godotenv` to load environment variables from a `.env` file during development.

```bash
go get -u github.com/joho/godotenv
```

`.env` file:
```
DB_PASSWORD=mysecretpassword
```

Go code:
```go
import (
    "fmt"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    password := os.Getenv("DB_PASSWORD")
    fmt.Println(password)
}
```

In production, you would set the environment variables directly on your server or using a secret management service.

## Implementing Rate Limiting and basic DoS protection

Rate limiting is a technique that is used to control the amount of traffic that a server receives. It can be used to prevent Denial of Service (DoS) attacks and to ensure that your application is available to all users.

Fiber has a built-in rate limiter middleware.

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/limiter"
    "time"
)

func main() {
    app := fiber.New()

    app.Use(limiter.New(limiter.Config{
        Max:        100,
        Expiration: 1 * time.Minute,
    }))

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
```

This configuration will limit each IP address to 100 requests per minute.

By following these security best practices, you can build more secure and reliable web applications in Go.

**Previous:** [Module 14: Authentication and JWT](../14/README.md)
**Next:** [Module 16: Caching for Performance](../16/README.md)
