package data

import (
	"context"
	"fmt"
	"os"
	"time"
	
	"../models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

var db *mongo.Database

func InitDBConnection() (error) {
	URI := fmt.Sprintf("mongodb://%s", os.Getenv("MONGO_URL"))
	clientOptions := options.Client().ApplyURI(URI).
		SetAuth(options.Credential{
			AuthSource: os.Getenv("MONGO_DB_NAME"),
			Username:   os.Getenv("MONGO_USER"),
			Password:   os.Getenv("MONGO_PASSWORD"),
		})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	db = client.Database(os.Getenv("MONGO_DB_NAME"))
	return nil
}


// Return all readings starting from fromDate
func GetReadings(fromDate time.Time) ([]models.SensorReading, error) {
	collection := db.Collection("measurements2")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	
	projection := bson.D{
		{"_id", 0},
		{"temperature", 1},
		{"timestamp", 1},
	}

	var query = bson.M{"timestamp": bson.M{"$gt": fromDate}}
	cur, err := collection.Find(ctx, query, options.Find().SetProjection(projection),)
	if err != nil { 
		return nil, err
	}
	defer cur.Close(ctx)

	// Pack results
	results := []models.SensorReading{}
	for cur.Next(ctx) {
		reading := models.SensorReading{}
		err = cur.Decode(&reading)
		if err != nil {
			return nil, err
		}
		results = append(results, reading)
	}

	return results, nil
}
