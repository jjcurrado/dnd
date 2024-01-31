package server

import (
	ai "dnd/openai"
	"dnd/templates"
	"dnd/utilities"
	"fmt"
	"net/http"
)

func CharacterPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("ERROR PARSING FORM: %v", err)))
		return
	}

	fmt.Println(r.Form)
	options := utilities.Options{
		Feats:      r.Form.Get("feats") == "true",
		Spells:     r.Form.Get("spells") == "true",
		Appearance: r.Form.Get("appearance") == "true",
	}

	prompt := r.Form.Get("prompt")

	if prompt == "" {
		w.Write([]byte("ERROR PARSING FORM: No prompt found"))
		return
	}

	fmt.Printf("Character prompt recieved: %s\n", prompt)
	character := ai.CreateCharacter(prompt, options)

	templates.Sheet(character).Render(r.Context(), w)
}
