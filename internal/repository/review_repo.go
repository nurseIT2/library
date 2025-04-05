package repository

import (
	"gorm.io/gorm"
	"github.com/nurseIT2/library/internal/models"
)

type ReviewRepository interface {
	GetAll() ([]models.Review, error)
	GetById(id uint) (*models.Review, error)
	GetByBookId(bookId uint) ([]models.Review, error)
	GetByUserId(userId uint) ([]models.Review, error)
	Create(review *models.Review) error
	Update(id uint, review *models.ReviewUpdate) error
	Delete(id uint) error
}

type ReviewRepositoryImpl struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepositoryImpl {
	return &ReviewRepositoryImpl{db: db}
}

func (r *ReviewRepositoryImpl) GetAll() ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepositoryImpl) GetById(id uint) (*models.Review, error) {
	var review models.Review
	err := r.db.First(&review, id).Error
	return &review, err
}

func (r *ReviewRepositoryImpl) GetByBookId(bookId uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("book_id = ?", bookId).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepositoryImpl) GetByUserId(userId uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("user_id = ?", userId).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepositoryImpl) Create(review *models.Review) error {
	var user models.User
	if err := r.db.First(&user, review.UserID).Error; err == nil {
		review.UserName = user.Username
	}
	return r.db.Create(review).Error
}

func (r *ReviewRepositoryImpl) Update(id uint, review *models.ReviewUpdate) error {
	return r.db.Model(&models.Review{}).Where("id = ?", id).Updates(review).Error
}

func (r *ReviewRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
} 