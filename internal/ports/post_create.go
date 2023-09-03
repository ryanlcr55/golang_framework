package ports

import (
	"context"
	"go_framework/internal/app/entities"
	"go_framework/internal/genapi"
	"net/http"
)

func (svr HttpServer) PostCreate(ctx context.Context, request genapi.PostCreateRequestObject) (genapi.PostCreateResponseObject, error) {
	data := entities.Post{
		UserNo:  request.Body.UserNo,
		Title:   request.Body.Title,
		Content: request.Body.Content,
		Status:  entities.PostStatus(request.Body.Status),
		ShowAt:  request.Body.ShowAt.Time,
	}

	if err := svr.app.Services.PostService.CreatePost(ctx, data); err != nil {
		return genapi.PostCreatedefaultJSONResponse{
			StatusCode: http.StatusBadGateway,
			Body: genapi.Error{
				Message: err.Error(),
			},
		}, err
	}

	return genapi.PostCreate200Response{}, nil
}
