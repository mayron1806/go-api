package main

import (
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/job"
	"github.com/mayron1806/go-api/internal/router"
)

func main() {
	logger := config.GetLogger("main")
	err := config.InitConfig()
	if err != nil {
		logger.Errorf("config error: %s", err.Error())
		panic(err)
	}

	job.Init()

	router.InitRouter()
}
