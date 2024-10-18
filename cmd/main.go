package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
    "github.com/andjrue/recipe-website-v2/internal/db"
    "github.com/andjrue/recipe-website-v2/internal/router"
)


func main() {
    // DB

    envErr := godotenv.Load()
    if envErr != nil {
        log.Println("Error opening ENV: ", envErr)
    }

    uri := os.Getenv("MONGODB_URI")
    client, err := db.ConnectToMongo(uri)
    if err != nil {
        log.Println("Issue connecting to DB: \n", err)
    }

    fmt.Println("Connected to DB")
    defer client.Disconnect(context.Background())
    
    // SERVER
    s := router.NewServer(":6969", client)

    // Hey, let's make sure we dont have another terminal open still running the same port. That will cause a lot of unecessary problems. 

    log.Println("Listening on port: ", s)

    router.Run(s)

    log.Println("Server closed?")


}

