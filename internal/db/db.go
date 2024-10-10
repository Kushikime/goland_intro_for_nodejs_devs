package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// * reference some value
// & point the block of memory that stores that obj

// x := 10
// y := &x

// Println(y) -> 0x14000...

// to actually get the value of y which is a pointer to x we need to use *
// Println(*y) -> 10

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

type Collection string

const (
	ProductsCollection Collection = "products"
)

const (
	MONGODB_URI = "mongodb://localhost:27017"
	DB_NAME     = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(MONGODB_URI)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client
		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
