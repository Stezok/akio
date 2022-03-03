package pengubot

import (
	"NFTProject/internal/meta"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *PenguBot) GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(GenerateMessage),
			tgbotapi.NewKeyboardButton(GeneratePromoMessage),
			tgbotapi.NewKeyboardButton(GenerateCustomMessage),
		),
	)
}

func (bot *PenguBot) GetPromoChooseKeyboard() tgbotapi.InlineKeyboardMarkup {
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

func (bot *PenguBot) GetBackgroundChooseKeyboard() tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton
	var row []tgbotapi.InlineKeyboardButton
	for i, frMeta := range bot.generator.GetBackgroundPartsList() {
		if i%4 == 0 && i != 0 {
			keyboard = append(keyboard, row)
			row = make([]tgbotapi.InlineKeyboardButton, 0)
		}
		arr := strings.Split(frMeta.FileName, ".")
		row = append(row, tgbotapi.NewInlineKeyboardButtonData(arr[0], "filename:"+arr[0]))
	}
	keyboard = append(keyboard, row)

	sep := tgbotapi.NewInlineKeyboardButtonData("===", "nodata:")
	row = []tgbotapi.InlineKeyboardButton{sep}
	keyboard = append(keyboard, row)

	row = make([]tgbotapi.InlineKeyboardButton, 0)
	for i := 1; i <= 5; i++ {
		text := fmt.Sprintf("%d", i)
		row = append(row, tgbotapi.NewInlineKeyboardButtonData(text, "count:"+text))
	}
	keyboard = append(keyboard, row)

	row = []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Generate", "commit_custom"),
	}
	keyboard = append(keyboard, row)

	return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}
