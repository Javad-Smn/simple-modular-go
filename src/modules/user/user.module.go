package user

import (
	"context"

	"github.com/Javad-Smn/simple-modular-go/src/db"
	libs "github.com/Javad-Smn/simple-modular-go/src/libs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection
var ctx context.Context

func collection() {
	client := db.Connect()
	userCollection = client.Database("go_test").Collection("users")
	ctx = context.TODO()
}
func routes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.Use(libs.ValidateUserToken())
	userRoute.GET("/me", getMe)
}
func InitModule(rg *gin.RouterGroup) {
	collection()
	routes(rg)
}
