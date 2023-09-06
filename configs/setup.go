package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoDriver struct {
	Client *mongo.Client
}


func (m *MongoDriver) Connect() error  {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(EnvMongoURI()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(),opts)

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    _, errg := mongo.Connect(ctx,opts)
    if errg != nil {
        log.Fatal(errg)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

	m.Client = client

	log.Println("Connected to MongoDB...")
	return nil
}


func (m *MongoDriver) Disconnect() error  {
   if err := m.Client.Disconnect(context.TODO()); err != nil {
	return err
   }
   log.Println("Disconnected from MongoDB...")
   return nil
}
