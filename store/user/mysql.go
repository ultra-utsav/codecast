package user

import (
	"context"
	"fmt"
	"net/http"

	error2 "github.com/codecast/error"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/codecast/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

// User store
type User struct {
	col *mongo.Collection
}

func New(col *mongo.Collection) *User {
	return &User{col: col}
}

func (u *User) Create(user *entities.User) error {
	_, er := u.FindByEmail(user.Email)
	if er == nil {
		return &error2.CodeError{
			Code:     http.StatusBadRequest,
			Err:      error2.ErrMongoOperation,
			Message:  "email already in use",
			Location: "DB",
		}
	}

	_, er = u.col.InsertOne(context.TODO(), user)
	if er != nil {

		return &error2.CodeError{
			Code:     http.StatusInternalServerError,
			Err:      error2.ErrMongoOperation,
			Message:  "unable to insert record",
			Location: "DB",
		}
	}

	return nil
}

func (u *User) FindByID(id int) (*entities.User, error) {
	var user entities.User

	er := u.col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if er != nil {
		return nil, &error2.CodeError{
			Code:     http.StatusInternalServerError,
			Err:      error2.ErrMongoOperation,
			Message:  fmt.Sprintf("unable to find record with given id: %v", id),
			Location: "DB",
		}
	}

	return &user, nil
}

func (u *User) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	er := u.col.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if er != nil {
		return nil, &error2.CodeError{
			Code:     http.StatusInternalServerError,
			Err:      error2.ErrMongoOperation,
			Message:  fmt.Sprintf("unable to find record with given email: %v", email),
			Location: "DB",
		}
	}

	return &user, nil
}
