package services

import (
	"context"

	"github.com/codepod/entities"
)

type User interface {
	Create(ctx context.Context, user *entities.User) error
	Find(ctx context.Context, filter string) (*entities.User, error)
	DeleteByID(ctx context.Context, id string) error
	Update(ctx context.Context, user *entities.User) error
}
