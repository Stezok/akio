package pengubot

import (
	"NFTProject/internal/generator/penguin"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PenguBot struct {
	*tgbotapi.BotAPI

	generator *penguin.PenguinGenerator

	callbackHandlers []func(*tgbotapi.Update) bool
	handlers         []func(*tgbotapi.Update) bool

	closeChan chan struct{}
}

func (bot *PenguBot) ReadMessages(updateChan tgbotapi.UpdatesChannel) {
	for {
		select {
		case update := <-updateChan:
			bot.HandleMessage(update)
		case <-bot.closeChan:
			return
		}
	}
}

func (bot *PenguBot) HandleMessage(u tgbotapi.Update) {
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

func (bot *PenguBot) Run() error {

	bot.InitHandlers()

	u := tgbotapi.NewUpdate(0)
	updateChan := bot.GetUpdatesChan(u)
	go bot.ReadMessages(updateChan)

	return nil
}

func NewPenguBot(token string, gen *penguin.PenguinGenerator) (*PenguBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &PenguBot{
		BotAPI:    bot,
		generator: gen,
		closeChan: make(chan struct{}, 1),
	}, nil
}
