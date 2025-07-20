package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	env "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DataBase *mongo.Database

func ConnectDB() {
	err := env.Load()
	logError(err, "Error loading .env file")

	mongoConnectionString := os.Getenv("MONGO_DB_URI")
	db := os.Getenv("MONGO_DB_NAME")

	if mongoConnectionString == "" {
		log.Fatal("MONGO_DB_URI configuration missing")
	}
	if db == "" {
		log.Fatal("MONGO_DB_NAME configuration missing")
	}

	params := options.Client().ApplyURI(mongoConnectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, params)
	logError(err, "Unable to connect to MongoDB:")

	err = client.Ping(ctx, nil)
	logError(err, "Unable to connect to MongoDB:")

	fmt.Println("Successfully Connected to MongoDB!")

	DataBase = client.Database(db)
}

func logError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
