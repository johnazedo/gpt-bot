package controllers

import (
	"github.com/johnazedo/gpt-bot/src/config"
	"github.com/johnazedo/gpt-bot/src/domain"
	tele "gopkg.in/telebot.v3"
)

func OnStart(c tele.Context) error {
	return c.Send(config.OnStartMessage)
}

type GetTextAnswerRepository interface {
	GetTextAnswer(prompt string) (*domain.TextResponse, error)
}

type Handle struct {
	Repository GetTextAnswerRepository
}

func (h *Handle) Ask(c tele.Context) error {
	resp, err := h.Repository.GetTextAnswer(c.Text())
	if err != nil {
		return err
	}
	return c.Send(resp.Choices[0].Text)
}
