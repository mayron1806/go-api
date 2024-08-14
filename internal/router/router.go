package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/config"
)

var (
	logger *config.Logger
)

func InitRouter() error {
	logger = config.GetLogger("Router")
	env := config.GetEnv()
	if env.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	logger.Info("registering routes...")
	err := registerRoutes(router)
	if err != nil {
		logger.Errorf("error registering routes: %s", err.Error())
		return err
	}
	logger.Info("routes registered")
	logger.Info("starting server...")
	router.Run(fmt.Sprintf(":%s", env.PORT))
	return nil
}
