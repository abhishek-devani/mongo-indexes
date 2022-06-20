package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Database(dbname string) (db *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("database connect error", err)
		return
	}
	db = client.Database(dbname)
	return
}
