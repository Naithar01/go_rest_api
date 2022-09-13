package models

import "time"

type Post struct {
	Id            int      `gorm:"primaryKey"`
	Content       string   `json:"content"`
	CategoryRefer int      `json:"category_id"`
	Category      Category `gorm:"foreignKey:CategoryRefer"`
	CreatedAt     time.Time
	UpdateAt      time.Time
}
