package models

type Client_metadata struct {
	ID uint `gorm:"primaryKey"`
	Ip string `gorm:"not null"`
	Os string `gorm:"size:128; not null"`
	Browser string `gorm:"size:128; not null"`
	Device string `gorm:"size:128; not null"`
}