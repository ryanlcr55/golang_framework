package services

import "go_framework/intermal/respositories"

type PostServices struct {
	repo respositories.IPostRepo
}

func (ps PostServices) CreatePost(post respositories.Post) error {
	return ps.repo.Create(&post)
}

func NewPostService(repo respositories.IPostRepo) PostServices {
	return PostServices{
		repo: repo,
	}
}
