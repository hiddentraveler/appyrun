package handlers

import "appyrun/internal/server"

type Handlers struct {
	server *server.Server
}

func NewHandlers(s *server.Server) *Handlers {
	return &Handlers{server: s}
}
