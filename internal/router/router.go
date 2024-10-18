package router

import (
	"encoding/json"
	"log"
	"net/http"
    
    "github.com/rs/cors"
	"github.com/gorilla/pat"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	addr string
	db   *mongo.Client
}




func NewServer(addr string, db *mongo.Client) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}



func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *Server) Run() {
	router := pat.New()

	// BACKEND USER INFORMATION

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			err := s.handleGetAllUsers(w, r)
			if err != nil {
				writeJson(w, http.StatusBadRequest, err)
			}
		} else if r.Method == "POST" {
			err := s.handleAddUser(w, r)
			if err != nil {
				writeJson(w, http.StatusBadRequest, err)
			}
		} else {
			writeJson(w, http.StatusNotFound, nil)
		}
	})

	router.HandleFunc("/users/{username}", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Request received")
		if r.Method == "GET" {
			// The use case is someone clicking into someone's profile I guess.

            // Do we really want this to be an error? I think we'll set it up eventually
            // so that someone can look at someone elses profile. 
            // I'll leave it like this for now I guess. 
            WriteJson(w, http.StatusBadRequest, nil)
		} else if r.Method == "PATCH" {
			err := s.handleUserUpdate(w, r)
			if err != nil {
				WriteJson(w, http.StatusBadRequest, err)
			}
		} else if r.Method == "DELETE" { // We want this here and not on Users, because we don't want someone deleting someone else.
			err := s.handleUserDelete(w, r)
			if err != nil {
				WriteJson(w, http.StatusBadRequest, err)
			}
		} else {
			log.Println("aaaaaaaa")
			WriteJson(w, http.StatusBadRequest, nil)
		}
	})

	// USER SIGN IN

	router.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// TODO -- Should query the DB, find the username and check if information entered in is correct. Should be prett simple?
			err := s.checkUserSignin(w, r)
			if err != nil {
                WriteJson(w, http.StatusBadRequest, err)
				log.Fatal("User creds do not match")
			}
		} 
	})
    
    // RECIPE ROUTES

    router.HandleFunc("/getAllRecipes", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            err := s.HandleGetAllRecipes(w, r) // TODO
            if err != nil {
                WriteJson(w, http.StatusBadRequest, err) // This really shouldn't happen, but you never know
                log.Println("No recipes available: %v", err)
            }
        } else if r.Method == "POST" {
            err := s.HandleAddRecipe(w, r) // TODO
            if err != nil {
                WriteJson(w, http.StatusBadRequest, err)
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

    err := http.ListenAndServe(s.addr, handler)
	if err != nil {
        log.Printf("error listen and server: %v", err)
    }
    log.Println("Server running")
}

