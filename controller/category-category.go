package controller

import (
	"github/com/Naithar01/go_rest_api/actions"
	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/models"

	"github.com/gofiber/fiber/v2"
)

func FindAllCategory(c *fiber.Ctx) error {
	categorys := []models.Category{}

	database.Database.Find(&categorys)

	responseCategorys := []actions.ResponseCategory{}

	posts := []models.Post{}

	for _, category := range categorys {
		actions.CreateFindPostByCategoryId(&posts, category.ID)

		responseCategory := actions.CreateResponseCategory(category, len(posts))
		responseCategorys = append(responseCategorys, responseCategory)
	}

	return c.Status(200).JSON(responseCategorys)
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

	if category.ID == 0 {
		return c.Status(401).SendString("Validate Category id")
	}

	findResponseCategory := actions.CreateFindResponseCategory(category)

	return c.Status(200).JSON(findResponseCategory)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(401).SendString("Validate Category Id")
	}

	database.Database.Unscoped().Where("id = ?", id).Delete(&models.Category{})

	return c.Status(200).SendString("Success")
}
