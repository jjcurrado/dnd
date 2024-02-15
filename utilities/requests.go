package utilities

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func ParseOptions(c echo.Context) Options {

	options := Options{
		Feats:      c.FormValue("feats") == "true",
		Spells:     c.FormValue("spells") == "true",
		Appearance: c.FormValue("appearance") == "true",
	}

	return options
}

func ParsePrompt(c echo.Context) (string, error) {

	prompt := c.FormValue("prompt")

	if prompt == "" {
		return "", errors.New("prompt is required")
	}
	return prompt, nil
}
