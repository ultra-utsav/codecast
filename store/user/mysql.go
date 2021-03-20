package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/codepod/filters"

	error2 "github.com/codepod/error"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/codepod/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

// User store
type User struct {
	col *mongo.Collection
}

func New(col *mongo.Collection) *User {
	return &User{col: col}
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

	_, er = u.col.InsertOne(ctx, user)
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

	er := u.col.FindOne(ctx, filter.GenFilter()).Decode(&user)
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

func (u *User) DeleteByID(ctx context.Context, id string) error {
	_, er := u.col.DeleteOne(ctx, bson.M{"_id": id})

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

func (u *User) Update(ctx context.Context, user entities.User) (*entities.User, error) {
	update := generateUpdate(user)

	res, er := u.col.UpdateByID(ctx, user.UserID, update)
	if er != nil || res.ModifiedCount == 0 {
		return nil, &error2.PodError{
			Code:     http.StatusInternalServerError,
			Err:      er.Error(),
			Message:  fmt.Sprintf("unable to update user details for userID: %v", user.UserID),
			Location: "DB",
		}
	}

	return nil, nil
}
