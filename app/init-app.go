package app

import (
	"log"

	"github.com/codecast/config"
	"github.com/codecast/driver"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database

func InitApp() {
	db, er := driver.GetConnection(config.MongoURI)
	if er == nil {
		log.Printf("error in connecting to database, %v", er)

		return
	}

	DB = db

	log.Println("app initialization completed..")
}
