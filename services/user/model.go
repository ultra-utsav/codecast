package user

import (
	"context"
	"net/http"
	"strings"

	error2 "github.com/codepod/error"

	"github.com/codepod/filters"

	"github.com/codepod/entities"
	"github.com/codepod/stores"
)

type Service struct {
	store stores.User
}

func New(store stores.User) *Service {
	return &Service{store: store}
}

func (s *Service) Create(ctx context.Context, user *entities.User) error {
	//TODO: password encryption is remaining
	return s.store.Create(ctx, user)
}

func (s *Service) Find(ctx context.Context, filter string) (*entities.User, error) {
	var f filters.User
	if strings.Contains(filter, "@") {
		f.Email = filter
	} else {
		f.ID = filter
	}

	u, er := s.store.Find(ctx, &f)
	if er != nil {
		return nil, &error2.PodError{
			Code: http.StatusInternalServerError,
			Err:  er.Error(),
		}
	}

	return u, nil
}

func (s *Service) DeleteByID(ctx context.Context, id string) error {
	er := s.store.DeleteByID(ctx, id)
	if er != nil {
		return &error2.PodError{
			Code: http.StatusInternalServerError,
			Err:  er.Error(),
		}
	}

	return nil
}

func (s *Service) Update(ctx context.Context, user *entities.User) error {
	return s.store.Update(ctx, user)
}
