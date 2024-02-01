package utilities

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

var AIClient = CreateClient()

func CreateClient() *openai.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var client = openai.NewClient(os.Getenv("API_KEY"))

	return client
}
