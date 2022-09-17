package router

import (
	"errors"
	"fmt"
	"github/com/Naithar01/go_rest_api/actions"
	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/models"

	"github.com/gofiber/fiber/v2"
)

type FindAllPostQuery struct {
	Category_id uint `query:"category_id"`
}

func AddPost(id uint, post *models.Post) error {
	var category models.Category

	database.Database.Find(&category, "id = ?", id)

	post.Category = category

	if category.ID == 0 {
		return errors.New("order does not exist")
	}

	return nil
}

func FindAllPost(c *fiber.Ctx) error {
	category_query := new(FindAllPostQuery)

	if err := c.QueryParser(category_query); err != nil {
		return c.Status(401).JSON(err.Error())
	}

	posts := []models.Post{}

	if category_query.Category_id != 0 {
		actions.CreateFindPostByCategoryIdResponse(&posts, category_query.Category_id)

		fmt.Println(category_query)

		responsePosts := []actions.ResponsePost{}

		for _, post := range posts {
			responsePost := actions.CreateResponsePost(post)
			responsePosts = append(responsePosts, responsePost)
		}

		return c.Status(200).JSON(responsePosts)

	}

	database.Database.Find(&posts)

	responsePosts := []actions.ResponsePost{}

	for _, post := range posts {
		responsePost := actions.CreateResponsePost(post)
		responsePosts = append(responsePosts, responsePost)
	}

	return c.Status(200).JSON(responsePosts)
}

func CreatePost(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := AddPost(post.CategoryRefer, &post); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if len(post.Tags) == 0 {
		post.Tags = nil
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

func DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(401).SendString("Validate Post Id")
	}

	database.Database.Unscoped().Where("id = ?", id).Delete(&models.Post{})

	return c.Status(200).SendString("Success")

}
