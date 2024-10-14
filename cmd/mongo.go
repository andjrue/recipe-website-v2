package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo(uri string) (*mongo.Client, error) {
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

func insertUser(db *mongo.Client, u *User) error {

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

func getAllUsers(db *mongo.Client) ([]User, error) {

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
