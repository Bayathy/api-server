package entity

import (
	"time"
)

type Article struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	Title     string
	Url       string
	Done      bool
	UserId    string
}
