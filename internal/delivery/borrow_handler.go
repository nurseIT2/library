package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/nurseIT2/library/internal/models"
	"github.com/nurseIT2/library/internal/services"
	"strconv"
)

type BorrowHandler struct {
	service *service.BorrowService
}

func NewBorrowHandler(service *service.BorrowService) *BorrowHandler {
	return &BorrowHandler{service: service}
}

// GetAllBorrows получение списка всех выдач книг
func (h *BorrowHandler) GetAllBorrows(c *gin.Context) {
	borrows, err := h.service.GetAllBorrows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список выдач"})
		return
	}

	c.JSON(http.StatusOK, borrows)
}

// GetBorrow получение выдачи по ID
func (h *BorrowHandler) GetBorrow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID выдачи"})
		return
	}

	borrow, err := h.service.GetBorrowByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись о выдаче не найдена"})
		return
	}

	c.JSON(http.StatusOK, borrow)
}

// GetUserBorrows получение выдач пользователя
func (h *BorrowHandler) GetUserBorrows(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	borrows, err := h.service.GetBorrowsByUserID(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список выдач"})
		return
	}

	c.JSON(http.StatusOK, borrows)
}

// GetCurrentBorrows получение текущих выдач
func (h *BorrowHandler) GetCurrentBorrows(c *gin.Context) {
	borrows, err := h.service.GetBorrowed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список выдач"})
		return
	}

	c.JSON(http.StatusOK, borrows)
}

// GetOverdueBorrows получение просроченных выдач
func (h *BorrowHandler) GetOverdueBorrows(c *gin.Context) {
	borrows, err := h.service.GetOverdue()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список просроченных выдач"})
		return
	}

	c.JSON(http.StatusOK, borrows)
}

// CreateBorrow создание новой выдачи
func (h *BorrowHandler) CreateBorrow(c *gin.Context) {
	var borrowCreate models.BorrowCreate
	if err := c.ShouldBindJSON(&borrowCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	borrow, err := h.service.CreateBorrow(borrowCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать запись о выдаче: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, borrow)
}

// ReturnBook возврат книги
func (h *BorrowHandler) ReturnBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID выдачи"})
		return
	}

	borrow, err := h.service.ReturnBook(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось вернуть книгу: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrow)
}

// UpdateBorrow обновление информации о выдаче
func (h *BorrowHandler) UpdateBorrow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID выдачи"})
		return
	}

	var borrowUpdate models.BorrowUpdate
	if err := c.ShouldBindJSON(&borrowUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	borrow, err := h.service.UpdateBorrow(uint(id), &borrowUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись о выдаче не найдена или не удалось обновить"})
		return
	}

	c.JSON(http.StatusOK, borrow)
}

// DeleteBorrow удаление записи о выдаче
func (h *BorrowHandler) DeleteBorrow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID выдачи"})
		return
	}

	if err := h.service.DeleteBorrow(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись о выдаче не найдена или не удалось удалить"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Запись о выдаче успешно удалена"})
} 