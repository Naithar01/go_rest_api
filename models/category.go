package models

type Category struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}
