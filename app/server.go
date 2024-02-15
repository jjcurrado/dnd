package app

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type server struct {
	e   *echo.Echo
	db  *sql.DB
	ai  *AIservice
	dnd *DNDservice
}
