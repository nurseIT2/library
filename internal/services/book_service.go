package service

import (
	"github.com/nurseIT2/library/internal/models"
	"github.com/nurseIT2/library/internal/repository"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetBookByID(id uint) (*models.Book, error) {
	return s.repo.GetById(id)
}

func (s *BookService) CreateBook(bookCreate models.BookCreate) (*models.Book, error) {
	book := models.Book{
		Title:       bookCreate.Title,
		Author:      bookCreate.Author,
		ISBN:        bookCreate.ISBN,
		PublishYear: bookCreate.PublishYear,
		Description: bookCreate.Description,
		GenreID:     bookCreate.GenreID,
		Quantity:    bookCreate.Quantity,
	}

	err := s.repo.Create(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookService) UpdateBook(id uint, bookUpdate *models.BookUpdate) (*models.Book, error) {
	err := s.repo.Update(id, bookUpdate)
	if err != nil {
		return nil, err
	}

	return s.repo.GetById(id)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}

func (s *BookService) SearchBooks(query string) ([]models.Book, error) {
	return s.repo.Search(query)
}

func (s *BookService) GetBooksByGenre(genreId uint) ([]models.Book, error) {
	return s.repo.GetByGenre(genreId)
} 