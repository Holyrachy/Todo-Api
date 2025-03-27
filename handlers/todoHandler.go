package handlers

import (
	"fmt"
	"strconv"
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
)

// Simulated in-memory storage
var todos = []models.Todo{}
var idCounter = 1

// Get all todos
func GetTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

// Get a single todo by ID
func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID not valid"})
	}

	for _, todo := range todos {
		if todo.ID == idConv {
			return c.JSON(todo)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
}

// Create a new todo
func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	todo.ID = idCounter
	idCounter++
	todos = append(todos, todo)
	return c.Status(201).JSON(todo)
}

// Update a todo
func UpdateTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID not valid"})
	}

	for i, todo := range todos {
		if idConv == todo.ID {
			if err := c.BodyParser(&todos[i]); err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
			}
			return c.JSON(todos[i])
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
}

// Delete a todo
func DeleteTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID not valid"})
	}

	for i, todo := range todos {
		if idConv == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)
			return c.Status(200).JSON(fiber.Map{"message": fmt.Sprintf("Todo with ID %v is deleted", id)})
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
}
