package stores

import (
	"context"

	"github.com/codepod/filters"

	"github.com/codepod/entities"
)

type User interface {
	Create(ctx context.Context, user *entities.User) error
	Find(ctx context.Context, filter *filters.User) (*entities.User, error)
	DeleteByID(ctx context.Context, id string) error
	Update(ctx context.Context, user *entities.User) error
}
