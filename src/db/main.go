package db

import (
	"context"
	"time"

	libs "github.com/Javad-Smn/simple-modular-go/src/libs"
	"github.com/fatih/color"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	var uri = `mongodb://` + libs.DotEnvVariable("DB_HOST") + `:` + libs.DotEnvVariable("DB_PORT")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	color.Green("Database Successfully Connected ⛁")
	return client
}
