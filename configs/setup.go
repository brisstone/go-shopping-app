package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// *mongo.Client: implies The return type of the function is a pointer to a mongo.Client object.
// The * denotes a pointer type.
func ConnectDB() *mongo.Client {
	// Create a ConnectDB function that first configures the client to use
	// the correct URI and check for errors.
	log.Print(EnvMongoURI())
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
	    log.Fatal(err)
	}

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(EnvMongoURI()))
	// if err != nil {
	// 	panic(err)
	// }

	// if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	// 	panic(err)
	// }

	// defined a timeout of 10 seconds we wanted to use when trying to connect.
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)
	// check if there is an error while connecting to the database
	// and cancel the connection if the connecting period exceeds 10 seconds
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	// pinged the database to test our connection and returned the client instance.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
// Create a DB variable instance of the ConnectDB.
// This will come in handy when creating collections
var DB *mongo.Client = ConnectDB()

// getting database collections
// Create a GetCollection function to retrieve
// and create collections on the database.
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
