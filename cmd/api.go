package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *Server) Run() {
	router := mux.NewRouter()

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
	var user, pass bool
	u := newUser("testemail4@gmail.com", "testuser4", "testpass4")
    log.Println("Requesting user checks - username pass")
	user, pass = checkUsernameAndPass(s.db, u.Username, u.Password)

	if user && pass {
		err := insertUser(s.db, u)
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
