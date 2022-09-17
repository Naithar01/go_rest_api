package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	Content       string         `json:"content"`
	CategoryRefer uint           `json:"category_id"`
	Category      Category       `json:"category" gorm:"foreignKey:CategoryRefer"`
	Tags          pq.StringArray `json:"tags" gorm:"type:text"`
}
