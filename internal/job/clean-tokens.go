package job

import (
	"time"

	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/model"
	"gorm.io/gorm"
)

func CleanTokens(logger *config.Logger, db *gorm.DB) {
	logger.Info("===========START CLEANING EXPIRED TOKENS==============")
	tx := db.Table("tokens").Where("expires_at < ?", time.Now()).Delete(&model.Token{})
	logger.Info("deleted", tx.RowsAffected, "expired tokens")
	logger.Info("===========FINISH CLEANING EXPIRED TOKENS=============")
}
