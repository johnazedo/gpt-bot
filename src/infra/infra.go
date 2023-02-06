package infra

import (
	. "github.com/johnazedo/gpt-bot/src/config"
	"github.com/johnazedo/gpt-bot/src/domain"
)

type GPTRepository struct {
	Service ApiService
	Config  *GPTConfig
}

func (r *GPTRepository) GetTextAnswer(prompt string) (*domain.TextResponse, error) {
	// TODO: Create a integration test for this
	request := domain.TextRequest{
		Model:       r.Config.Model,
		Prompt:      prompt,
		Temperature: r.Config.Temperature,
		MaxTokens:   r.Config.MaxTokens,
	}

	response := domain.TextResponse{}

	err := r.Service.Call("POST", r.Config.CompletionsUrl, request, response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
