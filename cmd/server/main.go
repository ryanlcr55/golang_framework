package server

import (
	"fmt"
	"go_framework/internal/app"
	"go_framework/internal/genproto"
	"go_framework/internal/ports"
	"go_framework/internal/server"
	"google.golang.org/grpc"
	"os"
	"strings"
)

func main() {

	application := app.InitializeServer()

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
	case "grpc":
		svc := ports.NewGrpcServer(application)
		server.RunGRPCServer(application, func(server *grpc.Server) {
			genproto.RegisterServiceServer(server, svc)
		})
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
