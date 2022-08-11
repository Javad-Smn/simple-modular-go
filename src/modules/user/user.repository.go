package user

import (
	"github.com/Javad-Smn/simple-modular-go/src/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne(data structs.User) *mongo.InsertOneResult {
	insertResult, err := userCollection.InsertOne(ctx, data)
	if err != nil {
		panic(err)
	}
	return insertResult
}
func FindUserPassword(name string) (string, error) {
	var user structs.User
	err := userCollection.FindOne(ctx, bson.D{primitive.E{Key: "name", Value: name}}, &options.FindOneOptions{Projection: bson.M{"password": 1}}).Decode(&user)

	if err != nil {
		return user.Password, err
	}
	return user.Password, nil
}
func FindOne(name string) (structs.User, error) {
	var user structs.User
	err := userCollection.FindOne(ctx, bson.D{primitive.E{Key: "name", Value: name}}, &options.FindOneOptions{Projection: bson.M{"password": 0}}).Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}
