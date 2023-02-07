package app

import "go_framework/internal/app/services"

type Application struct {
	Services Services
	//Logger
}

type Services struct {
	PostService services.PostServices
}

func NewApplication(
	PostService services.PostServices,
) Application {
	return Application{
		Services: Services{
			PostService: PostService,
		},
	}
}
