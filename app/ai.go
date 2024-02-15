package app

import (
	"github.com/sashabaranov/go-openai"
)

type AIservice struct {
	client *openai.Client
	tools  []openai.Tool
}

func (ai *AIservice) sendRequest() error {

	return nil
}
