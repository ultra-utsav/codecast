package app

import (
	"log"

	"github.com/codepod/config"
	"github.com/codepod/driver"
	"go.mongodb.org/mongo-driver/mongo"
)

// DB singleton instance
var DB *mongo.Database

// InitApp initializes an app
func InitApp() {
	db, er := driver.GetConnection(config.MongoURI)
	if er == nil {
		log.Printf("error in connecting to database, %v", er)

		return
	}

	DB = db

	log.Println("app initialization completed..")
}
