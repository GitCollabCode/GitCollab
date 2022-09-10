package data

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DATABASE            = "gitcollab"
	PROFILES_COLLECTION = "profiles"
)

type MongoDriver struct {
	Client mongo.Client
	Log    *logrus.Logger
}

func InitMongoDBDriver(log *logrus.Logger) (*MongoDriver, error) {
	uri := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create a new client and connect to the server
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	//ping server
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Info("Successfully connected and pinged MongoDB server.")

	return &MongoDriver{Client: *client, Log: log}, nil
}
