package taskbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *TaskBot) handleCallbackArtService() {

}

func (bot *TaskBot) handleCallbackAddService(c *tgbotapi.CallbackQuery) {
	uid = c.From.ID
}

func (bot *TaskBot) handleStart(u *tgbotapi.Update) error {
	if u.Message.Text != "/start" {
		return nil
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выбери свою команду")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Художник", dataArtServiceChoose),
			tgbotapi.NewInlineKeyboardButtonData("Реклама", dataAddServiceChoose),
		),
	)
	_, err := bot.Send(msg)
	return err
}

func (bot *TaskBot) handleMessage(u *tgbotapi.Update) error {
	uid := u.Message.From.ID
	if !bot.isTeamMember(uid) {
		return nil
	}

	if err := bot.handleStart(u); err != nil {
		return err
	}

	return nil
}
