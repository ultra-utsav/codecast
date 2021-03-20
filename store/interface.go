package store

import "github.com/codecast/entities"

type User interface {
	Create(*entities.User) error
	FindByID(int) (*entities.User, error)
	FindByEmail(string) (*entities.User, error)
}
