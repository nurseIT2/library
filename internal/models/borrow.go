package models

import (
	"gorm.io/gorm"
	"time"
)

// Borrow представляет запись о выдаче книги пользователю
type Borrow struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `json:"userId"`
	BookID     uint           `json:"bookId"`
	UserName   string         `json:"userName"` // Имя пользователя вместо объекта
	BookTitle  string         `json:"bookTitle"` // Название книги вместо объекта
	BorrowDate time.Time      `json:"borrowDate"`
	ReturnDate time.Time      `json:"returnDate"` // Планируемая дата возврата
	ReturnedAt *time.Time     `json:"returnedAt"` // Фактическая дата возврата (nil если не возвращена)
	Status     string         `gorm:"default:'borrowed'" json:"status"` // borrowed, returned, overdue
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// BorrowCreate структура для создания записи о выдаче книги
type BorrowCreate struct {
	UserID     uint      `json:"userId" binding:"required"`
	BookID     uint      `json:"bookId" binding:"required"`
	BorrowDate time.Time `json:"borrowDate"`
	ReturnDate time.Time `json:"returnDate" binding:"required"`
}

// BorrowUpdate структура для обновления записи о выдаче книги
type BorrowUpdate struct {
	ReturnDate time.Time  `json:"returnDate"`
	ReturnedAt *time.Time `json:"returnedAt"`
	Status     string     `json:"status"`
} 