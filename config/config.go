package config

import "gorm.io/gorm"

var (
	database *gorm.DB
	logger   *Logger
	env      *Env
)

func InitConfig() error {
	var err error
	logger = GetLogger("Config")
	logger.Info("Config Initializing...")

	env, err = loadEnv()
	if err != nil {
		return err
	}
	database, err = initDatabase()
	if err != nil {
		return err
	}
	return nil
}
func GetDatabase() *gorm.DB {
	return database
}
func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}
func GetEnv() *Env {
	return env
}
