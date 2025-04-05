package repository

import (
	"gorm.io/gorm"
	"github.com/nurseIT2/library/internal/models"
	"time"
)

type BorrowRepository interface {
	GetAll() ([]models.Borrow, error)
	GetById(id uint) (*models.Borrow, error)
	GetByUserId(userId uint) ([]models.Borrow, error)
	GetByBookId(bookId uint) ([]models.Borrow, error)
	GetBorrowed() ([]models.Borrow, error)
	GetOverdue() ([]models.Borrow, error)
	Create(borrow *models.Borrow) error
	Update(id uint, borrow *models.BorrowUpdate) error
	ReturnBook(id uint) error
	Delete(id uint) error
}

type BorrowRepositoryImpl struct {
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) *BorrowRepositoryImpl {
	return &BorrowRepositoryImpl{db: db}
}

func (r *BorrowRepositoryImpl) GetAll() ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Find(&borrows).Error
	return borrows, err
}

func (r *BorrowRepositoryImpl) GetById(id uint) (*models.Borrow, error) {
	var borrow models.Borrow
	err := r.db.First(&borrow, id).Error
	return &borrow, err
}

func (r *BorrowRepositoryImpl) GetByUserId(userId uint) ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Where("user_id = ?", userId).Find(&borrows).Error
	return borrows, err
}

func (r *BorrowRepositoryImpl) GetByBookId(bookId uint) ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Where("book_id = ?", bookId).Find(&borrows).Error
	return borrows, err
}

func (r *BorrowRepositoryImpl) GetBorrowed() ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Where("status = ?", "borrowed").Find(&borrows).Error
	return borrows, err
}

func (r *BorrowRepositoryImpl) GetOverdue() ([]models.Borrow, error) {
	var borrows []models.Borrow
	now := time.Now()
	err := r.db.Where("status = ? AND return_date < ?", "borrowed", now).Find(&borrows).Error
	return borrows, err
}

func (r *BorrowRepositoryImpl) Create(borrow *models.Borrow) error {
	// Получаем информацию о пользователе и книге
	var user models.User
	var book models.Book
	
	if err := r.db.First(&user, borrow.UserID).Error; err == nil {
		borrow.UserName = user.Username
	}
	
	if err := r.db.First(&book, borrow.BookID).Error; err == nil {
		borrow.BookTitle = book.Title
	}
	
	return r.db.Create(borrow).Error
}

func (r *BorrowRepositoryImpl) Update(id uint, borrow *models.BorrowUpdate) error {
	return r.db.Model(&models.Borrow{}).Where("id = ?", id).Updates(borrow).Error
}

func (r *BorrowRepositoryImpl) ReturnBook(id uint) error {
	now := time.Now()
	return r.db.Model(&models.Borrow{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"returned_at": now,
			"status":      "returned",
		}).Error
}

func (r *BorrowRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Borrow{}, id).Error
} 