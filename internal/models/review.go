package models

import (
	"gorm.io/gorm"
	"time"
)

// Review представляет отзыв о книге
type Review struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	BookID    uint           `json:"bookId"`
	UserID    uint           `json:"userId"`
	Rating    int            `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment   string         `json:"comment"`
	UserName  string         `json:"userName"` // Вместо полного объекта User
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// ReviewCreate структура для создания отзыва
type ReviewCreate struct {
	BookID  uint   `json:"bookId" binding:"required"`
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}

// ReviewUpdate структура для обновления отзыва
type ReviewUpdate struct {
	Rating  int    `json:"rating" binding:"min=1,max=5"`
	Comment string `json:"comment"`
} 