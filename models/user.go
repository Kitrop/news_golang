package models

type RoleStatus string

const (
    Admin      RoleStatus = "ADMIN"
    Active     RoleStatus = "ACTIVE"
    NotActive  RoleStatus = "NOT_ACTIVE"
    Banned     RoleStatus = "BANNED"
)

type User struct {
    ID       uint       `gorm:"primaryKey"`
    Username string     `gorm:"unique;size:20;not null" binding:"required"`
    Password string     `gorm:"not null" binding:"required"`
    Email    string     `gorm:"unique;size:100;not null" binding:"required"`
    Role     RoleStatus `gorm:"varchar(15);default:'NOT_ACTIVE'" json:"role"`
}

type UserLogin struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}