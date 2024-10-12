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
    db *mongo.Client
}

func NewServer(addr string, db *mongo.Client) *Server {
    return &Server{
        addr: addr,
        db: db,
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
            err := handleGetUsers(w, r)
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

    http.ListenAndServe(s.addr, router)
    log.Println("Server running")
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) error {
    return nil 
    // Do something
}

func (s *Server) handleAddUser(w http.ResponseWriter, r *http.Request) error {
    u := newUser("testemail1@gmail.com", "testuser1", "testpass1")
    insertUser(s.db, u)
    return writeJson(w, http.StatusOK, u)
}

