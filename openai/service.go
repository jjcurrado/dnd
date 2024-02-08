package ai

import (
	"context"
	util "dnd/utilities"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	openai "github.com/sashabaranov/go-openai"
)

// creates a character from the given description
func CreateCharacter(description string, options util.Options) util.Character {
	char := RequestCharacterStats(description)

	// get a spell list
	if char.Spellcaster && options.Spells {
		l := RequestSpellList(char)
		s := GetAllSpells(l)
		for _, spell := range s {
			util.PrintSpell(spell)
		}
	}

	return char
}

// Create the base character stats including: name, class, subclass, hp, ac, and ability scores
func RequestCharacterStats(prompt string) util.Character {
	res := sendRequest(prompt)
	char := util.Character{}
	err := GetAIResponse(res, &char)

	if err != nil {
		log.Fatal(err)
	}

	util.PrintCharacter(char)
	return char
}

// Given a character - make a list of spell names that they can cast
func RequestSpellList(char util.Character) []string {

	spells, cantrips := util.PreparedSpells(char)
	prompt := fmt.Sprintf("Create a list of %v spells and %v cantrips for a level %v %v . Only include spells that can be cast by a character of this level and class.", spells, cantrips, char.Level, char.Class)

	res := sendRequest(prompt)
	l := util.SpellList{}
	err := GetAIResponse(res, &l)

	if err != nil {
		log.Fatal(err)
	}

	return l.Spells
}

// Given the name of a spell, get its details including: name, level, range, duration, and description
func RequestSpellDetails(spell string) util.Spell {

	prompt := fmt.Sprintf("Give me the details for the spell %v", spell)
	res := sendRequest(prompt)
	s := util.Spell{}

	err := GetAIResponse(res, &s)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

func GetAIResponse(res openai.ChatCompletionResponse, out any) error {
	msg := res.Choices[0].Message
	return json.Unmarshal([]byte(msg.ToolCalls[0].Function.Arguments), &out)
}

// Loop through the list of spell names and request the details for each in parallel
func GetAllSpells(SpellNames []string) []util.Spell {
	details := make([]util.Spell, len(SpellNames))
	var wg sync.WaitGroup

	for i := 0; i < len(SpellNames); i++ {
		wg.Add(1)
		go func(SpellName string, index int) {
			defer wg.Done()
			spell, err := util.FindSpell(SpellName)

			if err == nil && spell.Name != "" {
				details[index] = spell
			} else {
				spell = RequestSpellDetails(SpellName)
				util.InsertSpell(spell)
				details[index] = spell
			}
		}(SpellNames[i], i)
	}
	wg.Wait()

	return details
}

// check for errors in the openai response
func checkResponse(res openai.ChatCompletionResponse, err error) {
	if err != nil {
		log.Fatalf("API error: %v", err)
	}
	if len(res.Choices) < 1 {
		log.Fatalf("Choices Empty!")
	}
}

// send a request to the openai API and check for errors
// return the response.
func sendRequest(prompt string) openai.ChatCompletionResponse {
	// prime the dialogue for the ai
	dialog := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a helpful assistant to a Dungeon Master for fifth edition Dungeons and Dragons.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	ctx := context.Background()

	res, err := util.AIClient.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4TurboPreview,
			Messages: dialog,
			Tools:    Functions,
		},
	)

	checkResponse(res, err)
	return res
}
