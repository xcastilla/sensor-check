package models

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Struct to pack results from DB
type SensorReading struct {
	Timestamp   time.Time
	Temperature float32
}

// Return all readings starting from fromDate
func GetReadings(db *mongo.Database, fromDate time.Time) ([]SensorReading, error) {
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
	results := []SensorReading{}
	for cur.Next(ctx) {
		reading := SensorReading{}
		err = cur.Decode(&reading)
		if err != nil {
			return nil, err
		}
		results = append(results, reading)
	}

	return results, nil
}
