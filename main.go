package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	user3 "github.com/codepod/delivery/user"
	"github.com/codepod/driver"
	user2 "github.com/codepod/services/user"
	"github.com/codepod/stores/user"
	"github.com/gorilla/mux"
)

func main() {
	db, er := driver.GetConnection(&driver.MySQLConfig{Host: "localhost", User: "root", Password: "", Port: "3306", DB: "codepod"})
	if er != nil {
		log.Fatal(er)
	}

	userStore := user.New(db)
	userService := user2.New(userStore)
	userHTTP := user3.New(userService)

	router := mux.NewRouter()

	router.HandleFunc("/user", userHTTP.Create).Methods(http.MethodPost)
	router.HandleFunc("/user", userHTTP.Find).Methods(http.MethodGet).Queries("filter", "{filter}")
	router.HandleFunc("/user/{id}", userHTTP.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/user/{id}", userHTTP.Update).Methods(http.MethodPut)

	log.Println("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
