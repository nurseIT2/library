package repository

import (
	"gorm.io/gorm"
	"github.com/nurseIT2/library/internal/models"
)

type GenreRepository interface {
	GetAll() ([]models.Genre, error)
	GetById(id uint) (*models.Genre, error)
	Create(genre *models.Genre) error
	Update(id uint, genre *models.GenreUpdate) error
	Delete(id uint) error
}

type GenreRepositoryImpl struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *GenreRepositoryImpl {
	return &GenreRepositoryImpl{db: db}
}

func (r *GenreRepositoryImpl) GetAll() ([]models.Genre, error) {
	var genres []models.Genre
	err := r.db.Find(&genres).Error
	return genres, err
}

func (r *GenreRepositoryImpl) GetById(id uint) (*models.Genre, error) {
	var genre models.Genre
	err := r.db.First(&genre, id).Error
	return &genre, err
}

func (r *GenreRepositoryImpl) Create(genre *models.Genre) error {
	return r.db.Create(genre).Error
}

func (r *GenreRepositoryImpl) Update(id uint, genre *models.GenreUpdate) error {
	return r.db.Model(&models.Genre{}).Where("id = ?", id).Updates(genre).Error
}

func (r *GenreRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Genre{}, id).Error
} 