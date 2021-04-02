package user

import (
	"github.com/codepod/entities"
)

func generateUpdate(u entities.User) (string, []interface{}) {
	var query string

	var values []interface{}

	if u.Name != "" {
		query += "name=?, "

		values = append(values, u.Name)
	}

	if u.Email != "" {
		query += "email=?, "

		values = append(values, u.Email)
	}

	if u.Password != "" {
		query += "password=?"

		values = append(values, u.Password)
	}

	return query, values
}
