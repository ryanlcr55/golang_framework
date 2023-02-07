package services

import (
	"context"
	"go_framework/internal/app/respositories"
)

type PostServices struct {
	repo respositories.IPostRepo
}

func (ps PostServices) CreatePost(ctx context.Context, post respositories.Post) error {
	return ps.repo.Create(ctx, &post)
}

func NewPostService(repo respositories.IPostRepo) PostServices {
	return PostServices{
		repo: repo,
	}
}
