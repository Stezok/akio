package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramHandler func(*tgbotapi.Update) bool

type Handler interface {
	GetCallbackHandlers() []TelegramHandler
	GetHandlers() []TelegramHandler
	BindBot(*tgbotapi.BotAPI)
}

type TelegramBot struct {
	*tgbotapi.BotAPI

	handler Handler

	closeChan chan struct{}
}

func (bot *TelegramBot) BindHandler(handler Handler) {
	handler.BindBot(bot.BotAPI)
	bot.handler = handler
}

func (bot *TelegramBot) ReadMessages(updateChan tgbotapi.UpdatesChannel) {
	for {
		select {
		case update := <-updateChan:
			bot.HandleMessage(update)
		case <-bot.closeChan:
			return
		}
	}
}

func (bot *TelegramBot) HandleMessage(u tgbotapi.Update) {
	if u.CallbackQuery != nil {
		for _, f := range bot.handler.GetCallbackHandlers() {
			if f(&u) {
				return
			}
		}
	} else {
		for _, f := range bot.handler.GetHandlers() {
			if f(&u) {
				return
			}
		}
	}
}

func (bot *TelegramBot) Run() error {

	u := tgbotapi.NewUpdate(0)
	updateChan := bot.GetUpdatesChan(u)
	go bot.ReadMessages(updateChan)

	return nil
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		BotAPI:    bot,
		closeChan: make(chan struct{}, 1),
	}, nil
}
