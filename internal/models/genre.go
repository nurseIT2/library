package models

import "time"

// Genre представляет жанр книги
type Genre struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GenreCreate структура для создания жанра
type GenreCreate struct {
	Name string `json:"name" binding:"required"`
}

// GenreUpdate структура для обновления жанра
type GenreUpdate struct {
	Name string `json:"name" binding:"required"`
} 