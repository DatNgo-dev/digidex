package api

import (
	"encoding/json"
	"net/http"

	"github.com/DatNgo-dev/digidex/util"
)

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(1)

	_ = util.Round2Dec(10.34434)

	json.NewEncoder(w).Encode(user)
}

// func (s *Server) handleDeleteUserByID(w http.ResponseWriter, r *http.Request) {
// 	user := s.store.Get(1)

// 	json.NewEncoder(w).Encode(user)
// }
