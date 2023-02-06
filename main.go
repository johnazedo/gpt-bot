package main

import (
	"github.com/johnazedo/gpt-bot/src/config"
	"github.com/johnazedo/gpt-bot/src/controllers"
	"github.com/johnazedo/gpt-bot/src/infra"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	_ = godotenv.Load()
	gpt := controllers.Handle{
		Repository: &infra.GPTRepository{
			Config: config.DefaultGPTConfig,
			Service: &infra.ApiServiceImpl{
				ApiKey: os.Getenv("CHATGPT_KEY"),
			},
		},
	}

	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_KEY"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", controllers.OnStart)
	b.Handle(tele.OnText, gpt.Ask)

	b.Start()
}
