package ports

import (
	"go_framework/internal/app"
)

func NewHttpServer(application *app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

type HttpServer struct {
	app *app.Application
}
