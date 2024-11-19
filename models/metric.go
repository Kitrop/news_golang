package models

import "time"

type Client_metadata struct {
	ID      uint   `gorm:"primaryKey"`
	Ip      string `gorm:"not null"`
	Os      string `gorm:"size:128; not null"`
	Browser string `gorm:"size:128; not null"`
	Device  string `gorm:"size:128; not null"`
}

type RequestMetric struct {
	ID           uint      `gorm:"primaryKey"`
	Path         string    `gorm:"index;not null"` // URL пути
	Method       string    `gorm:"not null"`       // HTTP метод
	StatusCode   int       `gorm:"not null"`       // Код ответа
	ResponseTime float64   `gorm:"not null"`
	Timestamp    time.Time `gorm:"default:current_timestamp"`
}