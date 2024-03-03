package app

import (
	database "dnd/db"
	"dnd/services"
	"log"
	"net/http"

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
	if err := s.e.Start(port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func (s *server) Close() {
	s.db.Close()
}
