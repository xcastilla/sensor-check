package data

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDBConnection() (*mongo.Database, error) {
	URI := fmt.Sprintf("mongodb://%s", os.Getenv("MONGO_URL"))
	clientOptions := options.Client().ApplyURI(URI).
		SetAuth(options.Credential{
			AuthSource: os.Getenv("MONGO_DB_NAME"),
			Username:   os.Getenv("MONGO_USER"),
			Password:   os.Getenv("MONGO_PASSWORD"),
		})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(os.Getenv("MONGO_DB_NAME"))
	return db, nil
}
