package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go_framework/internal/app/entities"
	"go_framework/internal/app/respositories"
	mock_respositories "go_framework/mock/respositories"
	"time"

	"testing"
)

func TestPostServices_CreatePost(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_respositories.NewMockIPostRepo(ctl)
	fakePost := entities.Post{
		No:      uuid.NewString(),
		Title:   "TestTitle",
		Content: "TestContent",
		Status:  entities.PostStatusEnable,
		ShowAt:  time.Now(),
	}
	gomock.InOrder(
		repo.EXPECT().Create(
			context.Background(),
			&fakePost,
		).Times(1).Return(nil),
	)

	type fields struct {
		repo respositories.IPostRepo
	}
	type args struct {
		ctx  context.Context
		post entities.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "success",
			fields: struct{ repo respositories.IPostRepo }{repo: repo},
			args: struct {
				ctx  context.Context
				post entities.Post
			}{
				ctx:  context.Background(),
				post: fakePost,
			},
			wantErr: false,
		},
		{
			name:   "no uuid",
			fields: struct{ repo respositories.IPostRepo }{repo: repo},
			args: struct {
				ctx  context.Context
				post entities.Post
			}{
				ctx: context.Background(),
				post: entities.Post{
					Title:   "Test",
					Content: "Test",
					ShowAt:  time.Now(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := PostServices{
				repo: tt.fields.repo,
			}
			if err := ps.CreatePost(tt.args.ctx, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
