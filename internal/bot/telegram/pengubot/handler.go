package pengubot

import (
	"NFTProject/internal/generator/penguin"
	"NFTProject/internal/meta"
	"bytes"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	GenerateMessage       = "♻️ Generate!"
	GeneratePromoMessage  = "♥️ Promo ♦️"
	GenerateCustomMessage = "Custom"
)

func (bot *PenguBot) handleStartCommand(u *tgbotapi.Update) bool {
	if u.Message.Text != "/start" {
		return false
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Welcome!")
	msg.ReplyMarkup = bot.GetMainMenuKeyboard()

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
		msg.ReplyMarkup = bot.GetPromoChooseKeyboard()
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

func (bot *PenguBot) handleGenerateCustomMessage(u *tgbotapi.Update) bool {

	if u.Message.Text != GenerateCustomMessage {
		return false
	}

	part := bot.generator.GetPartFiltred(penguin.Filter{
		Slot: meta.Back,
	})

	rc, err := bot.generator.GetPartReadCloserByFilename(part.FileName)
	if err != nil {
		log.Print(err)
		return true
	}
	defer rc.Close()
	fileReader := tgbotapi.FileReader{
		Name:   "background.png",
		Reader: rc,
	}

	msg := tgbotapi.NewPhoto(u.Message.Chat.ID, fileReader)
	msg.Caption = fmt.Sprintf("Фон: %s\nКоличество: 1", part.FileName)
	msg.ReplyMarkup = bot.GetBackgroundChooseKeyboard()

	_, err = bot.Send(msg)
	if err != nil {
		log.Print(err)
	}
	return true
}

func (bot *PenguBot) handleChangeCountCallback(u *tgbotapi.Update) bool {

	data := strings.Split(u.CallbackData(), ":")

	if data[0] != "count" {
		return false
	}

	arr := strings.Split(u.CallbackQuery.Message.Caption, "\n")
	text := fmt.Sprintf("%s\nКоличество: %s", arr[0], data[1])
	edit := tgbotapi.NewEditMessageCaption(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID, text)
	keyboard := bot.GetBackgroundChooseKeyboard()
	edit.ReplyMarkup = &keyboard
	_, err := bot.Send(edit)
	if err != nil {
		log.Print(err)
	}
	return true
}

func (bot *PenguBot) handleChangeBackgroundCallback(u *tgbotapi.Update) bool {

	data := strings.Split(u.CallbackData(), ":")

	if data[0] != "filename" {
		return false
	}

	rc, err := bot.generator.GetPartReadCloserByFilename(data[1] + ".png")
	if err != nil {
		log.Print(err)
		return true
	}
	defer rc.Close()

	fileReader := tgbotapi.FileReader{
		Name:   "background.png",
		Reader: rc,
	}

	arr := strings.Split(u.CallbackQuery.Message.Caption, "\n")
	text := fmt.Sprintf("Фон: %s\n%s", data[1]+".png", arr[1])

	inputPhoto := tgbotapi.NewInputMediaPhoto(fileReader)
	inputPhoto.Caption = text
	edit := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:    u.CallbackQuery.Message.Chat.ID,
			MessageID: u.CallbackQuery.Message.MessageID,
		},
		Media: inputPhoto,
	}
	keyboard := bot.GetBackgroundChooseKeyboard()
	edit.ReplyMarkup = &keyboard
	_, err = bot.Send(edit)
	if err != nil {
		log.Print(err)
	}
	return true
}

func (bot *PenguBot) handleCommitCustomCallback(u *tgbotapi.Update) bool {

	if u.CallbackData() != "commit_custom" {
		return false
	}

	customGenParams := bot.ParseCustomGenerationData(u.CallbackQuery.Message.Caption)

	edit := tgbotapi.NewEditMessageCaption(
		u.CallbackQuery.Message.Chat.ID,
		u.CallbackQuery.Message.MessageID,
		"Generating...",
	)
	_, err := bot.Send(edit)
	if err != nil {
		log.Print(err)
		return true
	}

	for i := 0; i < customGenParams.Count; i++ {
		var buf bytes.Buffer
		err := bot.generator.GenerateWithBackground(&buf, customGenParams.Filename)
		if err != nil {
			log.Print(err)
			continue
		}
		fileReader := tgbotapi.FileReader{
			Name:   "akiofriend.png",
			Reader: &buf,
		}

		doc := tgbotapi.NewDocument(u.CallbackQuery.Message.Chat.ID, fileReader)
		_, err = bot.Send(doc)
		if err != nil {
			log.Print(err)
		}
	}

	edit = tgbotapi.NewEditMessageCaption(
		u.CallbackQuery.Message.Chat.ID,
		u.CallbackQuery.Message.MessageID,
		"Done!",
	)
	_, err = bot.Send(edit)
	if err != nil {
		log.Print(err)
		return true
	}

	return true
}

func (bot *PenguBot) InitHandlers() {

	bot.callbackHandlers = []func(*tgbotapi.Update) bool{
		bot.handleGeneratePromoCallback,
		bot.handleChangeCountCallback,
		bot.handleChangeBackgroundCallback,
		bot.handleCommitCustomCallback,
	}

	bot.handlers = []func(*tgbotapi.Update) bool{
		bot.handleStartCommand,
		bot.handleGenerateMessage,
		bot.handleGeneratePromoMessage,
		bot.handleGenerateCustomMessage,
	}

}
