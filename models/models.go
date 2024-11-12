package models

import (
	"time"
)

type News struct {
	ID uint `gorm:"primaryKey"`
	Text string `gorm:"size:1000;not null"`
	CreatedAt time.Time `json:"-"`
}

type RoleStatus string

const (
	Admin RoleStatus = "ADMIN"
	Active RoleStatus = "ACTIVE"
	Not_Active RoleStatus = "NOT_ACTIVE"
	Banned RoleStatus = "BANNED"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"unique;size:20;not null"`
	Password string `gorm:"not null"`
	Email string `gorm:"unique;size:100;not null"`
	Role RoleStatus `gorm:"varchar(15);default:'NOT_ACTIVE'" json:"role"`
}