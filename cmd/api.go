package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
    addr string
}

func (s *Server) newServer(addr string) *Server {
    return &Server{
        addr: addr,
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
            err := handleAddUser(w, r)
            if err != nil {
                writeJson(w, http.StatusBadRequest, err)
            }
        } else {
            writeJson(w, http.StatusNotFound, nil)
        }
    })

    http.ListenAndServe(s.addr, router)
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) error {
    return nil 
    // Do something
}

func handleAddUser(w http.ResponseWriter, r *http.Request) error {
    return nil
    // Do something
}

