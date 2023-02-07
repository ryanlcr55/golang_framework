package entities

import "time"

type PostStatus int

const PostStatusEnable PostStatus = 1
const PostStatusDisable PostStatus = 2

type Post struct {
	ID        uint64
	No        string
	UserNo    string
	Title     string
	Content   string
	Status    PostStatus
	ShowAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
