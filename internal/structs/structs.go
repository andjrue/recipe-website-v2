package structs

import (
    
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Email    string   `bson:"email"`
	Username string   `bson:"username"`
	Password string   `bson:"password"`
	Recipes  []Recipe `bson:"recipes"`
}

type Recipe struct {
	Title       string
	Description string
	Ingredients string
	TimeToMake  string
	Steps       string
}

type Server struct {
	addr string
	db   *mongo.Client
}
