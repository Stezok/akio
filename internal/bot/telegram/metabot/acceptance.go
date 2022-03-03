package metabot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (bot *MetaBot) AcceptMetadata(u *tgbotapi.Update, text string) error {

	edit := tgbotapi.NewEditMessageCaption(
		u.CallbackQuery.Message.Chat.ID,
		u.CallbackQuery.Message.MessageID,
		text,
	)
	keyboard := GetAcceptKeyboard()
	edit.BaseEdit.ReplyMarkup = &keyboard

	_, err := bot.Send(edit)
	return err
}
