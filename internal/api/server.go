package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stooezy/dymplex/internal/config"
)

type Server struct {
	Config config.Server
	Http   *http.Server
}

func NewServer(config config.Server) *Server {
	return &Server{
		Config: config,
		Http: &http.Server{
			Addr: config.HttpServer.ListenAddress,
		},
	}
}

func (s *Server) Start() {
	go func() {
		if err := s.Http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("HTTP server error: %v", err))
		}
		slog.Info("Stopped serving new connection")
	}()

	slog.Info(fmt.Sprintf("Server started at %s", s.Config.HttpServer.ListenAddress))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := s.Http.Shutdown(shutdownCtx); err != nil {
		slog.Error("HTTP shutdown error: %v", err)
	}
	slog.Info("Graceful shutdown complete.")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Http.Shutdown(ctx)
}
