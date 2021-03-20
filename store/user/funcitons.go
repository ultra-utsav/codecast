package user

import (
	"github.com/codepod/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateUpdate(u entities.User) primitive.M {
	filter := bson.M{}

	if u.Email != "" {
		filter["email"] = u.Email
	}

	if u.Name != "" {
		filter["name"] = u.Name
	}

	if u.Password != "" {
		filter["password"] = u.Password
	}

	return filter
}
