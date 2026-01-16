package models

import (
	"time"
)

type Message struct {
	CreatedAt time.Time
	Text      string `gorm:"size:5000;not null"`
	Chat      Chat   `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;"`
	ChatID    uint   `gorm:"not null;index"`
	ID        uint   `gorm:"primaryKey"`
}
