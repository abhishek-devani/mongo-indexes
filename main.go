package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gsgit.gslab.com/poc/indexes"
	"gsgit.gslab.com/poc/models"
	"gsgit.gslab.com/poc/utils"
)

func main() {

	db := utils.Database("some_database")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	// col := client.Database("some_database").Collection("Some Collection")

	var tmp = []models.SimpleIndexes{
		{
			CollectionName: "temp",
			Indexes: []models.Simple{
				{
					Name: "temp",
					Keys: []primitive.E{
						{Key: "p1", Value: 1},
						{Key: "p2", Value: 1},
					},
					Unique: true,
					Sparse: true,
				},
				{
					Name: "filter_updatedat",
					Keys: []primitive.E{
						{Key: "p3", Value: 1},
						{Key: "p4", Value: 1},
						{Key: "p5", Value: 1},
					},
					Unique: false, Sparse: false,
				},
			},
		},
		{
			CollectionName: "some_database",
			Indexes: []models.Simple{
				{
					Name: "unique_project",
					Keys: []primitive.E{
						{Key: "p6", Value: 1},
						{Key: "p7", Value: 1},
						{Key: "p8", Value: 1},
					},
					Unique: true, Sparse: true,
				},
			},
		},
	}

	coll := db.Collection("some_database")
	var err error

	err = indexes.DeleteIndexes(coll, ctx)
	if err != nil {
		fmt.Println("Indexes().DropAll() ERROR:", err)
		os.Exit(1)
	}

	err = indexes.ApplyIndexes(db, tmp, ctx)
	if err != nil {
		fmt.Println("Indexes().CreateMany() ERROR:", err)
		os.Exit(1) // exit in case of error
	}

	err = indexes.GetIndexes(coll)
	if err != nil {
		fmt.Println("get index error: ", err)
		os.Exit(1) // exit in case of error
	}
}
