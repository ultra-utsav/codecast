package main

import (
	"log"
	"net/http"

	"github.com/codepod/app"
	"github.com/gorilla/mux"
)

func main() {
	app.InitApp()

	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
