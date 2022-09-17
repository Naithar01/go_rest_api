package actions

import (
	"github/com/Naithar01/go_rest_api/models"
	"time"
)

type ResponseCategory struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Posts     int    `json:"posts"`
	CreatedAt time.Time
}

type FindResponseCategory struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateResponseCategory(category models.Category, postLength int) ResponseCategory {
	return ResponseCategory{
		ID: category.ID, Title: category.Title, Posts: postLength, CreatedAt: category.CreatedAt,
	}
}

func CreateFindResponseCategory(category models.Category) FindResponseCategory {
	return FindResponseCategory{
		ID: category.ID, Title: category.Title, CreatedAt: category.CreatedAt, UpdatedAt: category.UpdatedAt,
	}
}
