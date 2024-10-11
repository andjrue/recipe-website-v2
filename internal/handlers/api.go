package handlers

import (
	"net/http"


  "github.com/gorilla/mux"
	"github.com/andjrue/recipe-website-v2/internal/server"
)

func (s *Server) Run() {
	r := mux.NewRouter()

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		err := handleUser(w, r)
		if err != nil {
			writeJson(w, http.StatusBadRequest, err)
		}
	})

}
