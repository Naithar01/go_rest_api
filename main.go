package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	var todos = []Todo{}

	app := fiber.New()

	app.Get("/api/todo", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Post("/api/todo", func(c *fiber.Ctx) error {
		var newTodo = &Todo{}

		if err := c.BodyParser(newTodo); err != nil {
			return err
		}

		newTodo.ID = len(todos) + 1

		todos = append(todos, *newTodo)

		return c.Status(201).JSON(todos)
	})

	app.Patch("/api/todo/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for index, todo := range todos {
			if todo.ID == id {
				todos[index].Done = !todos[index].Done
				break
			}
		}

		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
