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

	logger.Debug("config loading...")
	env, err = loadEnv()
	if err != nil {
		logger.Errorf("env config error: %s", err.Error())
		return err
	}
	logger.Debug("config loaded")

	logger.Debug("database loading")
	database, err = initDatabase()
	if err != nil {
		logger.Errorf("postgres connection error: %s", err.Error())
		return err
	}
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
func GetEnv() *Env {
	return env
}
