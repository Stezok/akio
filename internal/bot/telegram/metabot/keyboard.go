package metabot

import (
	"NFTProject/internal/meta"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetAcceptKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Accept", "accept"),
		),
	)
}

func GetSlotKeyboard() tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	var row []tgbotapi.InlineKeyboardButton
	for i, slot := range meta.Slots {
		if i%4 == 0 {
			if i != 0 {
				rows = append(rows, row)
			}
			row = make([]tgbotapi.InlineKeyboardButton, 0, 4)
		}
		row = append(row, tgbotapi.NewInlineKeyboardButtonData(string(slot), slot.GetData()))
	}
	if len(row) != 0 {
		rows = append(rows, row)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func GetColorKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("None", meta.NoneColor.GetData()),
			tgbotapi.NewInlineKeyboardButtonData("Red", meta.RedColor.GetData()),
			tgbotapi.NewInlineKeyboardButtonData("Pink", meta.PinkColor.GetData()),
			tgbotapi.NewInlineKeyboardButtonData("Cyan", meta.CyanColor.GetData()),
			tgbotapi.NewInlineKeyboardButtonData("Cosmic", meta.CosmicColor.GetData()),
		),
	)
}

func GetRarityKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Common", meta.Common.GetData()),
			tgbotapi.NewInlineKeyboardButtonData("Rare", meta.Rare.GetData()),
			tgbotapi.NewInlineKeyboardButtonData("Epic", meta.Epic.GetData()),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Legendary", meta.Legendary.GetData()),
			tgbotapi.NewInlineKeyboardButtonData("Promo", meta.Promo.GetData()),
		),
	)
}

func GetPromoTagKeyboard() tgbotapi.InlineKeyboardMarkup {
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
