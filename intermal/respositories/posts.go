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

type PostRepo interface {
	FindByNo(no string) (Post, error)
	Create(post Post) error
}
