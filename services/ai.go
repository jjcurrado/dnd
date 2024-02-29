package services

import (
	"context"
	util "dnd/utilities"
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

// A service to communicate with OpenAI using tool call functionality
type AIservice struct {
	client *openai.Client
	tools  []openai.Tool
}

func NewAIService() (*AIservice, error) {
	ai := &AIservice{}
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	ai.client = openai.NewClient(os.Getenv("API_KEY"))
	ai.tools = util.Functions
	return ai, nil
}

// Send a request to the openai API and check for errors
// return the response.
func (ai *AIservice) sendRequest(prompt string) openai.ChatCompletionResponse {
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
	res, err := ai.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4TurboPreview,
			Messages: dialog,
			Tools:    ai.tools,
		},
	)

	ai.checkAIResponse(res, err)
	return res
}

// check the response from OpenAI to get the function call arguments
// TODO: Handle empty choices, handle multiple choices (i.e. multiple character creations)
// probably need a custom interface of some sort here to handle different possible outcomes
func (ai *AIservice) getAIResponse(res openai.ChatCompletionResponse, out any) error {
	msg := res.Choices[0].Message
	return json.Unmarshal([]byte(msg.ToolCalls[0].Function.Arguments), &out)
}

// check for errors or nil response
func (ai *AIservice) checkAIResponse(res openai.ChatCompletionResponse, err error) {
	if err != nil {
		panic(err)
	}
	if len(res.Choices) < 1 {
		panic("Choices Empty!")
	}
}
