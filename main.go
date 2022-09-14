package main

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"

	"github/com/Naithar01/go_rest_api/actions"
	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/models"
)

func AddPost(id uint, post *models.Post) error {
	var category models.Category

	database.Database.Find(&category, "id = ?", id)

	post.Category = category

	if category.Id == 0 {
		return errors.New("order does not exist")
	}

	return nil
}

func main() {

	database.ConnectDB()

	app := fiber.New()

	app.Get("/api/category", func(c *fiber.Ctx) error {
		categorys := []models.Category{}

		database.Database.Find(&categorys)

		return c.Status(200).JSON(categorys)
	})

	app.Post("/api/category", func(c *fiber.Ctx) error {
		var category models.Category

		if err := c.BodyParser(&category); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		database.Database.Create(&category)

		return c.Status(201).JSON(category)
	})

	app.Get("/api/category/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Validate Category id")
		}

		var category models.Category

		database.Database.Find(&category, "id = ?", id)

		if category.Id == 0 {
			return c.Status(401).SendString("Validate Category id")
		}

		return c.Status(200).JSON(category)

	})

	app.Get("/api/post", func(c *fiber.Ctx) error {
		posts := []models.Post{}

		database.Database.Find(&posts)

		responsePosts := []actions.ResponsePost{}

		for _, post := range posts {
			responsePost := actions.CreateResponsePost(post)
			responsePosts = append(responsePosts, responsePost)
		}

		return c.Status(200).JSON(responsePosts)
	})

	app.Post("/api/post", func(c *fiber.Ctx) error {
		var post models.Post

		if err := c.BodyParser(&post); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		if err := AddPost(post.CategoryRefer, &post); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		database.Database.Create(&post)

		return c.Status(201).JSON(post)
	})

	app.Get("/api/post/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Validate Post Id")
		}

		var post models.Post

		database.Database.Find(&post, "id = ?", id)

		if post.Id == 0 {
			return c.Status(401).SendString("Validate Post Id")
		}

		return c.Status(200).JSON(post)
	})

	log.Fatal(app.Listen(":4000"))
}
