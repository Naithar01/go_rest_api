package models

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	Id            uint           `json:"id" gorm:"primaryKey"`
	Content       string         `json:"content"`
	CategoryRefer uint           `json:"category_id"`
	Category      Category       `json:"category" gorm:"foreignKey:CategoryRefer"`
	Tags          pq.StringArray `json:"tags" gorm:"type:text"`
	CreatedAt     time.Time
	UpdateAt      time.Time
}
