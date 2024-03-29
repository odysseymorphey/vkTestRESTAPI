package application

import (
	"context"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/cases"
	"os"
	"os/signal"
	"syscall"

	// "github.com/odysseymorphey/vkTestRESTAPI/internal/cases"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/config"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/server"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/storage/postgres"
	"go.uber.org/zap"
)

type Application struct {
	cancel  context.CancelFunc
	log     *zap.SugaredLogger
	cfg     *config.Config
	storage *postgres.Storage
	server  *server.Server
}

func (a *Application) Build(configPath string) {
	var err error
	a.log = a.initConfig()

	a.cfg, err = config.NewConfig(configPath)
	if err != nil {
		a.log.Fatal("Can't initialize cfg")
	}

	a.storage = a.buildPostgresStorage()

	svc := a.buildService(a.storage)

	a.server = a.buildServer(svc)

}

func (a *Application) Run() {
	a.log.Info("Application started")
	defer a.log.Info("Application stopped")

	var ctx context.Context

	ctx, a.cancel = context.WithCancel(context.Background())
	defer a.cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		select {
		case <-sig:
		case <-ctx.Done():
		}

		a.Stop()
	}()

	a.server.Run(ctx)
}

func (a *Application) Stop() {
	a.storage.Close()
	a.cancel()
	_ = a.log.Sync()
}

func (a *Application) initConfig() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		a.log.Fatal(err)
	}

	return logger.Sugar()
}

func (a *Application) buildPostgresStorage() *postgres.Storage {
	db, err := postgres.NewStorage(a.log, a.cfg.PostgresDSN())
	if err != nil {
		a.log.Fatal(err)
	}

	return db
}

func (a *Application) buildService(storage cases.Storage) *cases.FilmaryService {
	svc, err := cases.NewFilmaryService(a.log, storage)
	if err != nil {
		a.log.Fatal(err)
	}

	return svc
}

func (a *Application) buildServer(svc *cases.FilmaryService) *server.Server {
	srv, err := server.NewServer(a.log, svc, a.cfg.ServerPort())
	if err != nil {
		a.log.Fatal(err)
	}

	return srv
}
