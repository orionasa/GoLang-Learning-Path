package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Endpoint to get a list of all lessons
	app.Get("/api/lessons", func(c *fiber.Ctx) error {
		files, err := os.ReadDir("../../lessons")
		if err != nil {
			return err
		}

		var lessons []string
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
				lessons = append(lessons, file.Name())
			}
		}

		return c.JSON(lessons)
	})

	// Endpoint to get the content of a specific lesson
	app.Get("/api/lessons/:filename", func(c *fiber.Ctx) error {
		filename := c.Params("filename")
		// Sanitize filename to prevent directory traversal
		sanitizedFilename := filepath.Base(filename)
		if sanitizedFilename != filename || strings.Contains(sanitizedFilename, "..") {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid filename")
		}

		content, err := os.ReadFile(filepath.Join("../../lessons", sanitizedFilename))
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Lesson not found")
		}

		return c.SendString(string(content))
	})

	log.Fatal(app.Listen(":3000"))
}
