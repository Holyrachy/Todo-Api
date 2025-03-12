package main

import (
	"log"
	"todo-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
