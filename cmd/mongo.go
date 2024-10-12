package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo(uri string) (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI(uri)

    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
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
    coll := db.Database("recipe-website").Collection("users")

    result, err := coll.InsertOne(context.TODO(), u)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
    return nil
}
