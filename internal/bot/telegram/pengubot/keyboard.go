package pengubot

import (
	"NFTProject/internal/meta"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(GenerateMessage),
			tgbotapi.NewKeyboardButton(GeneratePromoMessage),
		),
	)
}

func GetPromoChooseKeyboard() tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton
	var row []tgbotapi.InlineKeyboardButton
	for i, tag := range meta.PromoTags {
		if i%4 == 0 && i != 0 {
			keyboard = append(keyboard, row)
			row = make([]tgbotapi.InlineKeyboardButton, 0)
		}
		row = append(row, tgbotapi.NewInlineKeyboardButtonData(string(tag), tag.GetData()))
	}
	keyboard = append(keyboard, row)
	return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}
