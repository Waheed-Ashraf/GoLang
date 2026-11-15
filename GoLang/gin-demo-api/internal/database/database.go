package database

import (
	"fmt"
	"my-app/internal/models"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})

	if err != nil {
		return nil, fmt.Errorf("Database Open Connection Error : %w", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Access Connection Error : %w", err)
	}

	sqlDb.SetMaxOpenConns(25)
	sqlDb.SetMaxIdleConns(20)
	sqlDb.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDb.Ping(); err != nil {
		return nil, fmt.Errorf("Ping Database Error : %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Book{})
}