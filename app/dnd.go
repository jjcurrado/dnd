package app

import (
	"database/sql"
	ai "dnd/openai"
	util "dnd/utilities"
)

type DNDservice struct {
	ai *AIservice
	db *sql.DB
}

func (dnd *DNDservice) createCharacter(prompt string, options util.Options) util.Character {
	return ai.CreateCharacter(prompt, options)
}
