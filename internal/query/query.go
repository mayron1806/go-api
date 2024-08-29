package query

import (
	"github.com/mayron1806/go-api/config"
	"gorm.io/gorm"
)

type Query struct {
	db     *gorm.DB
	logger *config.Logger
}
