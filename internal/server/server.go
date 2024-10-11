package server

import (
  "encoding/json"
  "net/http"
)

type Server struct {
  listenAddr string
  // Will need to add the DB here at some point
}

func NewApiServer(l string) *Server {
  return &Server{listenAddr: l}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  return json.NewEncoder(w).Encode(v)
}




