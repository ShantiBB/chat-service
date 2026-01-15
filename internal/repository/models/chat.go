package models

import (
	"time"
)

type Chat struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:200;not null"`
	Messages  []Message `gorm:"-"`
	CreatedAt time.Time
}
