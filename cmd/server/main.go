package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/sethjback/go-101/pkg/api"
	"github.com/sethjback/go-101/pkg/data"
)

func main() {
	slog.Info("starting")

	s := data.NewStore()

	// add a default admin user
	err := s.SaveUser(context.Background(), &data.User{ID: "admin", Name: "Admin", Admin: true})
	if err != nil {
		slog.Error("could not create user", "error", err)
		os.Exit(1)
	}

	srv := api.New(s)
	errs := srv.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errs:
		slog.Error("could not start server", "error", err)
		os.Exit(1)
	case <-sig:
		// stop gracefully
		slog.Info("stopping")

		err = srv.Stop()
		if err != nil {
			slog.Warn("error stopping server", "error", err)
		}
	}

	slog.Info("that's a wrap!")
}
