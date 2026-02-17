package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sethjback/go-101/pkg/data"
)

type Server struct {
	store data.Store
	r     chi.Router
	srv   *http.Server
}

func New(store data.Store) *Server {
	srv := &Server{store: store}

	srv.r = chi.NewRouter()
	srv.r.Use(middleware.Logger)
	srv.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	srv.srv = &http.Server{
		Addr:    ":8080",
		Handler: srv.r,
	}

	return srv
}

func (s *Server) Start() chan error {
	errs := make(chan error)
	go func(e chan error) {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errs <- err
		}
	}(errs)

	return errs
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
