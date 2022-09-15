package router

import (
	"errors"
	"github/com/Naithar01/go_rest_api/actions"
	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/models"

	"github.com/gofiber/fiber/v2"
)

type FindAllPostQuery struct {
	Category_id string `query:"category_id"`
}

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
	// db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	category_query := new(FindAllPostQuery)

	if err := c.QueryParser(category_query); err != nil {
		return c.Status(401).JSON(err.Error())
	}

	posts := []models.Post{}

	if len(category_query.Category_id) != 0 {
		database.Database.Find(&posts, "category_refer = ?", category_query.Category_id)

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
