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
	Id           int    `json:"id"`
	Content      string `json:"content"`
	CategoryName string `json:"categoryName"`
}

func middlewareCategory(c *fiber.Ctx) error {
	log.Println("middleWare Category pass")
	return c.Next()
}

func middlewarePost(c *fiber.Ctx) error {
	log.Println("middleWare Post pass")
	return c.Next()
}

func main() {
	var categorys = []Category{}
	var posts = []Post{}

	app := fiber.New()

	api := app.Group("/api")

	categoryRouter := api.Group("/category", middlewareCategory)
	postRouter := api.Group("/post", middlewarePost)

	categoryRouter.Get("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(categorys)
	})

	postRouter.Get("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(posts)
	})

	categoryRouter.Post("", func(c *fiber.Ctx) error {
		var newCategory = &Category{}

		if err := c.BodyParser(newCategory); err != nil {
			return err
		}

		newCategory.Id = len(categorys) + 1

		categorys = append(categorys, *newCategory)

		return c.Status(201).JSON(newCategory)
	})

	postRouter.Post("/:categoryId", func(c *fiber.Ctx) error {
		categoryId, err := c.ParamsInt("categoryId")

		if err != nil {
			return c.Status(401).SendString("Cant find Category Id")
		}

		var newPost = &Post{}

		if err := c.BodyParser(newPost); err != nil {
			return c.Status(401).SendString("Some Error, Fail Create Post")
		}

		for index, category := range categorys {
			if category.Id == categoryId {
				newPost.Id = len(posts) + 1
				newPost.CategoryName = category.Title
				posts = append(posts, *newPost)

				categorys[index].Posts = append(categorys[index].Posts, *newPost)

				return c.Status(201).JSON(category)

			}
		}

		return c.Status(401).SendString("Some Error, Fail Create Post")
	})

	postRouter.Get("/:id", func(c *fiber.Ctx) error {
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
