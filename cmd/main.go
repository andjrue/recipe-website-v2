package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main() {
    // DB

    envErr := godotenv.Load()
    if envErr != nil {
        log.Println("Error opening ENV: ", envErr)
    }

    uri := os.Getenv("MONGODB_URI")
    client, err := connectToMongo(uri)
    if err != nil {
        log.Println("Issue connecting to DB: \n", err)
    }

    fmt.Println("Connected to DB")
    defer client.Disconnect(context.Background())
    
    // SERVER
    s := NewServer(":6969", client)
    log.Println("Listening on port: ", s)

    s.Run()

    log.Println("Server closed?")

    select {}

}

