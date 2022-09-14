package router

import (
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

func FindCategoryById(c *fiber.Ctx) error {
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

}
