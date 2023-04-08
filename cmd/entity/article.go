package entity

import (
	"time"
)

type Article struct {
	Id        int32 `gorm:"primaryKey"`
	CreatedAt time.Time
	Title     string
	Url       string
	Done      bool
	UserId    string
}
