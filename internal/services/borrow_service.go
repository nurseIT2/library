package service

import (
	"errors"
	"github.com/nurseIT2/library/internal/models"
	"github.com/nurseIT2/library/internal/repository"
	"time"
)

type BorrowService struct {
	repo      repository.BorrowRepository
	bookRepo  repository.BookRepository
}

func NewBorrowService(repo repository.BorrowRepository, bookRepo repository.BookRepository) *BorrowService {
	return &BorrowService{
		repo:      repo,
		bookRepo:  bookRepo,
	}
}

func (s *BorrowService) GetAllBorrows() ([]models.Borrow, error) {
	return s.repo.GetAll()
}

func (s *BorrowService) GetBorrowByID(id uint) (*models.Borrow, error) {
	return s.repo.GetById(id)
}

func (s *BorrowService) GetBorrowsByUserID(userId uint) ([]models.Borrow, error) {
	return s.repo.GetByUserId(userId)
}

func (s *BorrowService) GetBorrowsByBookID(bookId uint) ([]models.Borrow, error) {
	return s.repo.GetByBookId(bookId)
}

func (s *BorrowService) GetBorrowed() ([]models.Borrow, error) {
	return s.repo.GetBorrowed()
}

func (s *BorrowService) GetOverdue() ([]models.Borrow, error) {
	return s.repo.GetOverdue()
}

func (s *BorrowService) CreateBorrow(borrowCreate models.BorrowCreate) (*models.Borrow, error) {
	// Проверяем, есть ли книга в наличии
	book, err := s.bookRepo.GetById(borrowCreate.BookID)
	if err != nil {
		return nil, err
	}

	if book.Quantity <= 0 {
		return nil, errors.New("книга не доступна для выдачи")
	}

	// Создаем запись о выдаче
	borrowDate := borrowCreate.BorrowDate
	if borrowDate.IsZero() {
		borrowDate = time.Now()
	}

	borrow := models.Borrow{
		UserID:     borrowCreate.UserID,
		BookID:     borrowCreate.BookID,
		BorrowDate: borrowDate,
		ReturnDate: borrowCreate.ReturnDate,
		Status:     "borrowed",
	}

	err = s.repo.Create(&borrow)
	if err != nil {
		return nil, err
	}

	// Уменьшаем количество доступных книг
	book.Quantity--
	s.bookRepo.Update(book.ID, &models.BookUpdate{Quantity: book.Quantity})

	return &borrow, nil
}

func (s *BorrowService) UpdateBorrow(id uint, borrowUpdate *models.BorrowUpdate) (*models.Borrow, error) {
	err := s.repo.Update(id, borrowUpdate)
	if err != nil {
		return nil, err
	}

	return s.repo.GetById(id)
}

func (s *BorrowService) ReturnBook(id uint) (*models.Borrow, error) {
	// Получаем запись о выдаче
	borrow, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	// Если книга уже возвращена, возвращаем ошибку
	if borrow.Status == "returned" {
		return nil, errors.New("книга уже возвращена")
	}

	// Обновляем статус
	err = s.repo.ReturnBook(id)
	if err != nil {
		return nil, err
	}

	// Увеличиваем количество доступных книг
	book, err := s.bookRepo.GetById(borrow.BookID)
	if err != nil {
		return nil, err
	}

	book.Quantity++
	s.bookRepo.Update(book.ID, &models.BookUpdate{Quantity: book.Quantity})

	return s.repo.GetById(id)
}

func (s *BorrowService) DeleteBorrow(id uint) error {
	return s.repo.Delete(id)
} 