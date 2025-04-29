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

var DB *mongo.Database

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	fmt.Println(os.Getenv("SECRET_KEYWORD"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error after trying to connect to Mongodb")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB doesn't replied the ping")
	}

	DB = client.Database(dbName)
	log.Println(DB)
	log.Println("MongoDB connection established")
}
