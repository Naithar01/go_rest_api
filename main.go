package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/router"
)

func main() {

	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New())

	// Category
	app.Get("/api/category", router.FindAllCategory)
	app.Post("/api/category", router.CreateCategory)
	app.Get("/api/category/:id", router.FindCategoryById)
	app.Delete("/api/category/:id", router.DeleteCategory)

	// Post
	app.Get("/api/post", router.FindAllPost)
	app.Get("/api/post/ids", router.FindAllPostOnlyId)
	app.Post("/api/post", router.CreatePost)
	app.Get("/api/post/:id", router.FindPostById)
	app.Delete("/api/post/:id", router.DeletePost)
	app.Patch("/api/post/:id", router.UpdatePost)

	log.Fatal(app.Listen(":4000"))
}
