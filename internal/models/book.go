package models

import (
	"gorm.io/gorm"
	"time"
)

// Book представляет книгу в библиотеке
type Book struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Author      string         `gorm:"not null" json:"author"`
	ISBN        string         `gorm:"uniqueIndex" json:"isbn"`
	PublishYear int            `json:"publishYear"`
	Description string         `json:"description"`
	GenreID     uint           `json:"genreId"`
	GenreName   string         `json:"genreName"` // Храним имя жанра вместо объекта
	Quantity    int            `gorm:"default:1" json:"quantity"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// BookCreate структура для создания книги
type BookCreate struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	PublishYear int    `json:"publishYear"`
	Description string `json:"description"`
	GenreID     uint   `json:"genreId"`
	Quantity    int    `json:"quantity"`
}

// BookUpdate структура для обновления книги
type BookUpdate struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	PublishYear int    `json:"publishYear"`
	Description string `json:"description"`
	GenreID     uint   `json:"genreId"`
	Quantity    int    `json:"quantity"`
} 