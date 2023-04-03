package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title  string
	Url    string
	Done   bool
	UserId int
	User   User `gorm:"foreignKey:UserId"`
}
