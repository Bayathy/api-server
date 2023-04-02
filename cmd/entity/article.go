package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title string
	Url   string
	Done  bool
	User  User `gorm:"embedded"`
}
