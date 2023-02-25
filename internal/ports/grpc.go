package ports

import (
	"context"
	"go_framework/internal/app"
	"go_framework/internal/app/entities"
	"go_framework/internal/genproto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	app *app.Application
	genproto.UnimplementedServiceServer
}

func NewGrpcServer(application *app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (srv GrpcServer) PostCreate(ctx context.Context, req *genproto.PostCreateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, srv.app.Services.PostService.CreatePost(
		ctx, entities.Post{
			UserNo:  req.UserNo,
			Title:   req.Title,
			Content: req.Content,
			Status:  entities.PostStatus(req.Status),
			ShowAt:  req.ShowAt.AsTime(),
		})
}
