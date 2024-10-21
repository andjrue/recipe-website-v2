package recipes

import (
	"context"
	"log"
	"os"

	"github.com/andjrue/recipe-website-v2/internal/structs"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Recipe struct {
	Title       string
	Description string
	Ingredients string
	TimeToMake  string
	Steps       string
}

type Server structs.Server

func NewRecipe(title, descrip, ingre, ttm, steps string) *Recipe {
	return &Recipe{
		Title:       title,
		Description: descrip,
		Ingredients: ingre,
		TimeToMake:  ttm,
		Steps:       steps,
	}
}

func HandleAddRecipe(s *structs.Server, db *mongo.Client, username string) error {
	type Recipe struct {
		Title       string
		Description string
		Ingredients string
		TimeToMake  string
		Steps       string
	}

	recipe := NewRecipe("test title 1", "test descrip 1", "test ingre 1", "test ttm 1", "test steps 1")
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Issue loading env - insertUser")
	}

	DB := os.Getenv("DB")
	coll := db.Database(DB).Collection("users")
	filter := bson.M{"username": username}

	update := bson.M{"$set": bson.M{"recipes": recipe}}

	res, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Issue upating user - %v", err)
	}

	if res.MatchedCount == 0 {
		log.Printf("No user found with username - %v", username)
	}

    log.Printf("Recipe successfully added for User - %v", username)

	return nil
}

func HandleGetAllRecipes(s *structs.Server) error {
	return nil
}
