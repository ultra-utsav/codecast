package driver

import (
	"context"
	"time"

	"github.com/codecast/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetConnection connects to mongo client and returns database connection
func GetConnection(uri string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)

	clientOptions.SetAppName(config.App)
	clientOptions.SetConnectTimeout(5 * time.Minute)

	er := clientOptions.Validate()
	if er != nil {
		return nil, er
	}

	client, er := mongo.Connect(context.TODO(), clientOptions)
	if er != nil {
		return nil, er
	}

	er = client.Ping(context.TODO(), nil)
	if er != nil {
		return nil, er
	}

	return client.Database(config.App), nil
}
