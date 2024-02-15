package app

import (
	"dnd/templates"
	"dnd/utilities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) handleCharacterCreate() echo.HandlerFunc {

	return func(c echo.Context) error {

		options := utilities.ParseOptions(c)
		prompt, err := utilities.ParsePrompt(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		character := s.dnd.createCharacter(prompt, options)
		return templates.Sheet(character).Render(c.Request().Context(), c.Response())
	}
}
