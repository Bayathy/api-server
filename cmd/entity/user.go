package entity

import (
	"time"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	Uuid      string
}
