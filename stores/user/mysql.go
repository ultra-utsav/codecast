package user

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/codepod/filters"

	"github.com/codepod/entities"
	error2 "github.com/codepod/error"
)

// User stores
type User struct {
	db *sql.DB
}

func New(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) Create(ctx context.Context, user *entities.User) error {
	_, er := u.Find(ctx, &filters.User{Email: user.Email})
	if er == nil {
		return &error2.PodError{
			Code:     http.StatusBadRequest,
			Message:  "email already in use",
			Location: "DB",
		}
	}

	query := "insert into users values(null,?,?,?)"
	// TODO: encrypt password
	_, er = u.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	if er != nil {
		return &error2.PodError{
			Code:     http.StatusInternalServerError,
			Err:      er.Error(),
			Message:  "unable to insert record",
			Location: "DB",
		}
	}

	return nil
}

func (u *User) Find(ctx context.Context, filter *filters.User) (*entities.User, error) {
	var user entities.User

	query := "select * from users where "
	q, val := filter.WhereClause()
	query += q

	er := u.db.QueryRowContext(ctx, query, val).Scan(&user.UserID, &user.Name, &user.Email, &user.Password)
	if er != nil {
		return nil, &error2.PodError{
			Code:     http.StatusInternalServerError,
			Err:      er.Error(),
			Message:  "unable to find record with given data",
			Location: "DB",
		}
	}

	return &user, nil
}

func (u *User) Update(ctx context.Context, user *entities.User) error {
	q, values := generateUpdate(*user)

	query := "update users set "

	if len(values) != 0 {
		query += q
	}

	_, er := u.db.ExecContext(ctx, user.UserID, query, values)
	if er != nil {
		return &error2.PodError{
			Code:     http.StatusInternalServerError,
			Err:      er.Error(),
			Message:  fmt.Sprintf("unable to update user details for userID: %v", user.UserID),
			Location: "DB",
		}
	}

	return nil
}

func (u *User) DeleteByID(ctx context.Context, id string) error {
	query := "delete form users where id=?"

	_, er := u.db.ExecContext(ctx, query, id)
	if er != nil {
		return &error2.PodError{
			Code:     http.StatusInternalServerError,
			Err:      er.Error(),
			Message:  fmt.Sprintf("unable to delete record with given id: %v", id),
			Location: "DB",
		}
	}

	return nil
}
