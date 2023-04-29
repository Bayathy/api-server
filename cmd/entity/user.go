package entity

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Uuid      string    `gorm:"uniqueIndex"`
	Articles  []Article `gorm:"foreignKey:UserId;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
