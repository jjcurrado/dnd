package app

import (
	"dnd/templates"
	"dnd/utilities"
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

// Middleware to log the host and method of incoming requests to the given handler
func (s *server) handleLogging(callback echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("----------------------\nHOST : %v\nMETHOD : %v\n----------------------\n", c.Request().URL, c.Request().Method)
		return callback(c)
	}
}

// Create a character sheet from form data and render the HTML to the response
func (s *server) handleCharacterCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		options := utilities.ParseOptions(c)
		prompt, err := utilities.ParsePrompt(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		character := s.dnd.CreateCharacter(prompt, options)
		return templates.Sheet(character).Render(c.Request().Context(), c.Response())
	}
}

// Gives the landing page for the website : index.html
func (s *server) handleHomePage() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		return tmpl.Execute(c.Response().Writer, nil)
	}
}

// Gives the current welcome message : welcome.html
func (s *server) handleWelcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl := template.Must(template.ParseFiles("templates/welcome.html"))
		return tmpl.Execute(c.Response().Writer, nil)
	}
}

// Gives the form with text input and submission for creating a character: prompt.html
func (s *server) handlePrompt() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl := template.Must(template.ParseFiles("templates/prompt.html"))
		return tmpl.Execute(c.Response().Writer, nil)
	}
}

func (s *server) handleSheetsView() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl := template.Must(template.ParseFiles("templates/sheets.html"))
		return tmpl.Execute(c.Response().Writer, nil)
	}
}

func (s *server) handleSheetView() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl := template.Must(template.ParseFiles("templates/charactersheet.html"))
		return tmpl.Execute(c.Response().Writer, nil)
	}
}
