package actions

import "github/com/Naithar01/go_rest_api/models"

type ResponseCategory struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Posts int    `json:"posts"`
}

func CreateResponseCategory(category models.Category, postLength int) ResponseCategory {
	return ResponseCategory{
		Id: category.Id, Title: category.Title, Posts: postLength,
	}
}
