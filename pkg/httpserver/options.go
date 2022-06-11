package httpserver

import (
	"time"
)

type option func(s *server)

// SetAddress - set address to http server
func SetAddress(addr string) option {
	return func(s *server) {
		s.server.Addr = addr
	}
}

// SetReadTimeout - set read timeout to http server
func SetReadTimeout(timeout time.Duration) option {
	return func(s *server) {
		s.server.ReadTimeout = timeout
	}
}

// SetWriteTimeout - set write timeout to http server
func SetWriteTimeout(timeout time.Duration) option {
	return func(s *server) {
		s.server.WriteTimeout = timeout
	}
}

// SetShutdownTimeout - set shutdown timeout to http server
func SetShutdownTimeout(timeout time.Duration) option {
	return func(s *server) {
		s.shutdownTimeout = timeout
	}
}
