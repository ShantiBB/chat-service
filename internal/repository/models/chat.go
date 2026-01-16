package models

import (
	"time"
)

type Chat struct {
	CreatedAt time.Time
	Title     string     `gorm:"size:200;not null"`
	Messages  []*Message `gorm:"-"`
	ID        uint       `gorm:"primaryKey"`
}
