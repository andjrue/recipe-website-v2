package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

func writeJson(w http.ResponseWriter, status int, v any) error {
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
		if r.Method == "GET" {
			// The use case is someone clicking into someone's profile I guess.
			writeJson(w, http.StatusBadRequest, nil)
			// TO DO
		} else if r.Method == "PATCH" {
			err := s.handleUserUpdate(w, r)
			if err != nil {
				writeJson(w, http.StatusBadRequest, err)
			}
		} else if r.Method == "DELETE" { // We want this here and not on Users, because we don't want someone deleting someone else.
			err := s.handleUserDelete(w, r)
			if err != nil {
				writeJson(w, http.StatusBadRequest, err)
			}
		} else {
			log.Println("aaaaaaaa")
			writeJson(w, http.StatusBadRequest, nil)
		}
	})

    // USER SIGN IN 

    router.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            // TODO -- Should query the DB, find the username and check if information entered in is correct. Should be prett simple?
            err := s.checkUserSignin(w, r) 
            if err != nil {
                writeJson(w, http.StatusBadRequest, nil)
                log.Fatal("User creds do not match")
            }
            writeJson(w, http.StatusOK, nil)
        }
    })

	http.ListenAndServe(s.addr, router)
	log.Println("Server running")
}

func (s *Server) handleGetAllUsers(w http.ResponseWriter, r *http.Request) error {

	users, err := getAllUsers(s.db)
	if err != nil {
		log.Printf("get all users err: %v", err)
	}

	return writeJson(w, http.StatusOK, users)
}

func (s *Server) handleAddUser(w http.ResponseWriter, r *http.Request) error {
    u := &User{Email: "testemail4@gmail.com", Username: "testuser4", Password: "testpass4"}
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

		err = insertUser(s.db, u)
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
