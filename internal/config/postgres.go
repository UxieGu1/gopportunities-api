package config

import (
	"fmt"
	"os"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbName, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(
		&schemas.User{},
		&schemas.Company{},
		&schemas.Candidate{},
		&schemas.Opening{},
		&schemas.Application{},
	)
	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}

	return db, nil
}