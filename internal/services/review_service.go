package service

import (
	"github.com/nurseIT2/library/internal/models"
	"github.com/nurseIT2/library/internal/repository"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) GetAllReviews() ([]models.Review, error) {
	return s.repo.GetAll()
}

func (s *ReviewService) GetReviewByID(id uint) (*models.Review, error) {
	return s.repo.GetById(id)
}

func (s *ReviewService) GetReviewsByBookID(bookId uint) ([]models.Review, error) {
	return s.repo.GetByBookId(bookId)
}

func (s *ReviewService) GetReviewsByUserID(userId uint) ([]models.Review, error) {
	return s.repo.GetByUserId(userId)
}

func (s *ReviewService) CreateReview(userId uint, reviewCreate models.ReviewCreate) (*models.Review, error) {
	review := models.Review{
		BookID:  reviewCreate.BookID,
		UserID:  userId,
		Rating:  reviewCreate.Rating,
		Comment: reviewCreate.Comment,
	}

	err := s.repo.Create(&review)
	if err != nil {
		return nil, err
	}

	return &review, nil
}

func (s *ReviewService) UpdateReview(id uint, reviewUpdate *models.ReviewUpdate) (*models.Review, error) {
	err := s.repo.Update(id, reviewUpdate)
	if err != nil {
		return nil, err
	}

	return s.repo.GetById(id)
}

func (s *ReviewService) DeleteReview(id uint) error {
	return s.repo.Delete(id)
} 