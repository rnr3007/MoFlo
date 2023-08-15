package models

import (
	"context"
	"fmt"
	"moflo-be/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CollectionUser() *mongo.Collection {
	collectionUser := config.DB.Database("moflo").Collection("user")
	return collectionUser
}

type User struct {
	UserId   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName string             `json:"fullName,omitempty" bson:"fullName,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

func InitUserCollection() {
	indexModel := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := CollectionUser().Indexes().CreateMany(context.Background(), indexModel)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
