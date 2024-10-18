package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

    "github.com/andjrue/recipe-website-v2/internal/structs"
    "github.com/andjrue/recipe-website-v2/internal/users"
    "github.com/andjrue/recipe-website-v2/internal/router"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Server structs.Server // I see, this is interesting.

/* One really weird thing with Go:
   I can update User no problem, but because a method is being called on Server (right below)
   you need to set up a new type for it that points to the package. 
*/


func (s *Server) CheckUserSignin(w http.ResponseWriter, r *http.Request) error {
	/*  We will need to query the DB with provided credentials from the user.
	        -> Username & Pass

	    Passwords will be hashed at this point, so that's really what were looking for.
	    I'm assuming the hashed passes will come out the same each time, so we can take what they entered, hash it, and then compare what is in the db.

	    This sounds pretty easy, but we'll see.
	*/

    var u users.User
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&u)
    if err != nil {
        http.Error(w, "Invalid request payload - checkUserSignin", http.StatusBadRequest)
    }

	log.Println("Checking username")

	DB := os.Getenv("DB")
	coll := s.Db.Database(DB).Collection("users")
	filter := bson.M{"username": u.Username}

	var result users.User

	err = coll.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            router.WriteJson(w, http.StatusBadRequest, err)

            log.Fatal("No user found with that username")
        }
    }

    err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(u.Password))
    if err != nil {
        router.WriteJson(w, http.StatusBadRequest, err)
        log.Fatal("Passwords do not match")
    }
    log.Println("Check sucsessful. User signed in.")
    return nil
}
