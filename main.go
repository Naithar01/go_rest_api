package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Posts []Post `json:"posts"`
}

type Post struct {
	Id int `json:"id"`
}

func main() {
	var categorys = []Category{}
	var posts = []Post{}

	app := fiber.New()

	app.Get("/api/category", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(categorys)
	})

	app.Get("/api/post", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(posts)
	})

	app.Post("/api/category", func(c *fiber.Ctx) error {
		var newCategory = &Category{}

		if err := c.BodyParser(newCategory); err != nil {
			return err
		}

		newCategory.Id = len(categorys) + 1

		categorys = append(categorys, *newCategory)

		return c.Status(201).JSON(newCategory)
	})

	app.Post("/api/post", func(c *fiber.Ctx) error {
		var newPost = &Post{}

		if err := c.BodyParser(newPost); err != nil {
			return err
		}

		newPost.Id = len(posts) + 1

		posts = append(posts, *newPost)

		return c.Status(201).JSON(newPost)
	})

	app.Get("/api/post/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for index, post := range posts {
			if post.Id == id {
				return c.JSON(posts[index])
			}
		}

		return c.Status(401).SendString("Cant find Post")

	})

	log.Fatal(app.Listen(":4000"))
}
