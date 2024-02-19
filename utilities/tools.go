package utilities

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

var Functions = []openai.Tool{
	{
		// CreateCharacter
		// makes the base stats and description for a character
		Type: openai.ToolTypeFunction,
		Function: openai.FunctionDefinition{
			Name:        "CreateCharacter",
			Description: "Create a character for the user with the given stats.",
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"Name": {
						Type:        jsonschema.String,
						Description: "Name of the character",
					},
					"Class": {
						Type:        jsonschema.String,
						Description: "the class of the character",
					},
					"Appearance": {
						Type:        jsonschema.String,
						Description: "a description of the physical appearance of the character",
					},
					"SubClass": {
						Type:        jsonschema.String,
						Description: "the subclass of the character",
					},
					"Level": {
						Type:        jsonschema.Integer,
						Description: "the level or challenge rating of the character",
					},
					"HitPoints": {
						Type:        jsonschema.Integer,
						Description: "the maximum hitpoints of the character",
					},
					"Race": {
						Type:        jsonschema.String,
						Description: "the class of the character",
					},
					"Subrace": {
						Type:        jsonschema.String,
						Description: "the sub race of the character, i.e. wood elf versus dark elf",
					},
					"Background": {
						Type:        jsonschema.String,
						Description: "the background of the character",
					},
					"AC": {
						Type:        jsonschema.Integer,
						Description: "the armor class of the character",
					},
					"Speed": {
						Type:        jsonschema.Integer,
						Description: "speed of the character in feet",
					},
					"Strength": {
						Type:        jsonschema.Integer,
						Description: "the strength ability score of the character. This is not the strength modifier but the score.",
					},
					"Intelligence": {
						Type:        jsonschema.Integer,
						Description: "the Intelligence ability score of the character. This is not the intelligence modifier but the score.",
					},
					"Constitution": {
						Type:        jsonschema.Integer,
						Description: "the Consititution ability score of the character. This is not the constituion modifier but the score.",
					},
					"Dexterity": {
						Type:        jsonschema.Integer,
						Description: "the Dexterity ability score of the character. This is not the constituion modifier but the score.",
					},
					"Wisdom": {
						Type:        jsonschema.Integer,
						Description: "the Wisdom ability score of the character. This is not the wisdom modifier but the score.",
					},
					"Charisma": {
						Type:        jsonschema.Integer,
						Description: "the Charisma ability score of the character. This is not the charisma modifier but the score.",
					},
					"Spellcaster": {
						Type:        jsonschema.Boolean,
						Description: "True if this character has the ability to cast spells, false if it does not.",
					},
				},
				Required: []string{
					"Name",
					"Class",
					"Level",
					"Race",
					"Background",
					"AC",
					"Speed",
					"Strength",
					"Intelligence",
					"Constitution",
					"Wisdom",
					"Dexterity",
					"Charisma",
					"Appearance",
					"Spellcaster",
				},
			},
		},
	},
	{
		// CreateSpellList
		// gets an array of strings - names of spells
		Type: openai.ToolTypeFunction,
		Function: openai.FunctionDefinition{
			Name:        "CreateSpellList",
			Description: "creates a list of spells to give to the user",
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"Spells": {
						Type: jsonschema.Array,
						Items: &jsonschema.Definition{
							Type:        jsonschema.String,
							Description: "The name of a spell from fifth edition dungeons and dragons",
						},
					},
				},
				Required: []string{"Spells"},
			},
		},
	},
	{
		Type: openai.ToolTypeFunction,
		Function: openai.FunctionDefinition{
			Name:        "CreateSpellDetails",
			Description: "Creates details of a spell for the user",
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"Name": {
						Type:        jsonschema.String,
						Description: "The name of the spell",
					},
					"Level": {
						Type:        jsonschema.Integer,
						Description: "The level of the spell, level 0 being a cantrip",
					},
					"Range": {
						Type:        jsonschema.String,
						Description: "The Range of the spell",
					},
					"Duration": {
						Type:        jsonschema.String,
						Description: "The duration of the spell, i.e. instantaneous or concentration up to 10 minutes",
					},
					"CastingTime": {
						Type:        jsonschema.String,
						Description: "How long it takes to cast the spell, i.e. 1 action or 1 bonus action",
					},
					"Description": {
						Type:        jsonschema.String,
						Description: "A description of the spell including the damage rolls and saving throws required at base level and upcast levels.",
					},
					"Components": {
						Type:        jsonschema.String,
						Description: "The components used for casting. V for verbal, S for somatic, and M for material",
					},
					"School": {
						Type:        jsonschema.String,
						Description: "The school of magic that the spell pertains to, like Evocation, Illusion, Conjuration, etc.",
					},
				},
				Required: []string{"Name", "Components", "School", "Level", "Range", "Duration", "CastingTime", "Description"},
			},
		},
	},
}
