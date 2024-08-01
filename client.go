package main

import (
	"os"

	"github.com/gage-technologies/mistral-go"
)

// MistralClient creates and returns a new MistralClient instance.
func MistralClient() *mistral.MistralClient {
	apiKey := os.Getenv("MISTRAL_API_KEY")
	if apiKey == "" {
		apiKey = "your-api-key" // Fallback if not using environment variables
	}
	return mistral.NewMistralClientDefault(apiKey)
}
