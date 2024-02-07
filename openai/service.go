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

//TODO:
// - use the ai in a more generic sense (i.e. have a function check which function calls its lookking for)

// CreateCharacter
// creates a character from the given description
//
// TODO:
// - make the options dynamic for request (which parameters are required)
func CreateCharacter(description string, options util.Options) util.Character {

	char := RequestCharacterStats(description)

	// get a spell list
	if char.Spellcaster && options.Spells {

		l := RequestSpellList(char)
		fmt.Printf("\n%v\n", l)

		s := GetAllSpells(l)

		for _, spell := range s {
			util.PrintSpell(spell)
		}
	}
	return char
}

// RequestCharacterStats
// Create the base character stats including: name, class, subclass, hp, ac, and ability scores
// TODO: Do some of these requests in parallel?
func RequestCharacterStats(prompt string) util.Character {
	res := sendRequest(prompt)

	char := GetCharacterResponse(res)
	return char
}

// RequestSpellList
// Given a character - make a list of spell names that they can cast
func RequestSpellList(char util.Character) []string {

	spells, cantrips := util.PreparedSpells(char)
	prompt := fmt.Sprintf("Create a list of %v spells and %v cantrips for a level %v %v . Only include spells that can be cast.", spells, cantrips, char.Level, char.Class)

	res := sendRequest(prompt)
	l := GetSpellListResponse(res)
	return l
}

// RequestSpellDetails
// Given the name of a spell, get its details including: name, level, range, duration, and description
func RequestSpellDetails(spell string) util.Spell {

	//TODO: check a database if we've already generated the spell details instead of prompting chat GPT for it
	// 		also check the official SRD api ?
	prompt := fmt.Sprintf("Give me the details for the spell %v", spell)
	fmt.Printf("Creating spell details:\n%v\n", prompt)

	res := sendRequest(prompt)
	s := GetSpellDetailsResponse(res)

	return s
}

// GetCharacterResponse
// Extract character data from the openai response
func GetCharacterResponse(res openai.ChatCompletionResponse) util.Character {
	msg := res.Choices[0].Message
	char := util.Character{}
	json.Unmarshal([]byte(msg.ToolCalls[0].Function.Arguments), &char)
	return char
}

// GetSpellListResponse
// Extract spell name list from the openai response
func GetSpellListResponse(res openai.ChatCompletionResponse) []string {
	msg := res.Choices[0].Message
	spell := util.SpellList{}
	json.Unmarshal([]byte(msg.ToolCalls[0].Function.Arguments), &spell)
	return spell.Spells
}

// GetSpellDetailsResponse
// Extract spell detail data from the openai response
func GetSpellDetailsResponse(res openai.ChatCompletionResponse) util.Spell {
	msg := res.Choices[0].Message
	spell := util.Spell{}
	json.Unmarshal([]byte(msg.ToolCalls[0].Function.Arguments), &spell)
	return spell
}

// GetAllSpells
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

// checkResponse
// check for errors in the openai response
func checkResponse(res openai.ChatCompletionResponse, err error) {
	if err != nil {
		log.Fatalf("API error: %v", err)
	}
	if len(res.Choices) < 1 {
		log.Fatalf("Choices Empty!")
	}
}

// sendRequest
// send a request to the openai API and check for errors
// return the response.
// TODO:
// - add rate limiter
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
