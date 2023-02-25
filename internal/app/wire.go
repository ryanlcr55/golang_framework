//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"go_framework/internal/adapters/gormadapter"
	"go_framework/internal/app/respositories"
	"go_framework/internal/app/services"
	"go_framework/internal/configs"
	"go_framework/internal/pkg/database"
)

func InitializeServer() *Application {
	wire.Build(
		configs.NewServerConfig,
		database.NewGormDb,
		gormadapter.NewGormPostRepo,
		wire.Bind(new(respositories.IPostRepo), new(*gormadapter.PostRepo)),
		gormadapter.NewTransactionHandler,
		wire.Bind(new(respositories.ITrxHandler), new(*gormadapter.TrxHandler)),
		services.NewPostService,
		NewApplication,
	)

	return &Application{}
}
