package filters

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name  string
	Email string
}

func (u User) GenFilter() primitive.M {
	f := bson.M{}
	if u.Name != "" {
		f["name"] = u.Name
	}

	if u.Email != "" {
		f["email"] = u.Email
	}

	return f
}
