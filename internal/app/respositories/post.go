package respositories

import (
	"context"
	"go_framework/internal/app/entities"
)

type IPostRepo interface {
	WithTransaction(tx any) (IPostRepo, error)
	FindByNo(ctx context.Context, no string) (entities.Post, error)
	Create(ctx context.Context, post *entities.Post) error
}
