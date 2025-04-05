package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null" json:"username"`
	Password  string         `gorm:"not null" json:"-"`
	Email     string         `gorm:"unique" json:"email"`
	FullName  string         `json:"fullName"`
	Phone     string         `json:"phone"`
	Role      string         `gorm:"default:'user'" json:"role"` // admin, user
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// UserUpdate структура для обновления пользователя
type UserUpdate struct {
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
}
