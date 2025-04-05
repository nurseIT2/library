package repository

import (
	"gorm.io/gorm"
	"github.com/nurseIT2/library/internal/models"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id uint) (*models.Book, error)
	Create(book *models.Book) error
	Update(id uint, book *models.BookUpdate) error
	Delete(id uint) error
	Search(query string) ([]models.Book, error)
	GetByGenre(genreId uint) ([]models.Book, error)
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db: db}
}

func (r *BookRepositoryImpl) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *BookRepositoryImpl) GetById(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *BookRepositoryImpl) Create(book *models.Book) error {
	if book.GenreID > 0 {
		var genre models.Genre
		if err := r.db.First(&genre, book.GenreID).Error; err == nil {
			book.GenreName = genre.Name
		}
	}
	return r.db.Create(book).Error
}

func (r *BookRepositoryImpl) Update(id uint, bookUpdate *models.BookUpdate) error {
	if bookUpdate.GenreID > 0 {
		var genre models.Genre
		if err := r.db.First(&genre, bookUpdate.GenreID).Error; err == nil {
			r.db.Model(&models.Book{}).Where("id = ?", id).Update("genre_name", genre.Name)
		}
	}
	return r.db.Model(&models.Book{}).Where("id = ?", id).Updates(bookUpdate).Error
}

func (r *BookRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}

func (r *BookRepositoryImpl) Search(query string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("title ILIKE ? OR author ILIKE ?", "%"+query+"%", "%"+query+"%").
		Find(&books).Error
	return books, err
}

func (r *BookRepositoryImpl) GetByGenre(genreId uint) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("genre_id = ?", genreId).Find(&books).Error
	return books, err
} 