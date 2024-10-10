package main

import (
	"fmt"
	"log"

	"example.com/eriktest/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New()

	app.Get("/health-check", healthCheck)
	app.Post("api/products", handlers.CreateProduct)

	fmt.Println("Hello world")

	log.Fatal(app.Listen(":3000"))
}
