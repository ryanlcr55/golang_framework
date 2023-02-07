package server

import (
	"fmt"
	"go_framework/internal/app"
	"google.golang.org/grpc"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func RunGRPCServer(application app.Application, registerServers ...func(server *grpc.Server)) {
	port := application.Configs.Server.GrpcPort
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	for _, register := range registerServers {
		register(grpcServer)
	}

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(grpcServer.Serve(listen))
}
