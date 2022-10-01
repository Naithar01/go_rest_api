package router

import (
	"github/com/Naithar01/go_rest_api/controller"

	"github.com/gofiber/fiber/v2"
)

func InitApp() *fiber.App {
	app := fiber.New()

	// Category
	app.Get("/api/category", controller.FindAllCategory)
	app.Post("/api/category", controller.CreateCategory)
	app.Get("/api/category/:id", controller.FindCategoryById)
	app.Delete("/api/category/:id", controller.DeleteCategory)

	// Post
	app.Get("/api/post", controller.FindAllPost)
	// // Search Post By Content
	app.Get("/api/post/search/content", controller.SearchPostByContent)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/post/:id", controller.FindPostById)
	app.Delete("/api/post/:id", controller.DeletePost)
	app.Patch("/api/post/:id", controller.UpdatePost)

	return app
}
