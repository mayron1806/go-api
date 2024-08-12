package config

import "gorm.io/gorm"

var (
	database *gorm.DB
	logger   *Logger
)

func InitConfig() error {
	logger = GetLogger("Config")
	db, err := initDatabase()

	if err != nil {
		logger.Errorf("postgres connection error: %s", err.Error())
		return err
	}
	database = db
	logger.Debug("postgres connection established")
	logger.Debug("config initialized")
	return nil
}
func GetDatabase() *gorm.DB {
	return database
}
func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}
