package src

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func OnStart(c tele.Context) error {
	return c.Send(OnStartMessage)
}

type GPTHandle struct{}

func (h *GPTHandle) AskGPT(c tele.Context) error {
	repo := GPTRepository{}
	resp, err := repo.GetGPTTextAnswer(c.Text())
	if err != nil {
		return err
	}

	result := fmt.Sprintf("%s: %s", MessagePrefix, resp.Choices[0].Text)
	return c.Send(result)
}
