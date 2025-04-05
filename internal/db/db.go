package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nurseIT2/library/internal/models"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbName := "mydatabase"
	dbUser := "myuser"
	dbPass := "mypassword"
	dbPort := "5444"
	sslmode := "disable"
	
	// Используем параметрический DSN для GORM
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
		dbHost, dbUser, dbPass, dbName, dbPort, sslmode)
		
	// Подключение для миграции и начальной настройки
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)
	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Выполняем миграции из файлов, если они есть
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Printf("Миграция: %v", err)
	}

	// Настраиваем GORM с более подробным логированием
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	
	// Создаем подключение GORM, используя DSN вместо соединения
	gormDB, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal(err)
	}
	DB = gormDB

	// Автомиграция для создания таблиц библиотеки
	err = DB.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Genre{},
		&models.Review{},
		&models.Borrow{},
	)
	if err != nil {
		log.Fatal("Ошибка автомиграции:", err)
	}
}
