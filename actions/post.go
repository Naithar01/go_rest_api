package actions

import (
	"github/com/Naithar01/go_rest_api/models"
	"time"
)

type ResponsePost struct {
	Id            int    `json:"id"`
	Content       string `json:"content"`
	CategoryRefer uint   `json:"category_id"`
	CreatedAt     time.Time
}

type FindResponsePost struct {
	Id            int             `json:"id"`
	Content       string          `json:"content"`
	CategoryRefer uint            `json:"category_id"`
	Category      models.Category `json:"category"`
	CreatedAt     time.Time
}

func CreateResponsePost(post models.Post) ResponsePost {
	return ResponsePost{Id: int(post.Id), Content: post.Content, CategoryRefer: post.CategoryRefer, CreatedAt: post.CreatedAt}
}

func CreateFindResponsePost(post models.Post, category models.Category) FindResponsePost {
	return FindResponsePost{
		Id: int(post.Id), Content: post.Content, CategoryRefer: post.CategoryRefer, Category: category, CreatedAt: post.CreatedAt,
	}
}
