package main

import (
	"log"
	"net/http"
	"os"
    "context"
    "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) checkUserSignin(w http.ResponseWriter, r *http.Request) error {
	/*  We will need to query the DB with provided credentials from the user.
	        -> Username & Pass

	    Passwords will be hashed at this point, so that's really what were looking for.
	    I'm assuming the hashed passes will come out the same each time, so we can take what they entered, hash it, and then compare what is in the db.

	    This sounds pretty easy, but we'll see.
	*/

    username := "testuser4"
    password := "testpass100"

	log.Println("Checking username")

	DB := os.Getenv("DB")
	coll := s.db.Database(DB).Collection("users")
	filter := bson.M{"username": username}

	var result User

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            writeJson(w, http.StatusBadRequest, err)

            log.Fatal("No user found with that username")
        }
    }

    err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
    if err != nil {
        writeJson(w, http.StatusBadRequest, err)
        log.Fatal("Passwords do not match")
    }
    log.Println("Check sucsessful. User signed in.")
    return nil
}
