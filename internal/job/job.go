package job

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/mayron1806/go-api/config"
)

var logger *config.Logger

func Init() {
	logger = config.GetLogger("Job")
	logger.Info("initializing job...")
	s, err := gocron.NewScheduler()
	if err != nil {
		logger.Errorf("error initializing job: %s", err.Error())
	}
	db := config.GetDatabase()
	s.NewJob(
		gocron.CronJob("0 0 4 * * *", true),
		gocron.NewTask(CleanTokens, logger, db),
	)
	s.Start()

	logger.Info("job initialized")
}
