package server

import (
	"encoding/json"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/dto"
	"net/http"
)

func (s *Server) createActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		s.log.Error("Method not allowed: ", r.Method)
	}

	ctx := r.Context()

	var actor dto.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.log.Error(err)
	}

	err = s.svc.CreateActor(ctx, actor)
	if err != nil {
		s.log.Error(err)
	}
}

func (s *Server) updateActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	ctx := r.Context()

	var actor dto.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.log.Error(err)
	}

	err = s.svc.UpdateActor(ctx, actor)
	if err != nil {
		s.log.Error(err)
	}
}
