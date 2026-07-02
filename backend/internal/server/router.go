package server

import (
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/auth"
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/health"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Auth Routes
	mux.HandleFunc("POST /auth/login", auth.Login)
	mux.HandleFunc("POST /auth/register", auth.Register)

	// Health Route(s?)
	mux.HandleFunc("GET /health", health.HealthCheck)

	return mux
}
