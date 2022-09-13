package router

import (
	"fmt"
	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/models"

	"github.com/gofiber/fiber/v2"
)

func FindALlPost(c *fiber.Ctx) error {
	posts := []models.Post{}

	database.Database.Find(&posts)

	return c.Status(200).JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := AddPost(post.CategoryRefer, &post); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	fmt.Println(post)

	database.Database.Create(&post)

	return c.Status(201).JSON(post)
}
