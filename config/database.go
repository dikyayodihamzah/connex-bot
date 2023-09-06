package config

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbName   = os.Getenv("DB_NAME")
)

func NewDB() *mongo.Database {
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s/", username, password, host, port)
	option := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.Background(), option)
	if err != nil {
		log.Fatal(err)
	}

	// Ping database
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(dbName)

	log.Println("Connected to database")
	return database
}
