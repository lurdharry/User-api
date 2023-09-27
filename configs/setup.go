package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(EnvMongoURI()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {

		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, error := mongo.Connect(ctx, opts)

	if error != nil {

		log.Fatal(error)
	}

	//ping the database
	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal("sskksksksksks")
	//     log.Fatal(err)
	// }

	log.Println("Connected to MongoDB...")
	return client
}

// //Client instance
var DB *mongo.Client = Connect()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("user-api-golang").Collection(collectionName)
	return collection
}
