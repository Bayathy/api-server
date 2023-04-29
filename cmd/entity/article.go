package entity

import (
	"time"
)

type Article struct {
	ID        uint `gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time
	Title     string
	Url       string
	Done      bool
	UserId    string
}
