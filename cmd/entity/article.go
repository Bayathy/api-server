package entity

import (
	"time"
)

type Article struct {
	Id        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	Url       string
	Done      bool
	UserId    string
}
