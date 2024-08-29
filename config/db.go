package config

import (
	"fmt"

	"github.com/mayron1806/go-api/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() (*gorm.DB, error) {
	logger := GetLogger("Database")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_NAME, env.DB_PASSWORD)
	logger.Info("Connecting to database...")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		logger.Errorf("postgres connection error: %s", err.Error())
		return db, err
	}
	logger.Info("Connection established")

	logger.Info("Migrating database...")
	err = db.AutoMigrate(
		&model.User{},
		&model.Organization{},
		&model.Member{},
		&model.Token{},
		&model.SocialProvider{},
	)
	if err != nil {
		logger.Errorf("postgres migration error: %s", err.Error())
		return db, err
	}
	logger.Info("Database migrated")

	return db, nil
}
