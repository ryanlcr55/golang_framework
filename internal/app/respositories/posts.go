package respositories

import (
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
	FindByNo(no string) (Post, error)
	Create(post *Post) error
}
