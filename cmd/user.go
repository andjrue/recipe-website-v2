package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	// Recipes []Recipe -- Will add this later
}

func newUser(email, username, password string) *User {
	return &User{
		Email:    email,
		Username: username,
		Password: password,
	}
}

func checkUsernameAndPass(db *mongo.Client, username, password string) (bool, bool) {


	userNoExists := make(chan bool, 1)
	passIsGood := make(chan bool, 1)

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Issue loading env - insertUser")
	}
	go func() {
// TODO -- Maybe make this more complicated? Not sure how I want to do this yet
		log.Println("Checking pass")
		if len(password) >= 8 {
			log.Println("Password passes")
			passIsGood <- true
            log.Println("Pass send to passisgood")
		}
		log.Println("a")
		passIsGood <- false
	}()

	go func() {
		log.Println("Checking username")
		DB := os.Getenv("DB")
		coll := db.Database(DB).Collection("users")
		filter := bson.M{"username": username}

		var result bson.M

		err := coll.FindOne(context.TODO(), filter).Decode(&result)
		log.Println("b")
		if err != nil {
			log.Println("c")
			if err == mongo.ErrNoDocuments {
				log.Printf("Did not find user in DB: %v", username)
				userNoExists <- true
                log.Println("sent to usernoexists")
			} 
		} else {
            log.Printf("Result from FindOne: %v", result)
            userNoExists <- false
        }

	}()

	log.Println("e")
	user := <-userNoExists
	pass := <-passIsGood
	log.Println("f")

	close(userNoExists)
	close(passIsGood)
	log.Println("g")

	log.Println("Checks completed successfully. User passed")
	return user, pass
}

func checkEmail(email string) error {
	return nil
}
