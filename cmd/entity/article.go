package entity

import (
	"time"
)

type Article struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	Title     string
	Url       string
	Done      bool
	UserId    int
	User      User `gorm:"foreignKey:UserId"`
}
