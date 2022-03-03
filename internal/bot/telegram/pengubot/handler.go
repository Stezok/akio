package pengubot

import (
	"NFTProject/internal/meta"
	"bytes"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	GenerateMessage      = "♻️ Generate!"
	GeneratePromoMessage = "♥️ Promo ♦️"
)

func (bot *PenguBot) handleStartCommand(u *tgbotapi.Update) bool {
	if u.Message.Text != "/start" {
		return false
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Welcome!")
	msg.ReplyMarkup = GetMainMenuKeyboard()

	_, err := bot.Send(msg)
	if err != nil {
		log.Print(err)
	}
	return true
}

func (bot *PenguBot) handleGenerateMessage(u *tgbotapi.Update) bool {
	if u.Message.Text != GenerateMessage {
		return false
	}

	go func() {
		var buf bytes.Buffer
		err := bot.generator.GenerateRandomSingle(&buf)
		if err != nil {
			log.Print(err)
			return
		}

		photo := tgbotapi.FileReader{
			Name:   "PenguAKIO.png",
			Reader: &buf,
		}
		msg := tgbotapi.NewPhoto(u.Message.Chat.ID, photo)
		_, err = bot.Send(msg)
		if err != nil {
			log.Print(err)
		}
	}()
	return true
}

func (bot *PenguBot) handleGeneratePromoMessage(u *tgbotapi.Update) bool {
	if u.Message.Text != GeneratePromoMessage {
		return false
	}

	go func() {
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите промо коллекцию")
		msg.ReplyMarkup = GetPromoChooseKeyboard()
		_, err := bot.Send(msg)
		if err != nil {
			log.Print(err)
		}
	}()

	return true
}

func (bot *PenguBot) handleGeneratePromoCallback(u *tgbotapi.Update) bool {
	data := strings.Split(u.CallbackData(), ":")
	if data[0] != "promotag" {
		return false
	}

	go func() {
		var buf bytes.Buffer
		err := bot.generator.GeneratePromo(&buf, meta.PromoTag(data[1]))
		if err != nil {
			log.Print(err)
			return
		}
		file := tgbotapi.FileReader{
			Name:   "PenguAKIO.png",
			Reader: &buf,
		}
		photo := tgbotapi.NewPhoto(u.CallbackQuery.Message.Chat.ID, file)
		_, err = bot.Send(photo)
		if err != nil {
			log.Print(err)
		}
	}()

	return true
}

func (bot *PenguBot) InitHandlers() {

	bot.callbackHandlers = []func(*tgbotapi.Update) bool{
		bot.handleGeneratePromoCallback,
	}

	bot.handlers = []func(*tgbotapi.Update) bool{
		bot.handleStartCommand,
		bot.handleGenerateMessage,
		bot.handleGeneratePromoMessage,
	}

}
