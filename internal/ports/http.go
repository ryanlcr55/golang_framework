package ports

import (
	"github.com/gin-gonic/gin"
	"go_framework/internal/app"
	"go_framework/internal/app/entities"
	"go_framework/internal/genapi"
	"net/http"
)

func NewHttpServer(application *app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

type HttpServer struct {
	app *app.Application
}

func (svr HttpServer) PostCreate(ctx *gin.Context) {
	req := genapi.PostCreate{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, genapi.Error{
			Message: "Post created failed",
		})
		return
	}

	data := entities.Post{
		UserNo:  req.UserNo,
		Title:   req.Title,
		Content: req.Content,
		Status:  entities.PostStatus(req.Status),
		ShowAt:  req.ShowAt.Time,
	}

	if err := svr.app.Services.PostService.CreatePost(ctx, data); err != nil {
		ctx.JSON(http.StatusInternalServerError, genapi.Error{
			Message: "Post created failed",
		})
	}
}
