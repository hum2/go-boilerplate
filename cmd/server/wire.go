//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	infraConfig "github.com/hum2/backend/internal/infrastructure/config"
	"github.com/hum2/backend/internal/infrastructure/db/ent"
	userRepo "github.com/hum2/backend/internal/infrastructure/repository/user"
	interfaceConfig "github.com/hum2/backend/internal/interface/config"
	userController "github.com/hum2/backend/internal/interface/controller/user"
	"github.com/hum2/backend/internal/interface/controller/user/gen"
	userUsecase "github.com/hum2/backend/internal/usecase/user"
)

type App struct {
	Conf           *interfaceConfig.Config
	UserController gen.ServerInterface
}

func NewApp(
	conf *interfaceConfig.Config,
	userController gen.ServerInterface,
) *App {
	return &App{
		Conf:           conf,
		UserController: userController,
	}
}

var WireSet = wire.NewSet(
	interfaceConfig.New,
	infraConfig.New,
	ent.New,
	ent.NewTransaction,
	userRepo.New,
	userUsecase.New,
	userController.New,
	NewApp,
)

func InitializeApp() (*App, error) {
	wire.Build(WireSet)
	return &App{}, nil
}
