package server

import (
	"fmt"
	"go_framework/internal/app"
	"go_framework/internal/server"
	"os"
	"strings"
)

func main() {

	application := app.InitializeServer()

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
	case "grpc":
		server.RunGRPCServer(application)
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
