package models

type Post struct {
	Id            uint     `json:"id" gorm:"primaryKey"`
	Content       string   `json:"content"`
	CategoryRefer uint     `json:"category_id" `
	Category      Category `json:"category" gorm:"foreignKey:CategoryRefer"`
}
