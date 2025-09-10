# 31. Introduction to Fiber

Welcome to our series on Fiber, a high-performance web framework for Go. If you're looking to build fast and scalable web applications, APIs, and microservices, you're in the right place.

## What is Fiber?

Fiber is a web framework for Go that is inspired by Express.js, one of the most popular web frameworks for Node.js. This means that if you have experience with Express, you'll find Fiber's API to be very familiar and intuitive.

But Fiber is more than just a Go version of Express. It's built on top of Fasthttp, the fastest HTTP engine for Go. This gives Fiber a significant performance advantage over the standard `net/http` package that many other Go web frameworks use.

## Core Philosophy

The core philosophy of Fiber is to be:
*   **Fast:** Fiber is designed for zero memory allocation and high performance.
*   **Simple:** The API is designed to be simple and easy to use, especially for developers coming from a Node.js background.
*   **Flexible:** Fiber provides a robust set of features, but it doesn't get in your way. You can use as much or as little of the framework as you need.

## Key Features

Here are some of the key features that make Fiber a great choice for your next project:
*   **High Performance:** Thanks to Fasthttp, Fiber is one of the fastest web frameworks in the Go ecosystem.
*   **Express-like API:** The API is simple, familiar, and easy to learn.
*   **Robust Routing:** Fiber has a powerful and flexible router that supports parameters, groups, and middleware.
*   **Extensive Middleware Support:** Fiber comes with a rich set of built-in middleware for common tasks like logging, authentication, and CORS. It's also easy to write your own.
*   **Template Engine Support:** Fiber supports multiple template engines for server-side rendering.
*   **WebSocket Support:** Built-in support for WebSockets makes it easy to build real-time applications.

## "Hello, World!" with Fiber

Let's look at a simple "Hello, World!" example to see how easy it is to get started with Fiber.

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
```

As you can see, the code is very similar to what you would write in Express.js. We create a new Fiber app, define a route for the root URL, and start the server.

In the upcoming lessons, we will dive deeper into the features of Fiber and learn how to build a complete web application with it. We'll start with routing in the next lesson.
