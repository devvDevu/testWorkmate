package app

import (
	"testWorkmate/cmd/app/adapters"
	"testWorkmate/cmd/app/handlers"
	"testWorkmate/cmd/app/repositories"
	"testWorkmate/cmd/app/services"
	"testWorkmate/cmd/app/usecases"
	"testWorkmate/internal/config"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg          *config.Config
	adapters     *adapters.Adapters
	usecases     *usecases.Usecases
	handlers     *handlers.Handlers
	services     *services.Services
	repositories *repositories.Repositories
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg:          cfg,
		adapters:     adapters.NewAdapters(),
		usecases:     usecases.NewUsecases(),
		handlers:     handlers.NewHandlers(),
		services:     services.NewServices(),
		repositories: repositories.NewRepositories(),
	}
}

func (a *App) GetAdapters() *adapters.Adapters {
	return a.adapters
}

func (a *App) MustInit(router *mux.Router) *App {
	const action = "App MustInit "

	{
		logrus.Info(action, "adapters starting")
		a.adapters.MustInit()
		logrus.Info(action, "adapters running")
	}

	{
		logrus.Info(action, "repositories starting")
		a.repositories.MustInit(a.adapters)
		logrus.Info(action, "repositories running")
	}

	{
		logrus.Info(action, "services starting")
		a.services.MustInit(a.repositories)
		logrus.Info(action, "services running")
	}

	{
		logrus.Info(action, "usecases starting")
		a.usecases.MustInit(a.services)
		logrus.Info(action, "usecases running")
	}

	{
		logrus.Info(action, "handlers starting")
		a.handlers.MustInit(a.usecases, router)
		logrus.Info(action, "handlers running")
	}

	logrus.Info(action, "done")

	return a
}
