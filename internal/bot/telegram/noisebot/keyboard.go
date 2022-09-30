package noisebot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	BACK_TAG    = "back"
	ANOTHER_TAG = "another"

	partChoiseKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Фон", BACK_TAG),
			tgbotapi.NewInlineKeyboardButtonData("Другое", ANOTHER_TAG),
		),
	)
)
