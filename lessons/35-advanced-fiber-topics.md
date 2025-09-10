# 35. Advanced Fiber Topics

Now that you have a good understanding of the basics of Fiber, let's explore some of its more advanced features that will help you build robust and full-featured web applications.

## Custom Error Handling

By default, Fiber's error handler will catch any errors returned from your handlers and send a generic `500 Internal Server Error` response. You can create a custom error handler to have more control over how errors are handled.

You can set a custom error handler in the Fiber config:

```go
import "errors"

app := fiber.New(fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        // Default error status code
        code := fiber.StatusInternalServerError

        // Retrieve the custom status code if it's a *fiber.Error
        var e *fiber.Error
        if errors.As(err, &e) {
            code = e.Code
        }

        // Send custom error page
        return c.Status(code).JSON(fiber.Map{
            "error": err.Error(),
        })
    },
})
```

## Data Validation

Validating incoming data is a crucial part of any web application. While Fiber doesn't have a built-in validation library, it's easy to integrate with popular validation libraries like `go-playground/validator`.

Here's an example of how you might use it:

```go
import "github.com/go-playground/validator/v10"

type User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

var validate = validator.New()

app.Post("/users", func(c *fiber.Ctx) error {
    user := new(User)

    if err := c.BodyParser(user); err != nil {
        return err
    }

    if err := validate.Struct(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // ... create the user
})
```

## Serving Static Files

Fiber makes it easy to serve static files like HTML, CSS, and JavaScript. You can use the `app.Static()` method for this.

```go
// Serve files from the "./public" directory at the root of the site
app.Static("/", "./public")

// Serve files from "./public" directory on the "/static" path
app.Static("/static", "./public")
```

## Template Engines

Fiber has built-in support for multiple template engines, allowing you to render dynamic HTML on the server. To use a template engine, you need to create a new engine and pass it to the Fiber config.

Here's an example using Go's built-in `html/template` engine:

```go
import "github.com/gofiber/template/html/v2"

// Create a new template engine
engine := html.New("./views", ".html")

app := fiber.New(fiber.Config{
    Views: engine,
})

app.Get("/", func(c *fiber.Ctx) error {
    // Render the index.html template with some data
    return c.Render("index", fiber.Map{
        "Title": "Hello, World!",
    })
})
```
You'll need to have a `views` directory with an `index.html` file for this to work.

Fiber provides a rich set of features that go beyond the basics of routing and handlers. By leveraging these advanced features, you can build powerful, secure, and full-featured web applications with ease.

In the next and final lesson of this series, we will put everything we've learned together to build a complete REST API with Fiber.
