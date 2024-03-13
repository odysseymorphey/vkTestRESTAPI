package application

import (
	"context"

	"github.com/odysseymorphey/vkTestRESTAPI/internal/config"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/storage/postgres"
	"go.uber.org/zap"
)

type Application struct {
	cancel  context.CancelFunc
	log     *zap.SugaredLogger
	cfg     *config.Config
	storage *postgres.Storage
}

func (a *Application) Build(configPath string) {
	var err error
	a.log = a.initConfig()

	a.cfg, err = config.NewConfig(configPath)
	if err != nil {
		a.log.Fatal("Can't initialize cfg")
	}

	a.storage = a.buildPostgresStorage()

}

func (a *Application) initConfig() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		a.log.Fatal(err)
	}

	return logger.Sugar()
}

func (a *Application) buildPostgresStorage() *postgres.Storage {
	st, err := postgres.NewStorage(a.log, a.cfg.PostgresDSN())
	if err != nil {
		a.log.Fatal()
	}

	return st
}
