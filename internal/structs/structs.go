package structs

import (
    
	"go.mongodb.org/mongo-driver/mongo"
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
	Addr string
	Db   *mongo.Client
}
