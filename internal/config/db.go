package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() (*gorm.DB, error) {
	logger := GetLogger("InitDatabase")

	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=go-api password=postgres sslmode=disable"), &gorm.Config{})

	if err != nil {
		logger.Errorf("postgres connection error: %s", err.Error())
		return db, err
	}
	err = db.AutoMigrate()

	if err != nil {
		logger.Errorf("postgres migration error: %s", err.Error())
		return db, err
	}
	return db, nil
}
