package httpserver

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultAddr            = ":80"
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeout = 5 * time.Second
)

type server struct {
	server          *http.Server
	logger          *logrus.Logger
	shutdownTimeout time.Duration
}

func NewServer(handler http.Handler, log *logrus.Logger, options ...option) *server {
	s := &server{
		server: &http.Server{
			Addr:         defaultAddr,
			Handler:      handler,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		},
		logger:          log,
		shutdownTimeout: defaultShutdownTimeout,
	}

	// set options
	for _, opt := range options {
		opt(s)
	}

	return s
}

func (s *server) Run() error {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-exit
		s.Stop()
	}()

	s.logger.Info("starting server on port" + s.server.Addr)
	err := s.server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (s *server) Stop() error {
	s.logger.Info("stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
