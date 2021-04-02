package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/codepod/entities"

	"github.com/codepod/services"
)

type User struct {
	service services.User
}

func New(service services.User) *User {
	return &User{service: service}
}

func (u *User) Create(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)

	var user entities.User

	er := json.Unmarshal(body, &user)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("unable to parse request body"))

		return
	}

	if user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("name can not be empty"))

		return
	}

	if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("email can not be empty"))

		return
	}

	if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("password can not be empty"))

		return
	}

	er = u.service.Create(req.Context(), &user)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("unable to create user, %v", er.Error())))
	}
}

func (u *User) Find(w http.ResponseWriter, req *http.Request) {
	filter := req.URL.Query()["filter"][0]

	res, er := u.service.Find(req.Context(), filter)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unable to find user"))

		return
	}

	res.Password = ""

	body, er := json.Marshal(res)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, _ = w.Write(body)
}

func (u *User) Update(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	body, _ := ioutil.ReadAll(req.Body)

	var user entities.User
	user.UserID = id

	er := json.Unmarshal(body, &user)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("unable to parse request body"))
	}

	er = u.service.Update(req.Context(), &user)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("unable to update user, %v", er.Error())))
	}
}

func (u *User) Delete(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	er := u.service.DeleteByID(req.Context(), id)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("unable to delete user"))

		return
	}
}
