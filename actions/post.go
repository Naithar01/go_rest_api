package actions

import (
	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/models"
	"time"

	"github.com/lib/pq"
)

type ResponsePost struct {
	ID            uint           `json:"id"`
	Content       string         `json:"content"`
	CategoryRefer uint           `json:"category_id"`
	Tags          pq.StringArray `json:"tags"`
	CreatedAt     time.Time
}

type FindResponsePost struct {
	ID            uint            `json:"id"`
	Content       string          `json:"content"`
	CategoryRefer uint            `json:"category_id"`
	Category      models.Category `json:"category"`
	Tags          pq.StringArray  `json:"tags"`
	CreatedAt     time.Time
}

func CreateResponsePost(post models.Post) ResponsePost {
	return ResponsePost{ID: uint(post.ID), Content: post.Content, CategoryRefer: post.CategoryRefer, CreatedAt: post.CreatedAt, Tags: post.Tags}
}

func CreateFindResponsePost(post models.Post, category models.Category) FindResponsePost {
	return FindResponsePost{
		ID: uint(post.ID), Content: post.Content, CategoryRefer: post.CategoryRefer, Category: category, CreatedAt: post.CreatedAt, Tags: post.Tags,
	}
}

func CreateFindPostByCategoryId(posts *[]models.Post, category_id uint) {
	database.Database.Find(&posts, "category_refer = ?", category_id)
}
