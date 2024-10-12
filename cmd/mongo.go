package main

import (
	"context"
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
