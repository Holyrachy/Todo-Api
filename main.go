package main

import (
	"log"
	"todo-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Middleware to handle OPTIONS requests
	app.Use(func(c *fiber.Ctx) error {
		origin := c.Get("Origin")
		// Respond to preflight requests
		if c.Method() == fiber.MethodOptions {

			c.Set("Access-Control-Allow-Origin", origin)
			c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Set("Access-Control-Expose-Headers", "Content-Type, Authorization")
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		origin := c.Get("Origin")

		c.Set("Access-Control-Allow-Origin", origin)
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Expose-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Max-Age", "3600")
		return c.Next()
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
