package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"io"
	"net/http"
	"os"
)

const OnStartText = "Pergunte algo para o ChatGPT."

func OnStart(c tele.Context) error {
	return c.Send(OnStartText)
}

type GPTHandle struct{}

func (h *GPTHandle) AskGPT(c tele.Context) error {
	textRequest := TextRequest{
		Model:       "text-davinci-003",
		Prompt:      c.Text(),
		Temperature: 0,
		MaxTokens:   1028,
	}

	postBody, err := json.Marshal(textRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	key := os.Getenv("CHATGPT_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(resBody))

	textResponse := TextResponse{}
	err = json.Unmarshal(resBody, &textResponse)
	if err != nil {
		return err
	}

	result := fmt.Sprintf("ChatGPT: %s", textResponse.Choices[0].Text)
	return c.Send(result)

}

type TextRequest struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	Temperature int64  `json:"temperature"`
	MaxTokens   int64  `json:"max_tokens"`
}

type TextResponse struct {
	Id      string                `json:"id"`
	Object  string                `json:"object"`
	Created int64                 `json:"created"`
	Choices []TextResponseChoices `json:"choices"`
}

type TextResponseChoices struct {
	Text     string `json:"text"`
	Index    int64  `json:"index"`
	LogProbs int64  `json:"logprobs"`
}
