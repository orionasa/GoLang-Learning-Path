# Module 14: Authentication and JWT

In this module, we'll cover how to implement authentication in a Go web application using JSON Web Tokens (JWT).

## Storing Passwords Securely with `bcrypt`

You should never store passwords in plain text. The `bcrypt` package provides a simple and secure way to hash and salt passwords.

Install the `bcrypt` package:

```bash
go get -u golang.org/x/crypto/bcrypt
```

Here's how to use it:

```go
import "golang.org/x/crypto/bcrypt"

// Hash a password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// Check if a password is correct
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

## Understanding JSON Web Tokens (JWT)

JWT is an open standard for creating access tokens that assert some number of claims. For example, a server could generate a token that has the claim "logged in as admin" and provide that to a client. The client could then use that token to prove that it is logged in as an admin.

A JWT consists of three parts: a header, a payload, and a signature. The header and payload are JSON objects that are Base64Url encoded. The signature is created by signing the header and payload with a secret key.

## Implementing JWT-based authentication flows in Fiber

We'll use the `github.com/golang-jwt/jwt/v4` package to work with JWTs.

Install the package:
```bash
go get -u github.com/golang-jwt/jwt/v4
```

Here's an example of how to create a JWT and use it to protect a route in Fiber:

```go
package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// In a real application, this would be in a config file
var jwtSecret = []byte("my-secret-key")

func login(c *fiber.Ctx) error {
	// For simplicity, we'll assume the user is authenticated
	// In a real application, you would check the user's credentials

	// Create the JWT claims
	claims := jwt.MapClaims{
		"username": "johndoe",
		"admin":    true,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func protected(c *fiber.Ctx) error {
	// In a real app, you would have middleware to extract the token
	// and validate it. For simplicity, we'll just get the user from the context.
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return c.SendString("Welcome " + name)
}

func main() {
	app := fiber.New()

	app.Post("/login", login)

	// This is a simplified example. In a real app, you would use middleware
	// to protect routes.
	app.Get("/protected", protected)

	app.Listen(":3000")
}
```

In a real application, you would create a middleware to extract the JWT from the `Authorization` header, validate it, and then store the user's information in the Fiber context.

This module has provided a basic overview of authentication and JWTs in Go. You can now implement a secure authentication system for your web applications.

**Previous:** [Module 13: Database Mastery with sqlc](../13/README.md)
**Next:** [Module 15: Advanced Security Practices](../15/README.md)
