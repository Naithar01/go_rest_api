package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello World")

	app := fiber.New()

	app.Get("/helloworld", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	log.Fatal(app.Listen(":4000"))
}
