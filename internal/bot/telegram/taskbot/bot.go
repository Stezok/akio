package taskbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TaskBot struct {
	Team  []int64
	Token string
	*tgbotapi.BotAPI
}

func (bot *TaskBot) Run() error {
	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		err := bot.handleMessage(&update)
		if err != nil {
			return err
		}
	}

	return nil
}
