package metabot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MetaBot struct {
	*tgbotapi.BotAPI

	callbackHandlers []func(*tgbotapi.Update) bool
	handlers         []func(*tgbotapi.Update) bool

	closeChan chan struct{}
}

func (bot *MetaBot) ReadMessages(updateChan tgbotapi.UpdatesChannel) {
	for {
		select {
		case update := <-updateChan:
			bot.HandleMessage(update)
		case <-bot.closeChan:
			return
		}
	}
}

func (bot *MetaBot) HandleMessage(u tgbotapi.Update) {
	if u.CallbackQuery != nil {
		for _, f := range bot.callbackHandlers {
			if f(&u) {
				return
			}
		}
	} else {
		for _, f := range bot.handlers {
			if f(&u) {
				return
			}
		}
	}
}

func (bot *MetaBot) Run() error {

	bot.InitHandlers()

	u := tgbotapi.NewUpdate(0)
	updateChan := bot.GetUpdatesChan(u)
	go bot.ReadMessages(updateChan)

	return nil
}

func NewMetaBot(token string) (*MetaBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &MetaBot{
		BotAPI:    bot,
		closeChan: make(chan struct{}, 1),
	}, nil
}
