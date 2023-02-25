package server

import (
	"fmt"
	"go_framework"
	"go_framework/internal/genproto"
	"go_framework/internal/ports"
	"go_framework/internal/server"
	"google.golang.org/grpc"
	"os"
	"strings"
)

func main() {

	application := golang_framework.InitializeServer()

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
		srv := ports.NewHttpServer(application)
		server.RunHttpServer(application, srv)
	case "grpc":
		srv := ports.NewGrpcServer(application)
		server.RunGRPCServer(application, func(server *grpc.Server) {
			genproto.RegisterServiceServer(server, srv)
		})
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
