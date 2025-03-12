package routes

import (
	"todo-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/todos", handlers.GetTodos)
	app.Get("/todos/:id", handlers.GetTodoByID)
	app.Post("/todos", handlers.CreateTodo)
	app.Put("/todos/:id", handlers.UpdateTodoByID)
	app.Delete("/todos/:id", handlers.DeleteTodoByID)
}
