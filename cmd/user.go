package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/andjrue/recipe-website-v2/cmd/recipes/recipes.go"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Email    string           `bson:"email"`
	Username string           `bson:"username"`
	Password string           `bson:"password"`
	Recipes  []recipes.Recipe `bson:"recipes"`
}

func newUser(email, username, password string, recipe recipes.Recipe) *User {
	return &User{
		Email:    email,
		Username: username,
		Password: password,
        Recipes: []recipe,
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

func (s *Server) handleGetAllUsers(w http.ResponseWriter, r *http.Request) error {

	users, err := getAllUsers(s.db)
	if err != nil {
		log.Printf("get all users err: %v", err)
	}

	return writeJson(w, http.StatusOK, users)
}

func (s *Server) handleAddUser(w http.ResponseWriter, r *http.Request) error {

	var u User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}

	log.Printf("User - %v", u)
	log.Println("Requesting user checks - username pass")
	user, pass := checkUsernameAndPass(s.db, u.Username, u.Password)
	log.Printf("User: %v\n pass: %v", user, pass)

	if user && pass {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			log.Fatal("Not able to hashpassword")
		}

		u.Password = string(hash)

		err = insertUser(s.db, &u)
		if err != nil {
			log.Printf("error adding user to db - user & pass: %v", err)
		}
		return writeJson(w, http.StatusOK, u)
	} else {
		return writeJson(w, http.StatusBadRequest, nil)
	}

}

func (s *Server) handleUserUpdate(w http.ResponseWriter, r *http.Request) error {
	np := "SuccessfullyUpdatedPass1"
	username := "testuser2"

	updateUser(s.db, username, np)
	return writeJson(w, http.StatusOK, nil)
}

func (s *Server) handleUserDelete(w http.ResponseWriter, r *http.Request) error {
	username := "testuser2"

	deleteUser(s.db, username)
	return writeJson(w, http.StatusOK, nil)
}
