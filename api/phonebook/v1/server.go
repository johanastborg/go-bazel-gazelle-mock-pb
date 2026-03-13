package v1

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	db Database
}
func NewServer(db Database) *Server {
	return &Server{db: db}
}

func (s *Server) PersonHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}

	p, err := s.db.GetPerson(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
