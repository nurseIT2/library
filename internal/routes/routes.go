package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nurseIT2/library/internal/auth"
	"github.com/nurseIT2/library/internal/db"
	"github.com/nurseIT2/library/internal/delivery"
	"github.com/nurseIT2/library/internal/middleware"
	"github.com/nurseIT2/library/internal/repository"
	"github.com/nurseIT2/library/internal/services"
)

func SetupRoutes(r *gin.Engine) {
	// Репозитории
	bookRepo := repository.NewBookRepository(db.DB)
	genreRepo := repository.NewGenreRepository(db.DB)
	reviewRepo := repository.NewReviewRepository(db.DB)
	borrowRepo := repository.NewBorrowRepository(db.DB)
	studentRepo := repository.NewStudentRepository(db.DB)

	// Сервисы
	bookService := service.NewBookService(bookRepo)
	genreService := service.NewGenreService(genreRepo)
	reviewService := service.NewReviewService(reviewRepo)
	borrowService := service.NewBorrowService(borrowRepo, bookRepo)
	studentService := service.NewStudentService(studentRepo)

	// Контроллеры
	bookHandler := delivery.NewBookHandler(bookService)
	genreHandler := delivery.NewGenreHandler(genreService)
	reviewHandler := delivery.NewReviewHandler(reviewService)
	borrowHandler := delivery.NewBorrowHandler(borrowService)
	studentHandler := delivery.NewStudentHandler(studentService)

	// Роуты для аутентификации
	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	// Роуты для управления студентами
	studentRoutes := r.Group("api/v1/students")
	{
		studentRoutes.GET("", studentHandler.GetAllStudents)
		studentRoutes.GET("/:id", studentHandler.GetStudent)
		studentRoutes.POST("", studentHandler.CreateStudent)
		studentRoutes.PUT("/:id", studentHandler.UpdateStudent)
		studentRoutes.DELETE("/:id", studentHandler.DeleteStudent)
	}

	// Библиотечные API роуты
	api := r.Group("api/v1")

	// Роуты для книг (общедоступные)
	booksRoutes := api.Group("/books")
	{
		booksRoutes.GET("", bookHandler.GetAllBooks)
		booksRoutes.GET("/:id", bookHandler.GetBook)
		booksRoutes.GET("/search", bookHandler.SearchBooks)
		booksRoutes.GET("/genre/:genreId", bookHandler.GetBooksByGenre)
		
		// Защищенные роуты для книг - требуют авторизации
		booksAuth := booksRoutes.Group("")
		booksAuth.Use(middleware.AuthMiddleware())
		{
			booksAuth.POST("", bookHandler.CreateBook)
			booksAuth.PUT("/:id", bookHandler.UpdateBook)
			booksAuth.DELETE("/:id", bookHandler.DeleteBook)
		}
	}

	// Роуты для жанров
	genresRoutes := api.Group("/genres")
	{
		genresRoutes.GET("", genreHandler.GetAllGenres)
		genresRoutes.GET("/:id", genreHandler.GetGenre)
		
		// Защищенные роуты для жанров
		genresAuth := genresRoutes.Group("")
		genresAuth.Use(middleware.AuthMiddleware())
		{
			genresAuth.POST("", genreHandler.CreateGenre)
			genresAuth.PUT("/:id", genreHandler.UpdateGenre)
			genresAuth.DELETE("/:id", genreHandler.DeleteGenre)
		}
	}

	// Роуты для отзывов
	reviewsRoutes := api.Group("/reviews")
	{
		reviewsRoutes.GET("", reviewHandler.GetAllReviews)
		reviewsRoutes.GET("/:id", reviewHandler.GetReview)
		reviewsRoutes.GET("/book/:bookId", reviewHandler.GetBookReviews)
		
		// Защищенные роуты для отзывов
		reviewsAuth := reviewsRoutes.Group("")
		reviewsAuth.Use(middleware.AuthMiddleware())
		{
			reviewsAuth.POST("", reviewHandler.CreateReview)
			reviewsAuth.PUT("/:id", reviewHandler.UpdateReview)
			reviewsAuth.DELETE("/:id", reviewHandler.DeleteReview)
		}
	}

	// Роуты для выдачи книг (все защищены)
	borrowsRoutes := api.Group("/borrows")
	borrowsRoutes.Use(middleware.AuthMiddleware())
	{
		borrowsRoutes.GET("", borrowHandler.GetAllBorrows)
		borrowsRoutes.GET("/:id", borrowHandler.GetBorrow)
		borrowsRoutes.GET("/user/:userId", borrowHandler.GetUserBorrows)
		borrowsRoutes.GET("/current", borrowHandler.GetCurrentBorrows)
		borrowsRoutes.GET("/overdue", borrowHandler.GetOverdueBorrows)
		borrowsRoutes.POST("", borrowHandler.CreateBorrow)
		borrowsRoutes.PUT("/:id", borrowHandler.UpdateBorrow)
		borrowsRoutes.POST("/:id/return", borrowHandler.ReturnBook)
		borrowsRoutes.DELETE("/:id", borrowHandler.DeleteBorrow)
	}
}
