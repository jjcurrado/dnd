package app

import (
	"dnd/templates"
	"dnd/utilities"
	"log"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

// Create a character sheet from form data and render the HTML to the response
func (s *server) handleCharacterCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		options := utilities.ParseOptions(c)
		prompt, err := utilities.ParsePrompt(c)
		if err != nil {
			slog.Error("Error parsing prompts", "msg", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		slog.Info("Creating Character...", "prompt", prompt)
		character, err := s.dnd.CreateCharacter(prompt, options)
		if err != nil {
			slog.Error("Error creating character", "msg", err.Error())
			return s.renderError(err.Error(), c)
		}
		slog.Info("Character created.", "character", character)

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

func (s *server) handleFileUploadView() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl := template.Must(template.ParseFiles("templates/fileupload.html"))
		return tmpl.Execute(c.Response().Writer, nil)
	}
}

func (s *server) handleFileUpload() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print(c.FormParams())
		//s, e := c.FormFile("f")
		//if e != nil {
		//panic(e)
		//}
		//log.Print(s.Filename)
		return nil
	}
}

func (s *server) renderError(message string, c echo.Context) error {
	return templates.Error(message).Render(c.Request().Context(), c.Response())
}
