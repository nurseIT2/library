package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/nurseIT2/library/internal/models"
	"github.com/nurseIT2/library/internal/services"
	"strconv"
)

type GenreHandler struct {
	service *service.GenreService
}

func NewGenreHandler(service *service.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

// GetAllGenres получение списка всех жанров
func (h *GenreHandler) GetAllGenres(c *gin.Context) {
	genres, err := h.service.GetAllGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список жанров"})
		return
	}

	c.JSON(http.StatusOK, genres)
}

// GetGenre получение жанра по ID
func (h *GenreHandler) GetGenre(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID жанра"})
		return
	}

	genre, err := h.service.GetGenreByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Жанр не найден"})
		return
	}

	c.JSON(http.StatusOK, genre)
}

// CreateGenre создание нового жанра
func (h *GenreHandler) CreateGenre(c *gin.Context) {
	var genreCreate models.GenreCreate
	if err := c.ShouldBindJSON(&genreCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	genre, err := h.service.CreateGenre(genreCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать жанр"})
		return
	}

	c.JSON(http.StatusCreated, genre)
}

// UpdateGenre обновление информации о жанре
func (h *GenreHandler) UpdateGenre(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID жанра"})
		return
	}

	var genreUpdate models.GenreUpdate
	if err := c.ShouldBindJSON(&genreUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	genre, err := h.service.UpdateGenre(uint(id), &genreUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Жанр не найден или не удалось обновить"})
		return
	}

	c.JSON(http.StatusOK, genre)
}

// DeleteGenre удаление жанра
func (h *GenreHandler) DeleteGenre(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID жанра"})
		return
	}

	if err := h.service.DeleteGenre(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Жанр не найден или не удалось удалить"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Жанр успешно удален"})
}