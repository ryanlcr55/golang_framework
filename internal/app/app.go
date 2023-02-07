package app

import (
	"go_framework/internal/app/services"
	"go_framework/internal/configs"
)

type Application struct {
	Configs  configs.Configs
	Services Services
	//Logger
}

type Services struct {
	PostService services.PostServices
}

func NewApplication(
	configs configs.Configs,
	PostSrv services.PostServices,
) Application {
	return Application{
		Configs: configs,
		Services: Services{
			PostService: PostSrv,
		},
	}
}
