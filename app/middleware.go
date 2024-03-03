package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// Middleware to log the host and method of incoming requests to the given handler
func (s *server) logRequests(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("----------------------\nHOST : %v\nMETHOD : %v\n----------------------\n", c.Request().URL, c.Request().Method)
		return next(c)
	}
}
