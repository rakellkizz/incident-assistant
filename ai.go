package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAISvc struct {
	client *openai.Client
}

func NewOpenAISvc() *OpenAISvc {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		return nil
	}
	return &OpenAISvc{client: openai.NewClient(key)}
}

func (s *OpenAISvc) Summarize(title, description string) (string, error) {
	if s == nil || s.client == nil {
		return "", fmt.Errorf("IA não configurada")
	}
	resp, err := s.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Você é um assistente que resume incidentes de forma objetiva para suporte corporativo."},
			{Role: "user", Content: fmt.Sprintf("Título: %s\nDescrição: %s\nResuma em 1 frase clara para agentes de suporte.", title, description)},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
