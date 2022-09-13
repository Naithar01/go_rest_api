package models

import "time"

type Category struct {
	Id        int32   `gorm:"primaryKey"`
	Title     string  `json:"title"`
	Posts     []*Post `json:"posts" gorm:"foreignKey:CategoryRefer"`
	CreatedAt time.Time
}
