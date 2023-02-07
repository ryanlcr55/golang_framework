package respositories

import (
	"context"
	"time"
)

type Post struct {
	ID        uint64
	No        string
	UserNo    string
	Title     string
	Content   string
	Status    string
	ShowAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type IPostRepo interface {
	WithTransaction(tx any) (IPostRepo, error)
	FindByNo(ctx context.Context, no string) (Post, error)
	Create(ctx context.Context, post *Post) error
}
