package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/nurseIT2/library/internal/models"
	"github.com/nurseIT2/library/internal/services"
	"strconv"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// GetAllBooks получение списка всех книг
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список книг"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBook получение книги по ID
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID книги"})
		return
	}

	book, err := h.service.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Книга не найдена"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook создание новой книги
func (h *BookHandler) CreateBook(c *gin.Context) {
	var bookCreate models.BookCreate
	if err := c.ShouldBindJSON(&bookCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	book, err := h.service.CreateBook(bookCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать книгу"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// UpdateBook обновление информации о книге
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID книги"})
		return
	}

	var bookUpdate models.BookUpdate
	if err := c.ShouldBindJSON(&bookUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	book, err := h.service.UpdateBook(uint(id), &bookUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Книга не найдена или не удалось обновить"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook удаление книги
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID книги"})
		return
	}

	if err := h.service.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Книга не найдена или не удалось удалить"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Книга успешно удалена"})
}

// SearchBooks поиск книг по названию или автору
func (h *BookHandler) SearchBooks(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указан поисковый запрос"})
		return
	}

	books, err := h.service.SearchBooks(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось выполнить поиск"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBooksByGenre получение книг по жанру
func (h *BookHandler) GetBooksByGenre(c *gin.Context) {
	genreId, err := strconv.ParseUint(c.Param("genreId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID жанра"})
		return
	}

	books, err := h.service.GetBooksByGenre(uint(genreId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить книги по жанру"})
		return
	}

	c.JSON(http.StatusOK, books)
} 