package main

import (
	"dnd/app"
	database "dnd/db"
	"dnd/services"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Failure to initialize database.\nError Message:%v", err)
	}

	ai, err := services.NewAIService()
	if err != nil {
		log.Fatalf("Failure to initialize OpenAI connection.\nError Message:%v", err)
	}

	e := echo.New()
	s := app.NewServer(e, ai, db)

	s.Start(":8000")
	s.Close()
}
