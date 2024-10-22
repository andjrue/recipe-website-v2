package recipes

import (
	"context"
	"fmt"
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
/*
    type Recipe struct {
		Title       string
		Description string
		Ingredients string
		TimeToMake  string
		Steps       string
	}
*/
	
    recipe := NewRecipe("test array", "test desc", "test ingre", "test ttm", "test steps")

    if db == nil {
        log.Fatal("didnt initialize db")
    }

    fmt.Printf("Recipe - %v", recipe)

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Issue loading env - insertUser")
	}

	DB := os.Getenv("DB")
    if DB == "" {
        log.Println("Didn't grab db env")
    }

	coll := db.Database(DB).Collection("users")
    if coll == nil {
        return fmt.Errorf("Collection not found in db")
    }
	filter := bson.M{"username": username}

    update := bson.M{"$push": bson.M{"recipes": recipe}} 
    //$set creates a unique problem here. It's "setting" the recipes value and will overwrite anything thats been there before. 
    // Good find honestly, I might be able to use that somewhere else, but this doesn't work for what we're trying to do. 

	res, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Issue updating user - %v", err)
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
