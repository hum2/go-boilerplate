//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/hum2/backend/internal/interface/controller/batch"
	ctr "github.com/hum2/backend/internal/interface/controller/batch/hellowire"
	usecase "github.com/hum2/backend/internal/usecase/batch/hellowire"
)

type App struct {
	Ctr batch.LambdaController
}

func NewApp(
	ctr batch.LambdaController,
) *App {
	return &App{
		Ctr: ctr,
	}
}

var WireSet = wire.NewSet(
	usecase.New,
	ctr.New,
	NewApp,
)

func InitializeApp() (*App, error) {
	wire.Build(WireSet)
	return &App{}, nil
}
