package router

import (
	"log"
	"net/http"

	"github.com/andjrue/recipe-website-v2/internal/db"
	"github.com/andjrue/recipe-website-v2/internal/structs"
	"github.com/andjrue/recipe-website-v2/internal/users"
	"github.com/gorilla/pat"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server structs.Server



func NewServer(addr string, db *mongo.Client) *Server {
	return &Server{
		Addr: addr,
		Db:   db,
	}
}




func (s *Server) Run() {
	router := pat.New()

	// BACKEND USER INFORMATION

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			err := users.HandleGetAllUsers(w, r)
			if err != nil {
				db.WriteJson(w, http.StatusBadRequest, err)
			}
		} else if r.Method == "POST" {
			err := users.HandleAddUser(w, r)
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
			err := db.HandleUserUpdate(w, r)
			if err != nil {
				db.WriteJson(w, http.StatusBadRequest, err)
			}
		} else if r.Method == "DELETE" { // We want this here and not on Users, because we don't want someone deleting someone else.
			err := db.HandleUserDelete(w, r)
			if err != nil {
				db.WriteJson(w, http.StatusBadRequest, err)
			}
		} else {
			log.Println("aaaaaaaa")
			db.WriteJson(w, http.StatusBadRequest, nil)
		}
	})

	// USER SIGN IN

	router.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// TODO -- Should query the DB, find the username and check if information entered in is correct. Should be prett simple?
			err := db.CheckUserSignin(w, r)
			if err != nil {
                db.WriteJson(w, http.StatusBadRequest, err)
				log.Fatal("User creds do not match")
			}
		} 
	})
    
    // RECIPE ROUTES

    router.HandleFunc("/getAllRecipes", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            err := db.HandleGetAllRecipes(w, r) // TODO
            if err != nil {
                db.WriteJson(w, http.StatusBadRequest, err) // This really shouldn't happen, but you never know
                log.Println("No recipes available: %v", err)
            }
        } else if r.Method == "POST" {
            err := db.HandleAddRecipe(w, r) // TODO
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

