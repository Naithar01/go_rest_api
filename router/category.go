package router

import (
	"errors"
	"fmt"
	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/models"

	"github.com/gofiber/fiber/v2"
)

func FindAllCategory(c *fiber.Ctx) error {
	categorys := []models.Category{}

	database.Database.Find(&categorys)

	return c.Status(200).JSON(categorys)
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	database.Database.Create(&category)

	return c.Status(201).JSON(category)
}

func AddPost(id int, post *models.Post) error {
	database.Database.Find(&post.Category, "id = ?", id)
	fmt.Println(id)
	if post.Category.Id == 0 {
		return errors.New("order does not exist")
	}

	return nil
}
