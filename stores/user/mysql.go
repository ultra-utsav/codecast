package user

import (
	"context"
	"database/sql"
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

	return er
}

func (u *User) Find(ctx context.Context, filter *filters.User) (*entities.User, error) {
	var user entities.User

	query := "select * from users where "
	q, val := filter.WhereClause()
	query += q

	er := u.db.QueryRowContext(ctx, query, val).Scan(&user.UserID, &user.Name, &user.Email, &user.Password)
	if er != nil {
		return nil, er
	}

	return &user, er
}

func (u *User) Update(ctx context.Context, user *entities.User) error {
	q, values := generateUpdate(*user)

	query := "update users set "

	if len(values) != 0 {
		query += q
	}

	query += " where id=?"
	values = append(values, user.UserID)

	_, er := u.db.ExecContext(ctx, query, values...)

	return er
}

func (u *User) DeleteByID(ctx context.Context, id string) error {
	query := "delete from users where id=?"

	_, er := u.db.ExecContext(ctx, query, id)

	return er
}
