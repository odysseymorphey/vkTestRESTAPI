package server

import (
	"encoding/json"
	"fmt"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/dto"
	"net/http"
	"strconv"
)

func (s *Server) mock(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) createActor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		s.log.Error("Method not allowed: ", r.Method)
	}
}

func (s *Server) updateActor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
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

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		s.log.Error("Method not allowed: ", r.Method)
	}
}

func (s *Server) deleteActor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		ctx := r.Context()

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			s.log.Error(err)
		}

		err = s.svc.DeleteActor(ctx, id)
		if err != nil {
			s.log.Error(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		s.log.Error("Method not allowed: ", r.Method)
	}
}

func (s *Server) createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ctx := r.Context()

		var movie dto.Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.log.Error(err)
		}

		fmt.Fprint(w, movie)

		err = s.svc.CreateMovie(ctx, movie)
		if err != nil {

		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	}
}
