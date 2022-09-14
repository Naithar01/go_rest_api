package router

import (
	"errors"
	"github/com/Naithar01/go_rest_api/actions"
	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/models"

	"github.com/gofiber/fiber/v2"
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

func FindAllPost(c *fiber.Ctx) error {
	posts := []models.Post{}

	database.Database.Find(&posts)

	responsePosts := []actions.ResponsePost{}

	for _, post := range posts {
		responsePost := actions.CreateResponsePost(post)
		responsePosts = append(responsePosts, responsePost)
	}

	return c.Status(200).JSON(responsePosts)
}

func CreatePostfunc(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := AddPost(post.CategoryRefer, &post); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	database.Database.Create(&post)

	return c.Status(201).JSON(post)
}

func FindPostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(401).SendString("Validate Post Id")
	}

	var post models.Post

	database.Database.Find(&post, "id = ?", id)

	if post.Id == 0 {
		return c.Status(401).SendString("Validate Post Id")
	}

	var category models.Category

	database.Database.Find(&category, "id = ?", post.CategoryRefer)

	return c.Status(200).JSON(actions.CreateFindResponsePost(post, category))
}
