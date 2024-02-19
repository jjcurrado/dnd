package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type server struct {
	e *echo.Echo

	db     *sql.DB
	spells *spellTable

	ai  *AIservice
	dnd *DNDservice
}

func NewServer() *server {

	s := &server{}
	s.e = echo.New()

	var err error

	s.db, err = newDB()
	if err != nil {
		panic(err)
	}
	s.spells = newSpellTable(s.db)

	s.ai, err = newAIService()
	if err != nil {
		panic(err)
	}

	s.dnd = newDNDService(s.ai, s.spells)

	s.e.Static("/static", "static")
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
