package app

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

// Middleware to log the host and method of incoming requests to the given handler
func (s *server) logRequests(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Info("Request recieved", "host", c.Request().URL, "method", c.Request().Method)
		return next(c)
	}
}
