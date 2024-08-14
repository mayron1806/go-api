package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Env struct {
	ENV string

	PORT string

	DB_HOST     string `validate:"required"`
	DB_PORT     string `validate:"required"`
	DB_USER     string `validate:"required"`
	DB_PASSWORD string `validate:"required"`
	DB_NAME     string `validate:"required"`

	SMTP_HOST          string
	SMTP_PORT          string
	SMTP_USER          string
	SMTP_PASS          string
	SMTP_FROM          string
	SHOULD_SEND_EMAILS bool

	JWT_SECRET                 string `validate:"required"`
	JWT_REFRESH_TOKEN_DURATION int
	JWT_ACCESS_TOKEN_DURATION  int
	JWT_ISSUER                 string
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return value
	}
	return defaultValue
}

func loadEnv() (*Env, error) {
	logger := GetLogger("Env")
	logger.Info("loading .env file...")
	err := godotenv.Load()

	if err != nil {
		logger.Errorf("error loading .env file: %s", err.Error())
		return nil, errors.New("error loading .env file")
	}
	envVars := Env{
		ENV:  getEnv("ENV", "development"),
		PORT: getEnv("PORT", "8080"),

		DB_HOST:     getEnv("DB_HOST", ""),
		DB_PORT:     getEnv("DB_PORT", ""),
		DB_USER:     getEnv("DB_USER", ""),
		DB_PASSWORD: getEnv("DB_PASSWORD", ""),
		DB_NAME:     getEnv("DB_NAME", ""),

		SMTP_HOST:          getEnv("SMTP_HOST", ""),
		SMTP_PORT:          getEnv("SMTP_PORT", ""),
		SMTP_USER:          getEnv("SMTP_USER", ""),
		SMTP_PASS:          getEnv("SMTP_PASS", ""),
		SMTP_FROM:          getEnv("SMTP_FROM", ""),
		SHOULD_SEND_EMAILS: os.Getenv("SHOULD_SEND_EMAILS") == "true" || os.Getenv("ENV") == "production",

		JWT_SECRET:                 getEnv("JWT_SECRET", ""),
		JWT_ACCESS_TOKEN_DURATION:  getEnvAsInt("JWT_ACCESS_TOKEN_DURATION", 3600),
		JWT_REFRESH_TOKEN_DURATION: getEnvAsInt("JWT_REFRESH_TOKEN_DURATION", 60*60*24*30),
		JWT_ISSUER:                 getEnv("JWT_ISSUER", ""),
	}

	envValidator := validator.New()
	err = envValidator.Struct(envVars)
	if err != nil {
		logger.Errorf("error validating .env file: %s", err.Error())
		return nil, err
	}
	logger.Info(".env file loaded")
	return &envVars, nil
}
