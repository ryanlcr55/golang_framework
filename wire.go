//go:build wireinject
// +build wireinject

package golang_framework

import (
	"github.com/google/wire"
	"go_framework/internal/adapters/gormadapter"
	"go_framework/internal/app"
	"go_framework/internal/app/respositories"
	"go_framework/internal/app/services"
	"go_framework/internal/pkg/configs"
	"go_framework/internal/pkg/database"
)

func InitializeServer() *app.Application {
	wire.Build(
		configs.NewServerConfig,
		database.NewGormDb,
		gormadapter.NewGormPostRepo,
		wire.Bind(new(respositories.IPostRepo), new(*gormadapter.PostRepo)),
		gormadapter.NewTransactionHandler,
		wire.Bind(new(respositories.ITrxHandler), new(*gormadapter.TrxHandler)),
		services.NewPostService,
		app.NewApplication,
	)

	return &app.Application{}
}
