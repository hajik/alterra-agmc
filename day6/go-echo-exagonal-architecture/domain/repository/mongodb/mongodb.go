package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var ctx = func() context.Context {
// 	return context.Background()
// }()

type MongoDB struct {
}

func (m *MongoDB) NewDB() (*mongo.Database, error) {
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	return nil, err
	// }

	// if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	// 	panic(err)
	// }

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	return nil, err
	// }

	// err = client.Connect(context.TODO())
	// if err != nil {
	// 	return nil, err
	// }

	// Check the connection
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client.Database("belajar_golang"), nil

}
