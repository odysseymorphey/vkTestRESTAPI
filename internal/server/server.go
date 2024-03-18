package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	router *http.ServeMux
	port   int
	log    *zap.SugaredLogger
	server *http.Server
	svc    FilmaryService
}

func NewServer(log *zap.SugaredLogger, svc FilmaryService, port int) (*Server, error) {
	mux := http.NewServeMux()

	server := &Server{
		router: mux,
		port:   port,
		log:    log,
		svc:    svc,
	}

	mux.HandleFunc("/createActor", server.createActor)
	mux.HandleFunc("/updateActor", server.updateActor)
	mux.HandleFunc("/deleteActor", server.deleteActor)
	mux.HandleFunc("/createMovie", server.createMovie)
	mux.HandleFunc("/deleteMovie", server.deleteMovie)

	return server, nil
}

func (s *Server) Run(ctx context.Context) {
	addr := fmt.Sprintf("0.0.0.0:%d", s.port)
	s.log.Infof("Server listening at %s", addr)

	go s.stopProcess(ctx)

	s.server = &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.log.Fatal(err)
	}
}

func (s *Server) stopProcess(ctx context.Context) {
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Error(err)
	}
}
