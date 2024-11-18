package models

import (
	"time"
)

type News struct {
	ID uint `gorm:"primaryKey"`
	Text string `gorm:"size:1000;not null" binding:"required"`
	CreatedAt time.Time `json:"-"`
}