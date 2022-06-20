package indexes

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gsgit.gslab.com/poc/models"
)

func GetIndexes(coll *mongo.Collection) error {
	// collection := db.Collection("temp")
	indexView := coll.Indexes()
	opts := options.ListIndexes().SetMaxTime(2 * time.Second)
	cursor, err := indexView.List(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
		return err
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
		return err
	}

	for _, v := range result {
		for _, v1 := range v {
			if reflect.ValueOf(v1).Kind() == reflect.Map {
				v1a := v1.(primitive.M)
				var index bson.D
				// fmt.Printf("%v: {\n", k1)
				for k2, v2 := range v1a {
					// if k2 == "_id" {
					// 	continue
					// }
					index = append(index, primitive.E{Key: k2, Value: v2})
					// fmt.Printf("  %v: %v\n", k2, v2)
				}
				fmt.Println(index)
				// fmt.Printf("}\n")
			} else {
				fmt.Println()
				// fmt.Printf("%v: %v\n", k1, v1)
			}
		}
		fmt.Println()
	}
	return nil
}

func ApplyIndexes(db *mongo.Database, tmp []models.SimpleIndexes, ctx context.Context) error {

	for _, val := range tmp {
		coll := db.Collection(val.CollectionName)
		for _, v1 := range val.Indexes {
			var mod []mongo.IndexModel
			Unique := v1.Unique
			Sparse := v1.Sparse
			// var index bson.D
			// for _, v2 := range v1.Keys {
			// 	index = append(index, primitive.E{Key: v2, Value: 1})
			// }
			// fmt.Println(index)
			// fmt.Println()
			mod = []mongo.IndexModel{
				{
					Keys: v1.Keys,
					Options: &options.IndexOptions{
						Unique: &Unique,
						Sparse: &Sparse,
					},
				},
			}
			opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
			_, err := coll.Indexes().CreateMany(ctx, mod, opts)

			// Check for the options errors
			if err != nil {
				return err
			} else {
				fmt.Println("CreateMany() option:", opts)
			}
		}
	}
	return nil
}

func DeleteIndexes(coll *mongo.Collection, ctx context.Context) error {

	_, err := coll.Indexes().DropAll(ctx)

	if err != nil {
		return err
	} else {
		// fmt.Println("DropAll() result type:", reflect.TypeOf(dropAllResult))
		// fmt.Println("DropAll() result:", dropAllResult)
		fmt.Println("Delete Successfully")
	}
	return err
}
