package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Age      int32              `bson:"age,omitempty" json:"age"`
	Password string             `bson:"password,omitempty"`
}

type UserResponse struct {
	ID   primitive.ObjectID `json:"id"`
	Name string             ` json:"name"`
	Age  int32              ` json:"age"`
}
