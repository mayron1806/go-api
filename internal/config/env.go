package config

import (
	"errors"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Env struct {
	DB_HOST     string `validate:"required"`
	DB_PORT     string `validate:"required"`
	DB_USER     string `validate:"required"`
	DB_PASSWORD string `validate:"required"`
	DB_NAME     string `validate:"required"`
}

func loadEnv() (*Env, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, errors.New("error loading .env file")
	}
	envVars := Env{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
	envValidator := validator.New()
	err = envValidator.Struct(envVars)
	if err != nil {
		return nil, err
	}
	return &envVars, nil
}
