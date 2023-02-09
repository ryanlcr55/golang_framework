//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"go_framework/internal/adapters/gormRepo"
	"go_framework/internal/app/respositories"
	"go_framework/internal/app/services"
	"go_framework/internal/configs"
	"go_framework/internal/pkg/database"
)

func InitializeServer() Application {
	wire.Build(
		configs.NewServerConfig,
		database.NewGormDb,
		gormRepo.NewGormPostRepo,
		wire.Bind(new(respositories.IPostRepo), new(gormRepo.PostRepo)),
		services.NewPostService,
		NewApplication,
	)

	return Application{}
}
