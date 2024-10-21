package router

import (
	"log"
	"net/http"

	"github.com/andjrue/recipe-website-v2/internal/auth"
	"github.com/andjrue/recipe-website-v2/internal/db"
	"github.com/andjrue/recipe-website-v2/internal/recipes"
	"github.com/andjrue/recipe-website-v2/internal/structs"
	"github.com/andjrue/recipe-website-v2/internal/users"
	"github.com/gorilla/pat"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)



func NewServer(addr string, db *mongo.Client) *structs.Server {
	return &structs.Server{
		Addr: addr,
		Db:   db,
	}
}




func Run(s *structs.Server) {
	router := pat.New()

	// BACKEND USER INFORMATION

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			err := users.HandleGetAllUsers(s, w, r) 
			if err != nil {
				db.WriteJson(w, http.StatusBadRequest, err)
			}
		} else if r.Method == "POST" {
			err := users.HandleAddUser(s, w, r)
			if err != nil {
				db.WriteJson(w, http.StatusBadRequest, err)
			}
		} else {
			db.WriteJson(w, http.StatusNotFound, nil)
		}
	})

	router.HandleFunc("/users/{username}", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Request received")
		if r.Method == "GET" {
			// The use case is someone clicking into someone's profile I guess.

            // Do we really want this to be an error? I think we'll set it up eventually
            // so that someone can look at someone elses profile. 
            // I'll leave it like this for now I guess. 
            db.WriteJson(w, http.StatusBadRequest, nil)
		} else if r.Method == "PATCH" {
			err := users.HandleUserUpdate(s, w, r)
			if err != nil {
				db.WriteJson(w, http.StatusBadRequest, err)
			}
		} else if r.Method == "DELETE" { // We want this here and not on Users, because we don't want someone deleting someone else.
			err := users.HandleUserDelete(s, w, r)
			if err != nil {
				db.WriteJson(w, http.StatusBadRequest, err)
			}
		} else {
			log.Println("aaaaaaaa")
			db.WriteJson(w, http.StatusBadRequest, nil)
		}
	})

	// USER SIGN INs.

	router.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// TODO -- Should query the DB, find the username and check if information entered in is correct. Should be prett simple?
			err := auth.CheckUserSignin(s, w, r)
			if err != nil {
                db.WriteJson(w, http.StatusBadRequest, err)
				log.Fatal("User creds do not match")
			}
		} 
	})
    
    // RECIPE ROUTES

    router.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            err := recipes.HandleGetAllRecipes(s) // TODO
            if err != nil {
                db.WriteJson(w, http.StatusBadRequest, err) // This really shouldn't happen, but you never know
                log.Println("No recipes available: %v", err)
            }
        } else if r.Method == "POST" {
            err := recipes.HandleAddRecipe(s, s.Db, "testuser1") // TODO -- Will eventually be from the user struct
            if err != nil {
                db.WriteJson(w, http.StatusBadRequest, err)
                log.Println("Error adding recipe to user: %v", err)
            }
        }
    })

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)

    err := http.ListenAndServe(s.Addr, handler)
	if err != nil {
        log.Printf("error listen and server: %v", err)
    }
    log.Println("Server running")
}

