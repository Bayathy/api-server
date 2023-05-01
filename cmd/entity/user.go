package entity

import (
	"time"
)

type User struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	Uuid      string    `gorm:"uniqueIndex"`
	Articles  []Article `gorm:"foreignKey:UserId;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
