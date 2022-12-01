package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/hello", hello)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(500) // 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("I made a 🙊 for you.")
}
