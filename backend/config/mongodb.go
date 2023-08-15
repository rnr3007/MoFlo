package config

import (
	"context"
	"moflo-be/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TimeOut)
	defer cancel()
	DB, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}
