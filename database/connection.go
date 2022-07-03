package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

var (
	MongoURL string
)

func loadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "development" {
		log.Println("Running in development mode")
		MongoURL = os.Getenv("DEBUG_MONGO_URL") // Get url from env

	} else {
		log.Println("Running in production mode")
		MongoURL = os.Getenv("MONGO_URL") // Get url from env
	}
}

func Connect() error {

	loadVar()

	// Set client options
	clientOptions := options.Client().ApplyURI(MongoURL)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Connect to MongoDB
	client, e := mongo.Connect(ctx, clientOptions)
	if e != nil {
		return e
	}

	// Check the connection
	e = client.Ping(ctx, nil)
	if e != nil {
		return e
	}

	fmt.Println("Connected to mongoDB !")
	// get collection as ref
	db := client.Database("uca")

	Mg = MongoInstance{Client: client, Db: db}

	return nil
}
