package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Simple struct {
	Name   string
	Keys   []primitive.E
	Unique bool
	Sparse bool
}

type SimpleIndexes struct {
	CollectionName string
	Indexes        []Simple
}
