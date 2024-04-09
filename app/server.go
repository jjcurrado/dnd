package app

import (
	database "dnd/db"
	"dnd/services"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// HTTP server representing entire application
type server struct {
	e   *echo.Echo
	db  *database.DB
	ai  *services.AIservice
	dnd *services.DNDservice
}

// Instantiate a new server and initialize services and routing
func NewServer(e *echo.Echo, ai *services.AIservice, db *database.DB) *server {
	s := &server{}
	s.e = e
	s.db = db
	s.ai = ai
	s.dnd = services.NewDNDService(ai, db)
	s.e.Static("/static", "static")
	s.e.Use(s.logRequests)
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.e.ServeHTTP(w, r)
}

func (s *server) Start(port string) {
	slog.Info("Starting Server...")
	if err := s.e.Start(port); err != http.ErrServerClosed {
		slog.Error("Error starting server", "msg", err.Error())
		os.Exit(1)
	}
	slog.Info("Started.")
}

func (s *server) Close() {
	slog.Info("Closing server...")
	s.db.Close()
	slog.Info("Closed.")
}
