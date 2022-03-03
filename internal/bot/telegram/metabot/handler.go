package metabot

import (
	"NFTProject/internal/meta"
	"bytes"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *MetaBot) HandleFragmentSent(u *tgbotapi.Update) bool {

	sizes := u.Message.Photo
	if len(sizes) == 0 {
		return false
	}
	photoConfig := tgbotapi.NewPhoto(u.Message.From.ID, tgbotapi.FileID(sizes[0].FileID))
	photoConfig.Caption = fmt.Sprintf("Текущие атрибуты\nИмя: %s\n\nВыберите слот фрагмента", u.Message.Caption)
	photoConfig.BaseChat.ReplyMarkup = GetSlotKeyboard()

	_, err := bot.Send(photoConfig)
	if err != nil {
		log.Print(err)
	}

	return true
}

func (bot *MetaBot) HandleSlotPickCallback(u *tgbotapi.Update) bool {
	data := strings.Split(u.CallbackData(), ":")

	if data[0] != "slot" {
		return false
	}

	attributes := strings.Split(u.CallbackQuery.Message.Caption, "\n\n")[0]
	newCaption := fmt.Sprintf("%s\nСлот: %s\n\nВыберите цвет фрагмента", attributes, data[1])
	keyboard := GetColorKeyboard()
	edit := tgbotapi.NewEditMessageCaption(
		u.CallbackQuery.Message.Chat.ID,
		u.CallbackQuery.Message.MessageID,
		newCaption,
	)
	edit.BaseEdit.ReplyMarkup = &keyboard

	_, err := bot.Send(edit)
	if err != nil {
		log.Print(err)
	}
	return true
}

func (bot *MetaBot) HandleColorPickCallback(u *tgbotapi.Update) bool {

	data := strings.Split(u.CallbackData(), ":")

	if data[0] != "color" {
		return false
	}

	attributes := strings.Split(u.CallbackQuery.Message.Caption, "\n\n")[0]
	newCaption := fmt.Sprintf("%s\nЦвет: %s\n\nВыберите редкость фрагмента", attributes, data[1])
	keyboard := GetRarityKeyboard()
	edit := tgbotapi.NewEditMessageCaption(
		u.CallbackQuery.Message.Chat.ID,
		u.CallbackQuery.Message.MessageID,
		newCaption,
	)
	edit.BaseEdit.ReplyMarkup = &keyboard

	_, err := bot.Send(edit)
	if err != nil {
		log.Print(err)
	}
	return true
}

func (bot *MetaBot) HandleRarityPickCallback(u *tgbotapi.Update) bool {

	data := strings.Split(u.CallbackData(), ":")

	if data[0] != "rarity" {
		return false
	}

	if data[1] == string(meta.Promo) {

		attributes := strings.Split(u.CallbackQuery.Message.Caption, "\n\n")[0]
		newCaption := fmt.Sprintf("%s\nРедкость: %s\n\nВыберите промо-коллекцию", attributes, data[1])
		keyboard := GetPromoTagKeyboard()
		edit := tgbotapi.NewEditMessageCaption(
			u.CallbackQuery.Message.Chat.ID,
			u.CallbackQuery.Message.MessageID,
			newCaption,
		)
		edit.BaseEdit.ReplyMarkup = &keyboard

		_, err := bot.Send(edit)
		if err != nil {
			log.Print(err)
		}

	} else {
		attributes := strings.Split(u.CallbackQuery.Message.Caption, "\n\n")[0]
		newCaption := fmt.Sprintf("%s\nРедкость: %s", attributes, data[1])
		err := bot.AcceptMetadata(u, newCaption)
		if err != nil {
			log.Print(err)
		}
	}

	return true
}

func (bot *MetaBot) HandlePromoTagPickCallback(u *tgbotapi.Update) bool {

	data := strings.Split(u.CallbackData(), ":")

	if data[0] != "promotag" {
		return false
	}

	attributes := strings.Split(u.CallbackQuery.Message.Caption, "\n\n")[0]
	newCaption := fmt.Sprintf("%s\nПромо-коллекция: %s", attributes, data[1])
	err := bot.AcceptMetadata(u, newCaption)
	if err != nil {
		log.Print(err)
	}
	return true
}

func (bot *MetaBot) HandleAcceptMetadata(u *tgbotapi.Update) bool {

	if u.CallbackData() != "accept" {
		return false
	}

	fragmentMeta := ParseMetaAttributes(u.CallbackQuery.Message.Caption)
	buffer := &bytes.Buffer{}
	err := fragmentMeta.WriteJson(buffer)
	if err != nil {
		log.Print(err)
		return true
	}

	file := tgbotapi.FileReader{
		Name:   fragmentMeta.FileName,
		Reader: buffer,
	}

	doc := tgbotapi.NewDocument(u.CallbackQuery.Message.Chat.ID, file)
	_, err = bot.Send(doc)
	if err != nil {
		log.Print(err)
	}

	return true
}

func (bot *MetaBot) InitHandlers() {

	bot.callbackHandlers = []func(*tgbotapi.Update) bool{
		bot.HandleSlotPickCallback,
		bot.HandleColorPickCallback,
		bot.HandleRarityPickCallback,
		bot.HandlePromoTagPickCallback,
		bot.HandleAcceptMetadata,
	}

	bot.handlers = []func(*tgbotapi.Update) bool{
		bot.HandleFragmentSent,
	}

}
