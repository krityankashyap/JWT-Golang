package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env") // load the env variable

	if err != nil {
		log.Fatal("Error loading the env variable");
	}

	MongoDb:= os.Getenv("MONGO_URL");

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err);
	}
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second); // WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).

	defer cancel(); // it calls automatically after the function is executed

	err = client.Connect(ctx);
  if err != nil {
		log.Fatal(err);
	}

	// database connection is established
  fmt.Println("Database connection established successfully");

	return client;

}

var Client *mongo.Client = DBinstance();

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName);

	return collection;
}

