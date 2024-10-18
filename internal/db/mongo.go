package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
    "net/http"

	"github.com/andjrue/recipe-website-v2/internal/structs"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




type User *structs.User

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}


func ConnectToMongo(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func InsertUser(db *mongo.Client, u *structs.User) error {

    envErr := godotenv.Load()
    if envErr != nil {
        log.Fatal("Issue loading env - insertUser")
    }

    DB := os.Getenv("DB") 

	coll := db.Database(DB).Collection("users")

	result, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}

func GetAllUsers(db *mongo.Client) ([]User, error) {

    envErr := godotenv.Load()
    if envErr != nil {
        log.Fatal("Issue loading env - insertUser")
    }

    DB := os.Getenv("DB") 
	coll := db.Database(DB).Collection("users")

	filter := bson.M{}
	// NOTE TO SELF - If order does matter use bson.D

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Issue getting all users: ", err)
		panic(err)
	}

	var results []User

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
    return results, nil
}

func UpdateUser(db *mongo.Client, username, newPassword string) error {

    envErr := godotenv.Load()
    if envErr != nil {
        log.Fatal("Issue loading env - insertUser")
    }
    DB := os.Getenv("DB")
    coll := db.Database(DB).Collection("users")
    filter := bson.M{"username": username}
    
    update := bson.M{"$set": bson.M{"password": newPassword}}

    res, err := coll.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        panic(err)
    }
    if res.MatchedCount == 0 {
        log.Printf("No user found with username - %v", username)
    } 

    log.Printf("Successfully updated password for %v", username)
    return nil

}

func DeleteUser(db *mongo.Client, username string) error {
    envErr := godotenv.Load()
    if envErr != nil {
        log.Fatal("Issue loading env - insertUser")
    }
    
    DB := os.Getenv("DB")
    coll := db.Database(DB).Collection("users")
    filter := bson.M{"username": username}

    res, err := coll.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Printf("Issue deleting user - %v: %v\n", username, err)
        panic(err)
    }
    log.Printf("User successfully deleted: %v", res)

    return nil

}
