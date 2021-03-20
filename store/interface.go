package store

import (
	"context"

	"github.com/codepod/entities"
	"github.com/codepod/filters"
)

type User interface {
	Create(ctx context.Context, user *entities.User) error
	Find(ctx context.Context, filter *filters.User) (*entities.User, error)
	DeleteByID(ctx context.Context, id string) error
	Update(ctx context.Context, user entities.User) (*entities.User, error)
}