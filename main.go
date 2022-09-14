package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/router"
)

func main() {

	database.ConnectDB()

	app := fiber.New()

	app.Get("/api/category", router.FindAllCategory)

	app.Post("/api/category", router.CreateCategory)

	app.Get("/api/category/:id", router.FindCategoryById)

	app.Get("/api/post", router.FindAllPost)

	app.Post("/api/post", router.CreatePostfunc)

	app.Get("/api/post/:id", router.FindPostById)

	log.Fatal(app.Listen(":4000"))
}
