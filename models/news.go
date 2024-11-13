package models

import (
	"time"
)

type News struct {
	ID uint `gorm:"primaryKey"`
	Text string `gorm:"size:1000;not null"`
	CreatedAt time.Time `json:"-"`
}