package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/nurseIT2/library/internal/models"
	"github.com/nurseIT2/library/internal/services"
	"strconv"
)

type ReviewHandler struct {
	service *service.ReviewService
}

func NewReviewHandler(service *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

// GetAllReviews получение списка всех отзывов
func (h *ReviewHandler) GetAllReviews(c *gin.Context) {
	reviews, err := h.service.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список отзывов"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// GetReview получение отзыва по ID
func (h *ReviewHandler) GetReview(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID отзыва"})
		return
	}

	review, err := h.service.GetReviewByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Отзыв не найден"})
		return
	}

	c.JSON(http.StatusOK, review)
}

// GetBookReviews получение отзывов по ID книги
func (h *ReviewHandler) GetBookReviews(c *gin.Context) {
	bookId, err := strconv.ParseUint(c.Param("bookId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID книги"})
		return
	}

	reviews, err := h.service.GetReviewsByBookID(uint(bookId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить отзывы"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// CreateReview создание нового отзыва
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	// Получаем ID пользователя из контекста (предполагается, что аутентификация уже пройдена)
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Для добавления отзыва необходимо авторизоваться"})
		return
	}

	var reviewCreate models.ReviewCreate
	if err := c.ShouldBindJSON(&reviewCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	review, err := h.service.CreateReview(userId.(uint), reviewCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать отзыв"})
		return
	}

	c.JSON(http.StatusCreated, review)
}

// UpdateReview обновление отзыва
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID отзыва"})
		return
	}

	var reviewUpdate models.ReviewUpdate
	if err := c.ShouldBindJSON(&reviewUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	review, err := h.service.UpdateReview(uint(id), &reviewUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Отзыв не найден или не удалось обновить"})
		return
	}

	c.JSON(http.StatusOK, review)
}

// DeleteReview удаление отзыва
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID отзыва"})
		return
	}

	if err := h.service.DeleteReview(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Отзыв не найден или не удалось удалить"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Отзыв успешно удален"})
} 