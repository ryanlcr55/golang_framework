package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_framework/internal/app"
	"go_framework/internal/ports"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func RunHttpServer(application app.Application, srv ports.HttpServer) {
	r := gin.Default()
	r = ports.RegisterHandlers(r, srv)

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%s", application.Configs.Server.HttpPort),
		Handler: r,
	}

	ch := make(chan os.Signal, 1)
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("SERVER GG惹:", err)
		}
	}()
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
