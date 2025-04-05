package service

import (
	"github.com/nurseIT2/library/internal/models"
	"github.com/nurseIT2/library/internal/repository"
)

type GenreService struct {
	repo repository.GenreRepository
}

func NewGenreService(repo repository.GenreRepository) *GenreService {
	return &GenreService{repo: repo}
}

func (s *GenreService) GetAllGenres() ([]models.Genre, error) {
	return s.repo.GetAll()
}

func (s *GenreService) GetGenreByID(id uint) (*models.Genre, error) {
	return s.repo.GetById(id)
}

func (s *GenreService) CreateGenre(genreCreate models.GenreCreate) (*models.Genre, error) {
	genre := models.Genre{
		Name: genreCreate.Name,
	}

	err := s.repo.Create(&genre)
	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (s *GenreService) UpdateGenre(id uint, genreUpdate *models.GenreUpdate) (*models.Genre, error) {
	err := s.repo.Update(id, genreUpdate)
	if err != nil {
		return nil, err
	}

	return s.repo.GetById(id)
}

func (s *GenreService) DeleteGenre(id uint) error {
	return s.repo.Delete(id)
} 