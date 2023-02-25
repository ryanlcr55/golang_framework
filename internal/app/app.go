package app

import (
	"go_framework/internal/app/respositories"
	"go_framework/internal/app/services"
	"go_framework/internal/configs"
)

type Application struct {
	Configs    *configs.Configs
	Services   Services
	TrxHandler respositories.ITrxHandler
	//Logger
}

type Services struct {
	PostService services.PostServices
}

func NewApplication(
	configs *configs.Configs,
	PostSrv services.PostServices,
	trxHandler respositories.ITrxHandler,
) Application {
	return Application{
		Configs: configs,
		Services: Services{
			PostService: PostSrv,
		},
		TrxHandler: trxHandler,
	}
}
