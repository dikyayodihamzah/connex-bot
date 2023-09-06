package config

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	// botID    = os.Getenv("TELEGRAM_BOT_ID")
	botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
)

func NewBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	return bot
}
