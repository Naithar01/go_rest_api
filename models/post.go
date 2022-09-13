package models

import "time"

type Post struct {
	Id            int      `json:"id" gorm:"primaryKey"`
	Content       string   `json:"content"`
	CategoryRefer int      `json:"category_id"`
	Category      Category `json:"category" gorm:"foreignKey:CategoryRefer"`
	CreatedAt     time.Time
	UpdateAt      time.Time
}
