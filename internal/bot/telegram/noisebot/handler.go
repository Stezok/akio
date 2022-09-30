package noisebot

import (
	"NFTProject/internal/bot/telegram"
	"NFTProject/pkg/imago"
	"bytes"
	"image/png"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NoisebotHandler struct {
	bot *tgbotapi.BotAPI
}

func (h *NoisebotHandler) GetHandlers() []telegram.TelegramHandler {
	return []telegram.TelegramHandler{
		h.HandlerReciveDocument,
		h.HandlerUpdateLogger,
	}
}

func (h *NoisebotHandler) GetCallbackHandlers() []telegram.TelegramHandler {
	return []telegram.TelegramHandler{
		h.CallbackAnotherPart,
		h.CallbackBackPart,
	}
}

func (h *NoisebotHandler) BindBot(bot *tgbotapi.BotAPI) {
	h.bot = bot
}

func (h *NoisebotHandler) HandlerUpdateLogger(u *tgbotapi.Update) bool {
	log.Printf("%+v", *u.Message)
	return false
}

func (h *NoisebotHandler) HandlerReciveDocument(u *tgbotapi.Update) bool {
	if u.Message.Document == nil {
		return false
	}

	docConfig := tgbotapi.NewDocument(u.Message.Chat.ID, tgbotapi.FileID(u.Message.Document.FileID))
	docConfig.ReplyMarkup = partChoiseKeyboard
	h.bot.Send(docConfig)
	return true
}

func (h *NoisebotHandler) ProcessAndSendPhoto(u *tgbotapi.Update, mode imago.BlendMode, freq float64) {
	file, err := h.bot.GetFile(tgbotapi.FileConfig{
		FileID: u.CallbackQuery.Message.Document.FileID,
	})
	if err != nil {
		log.Print(err)
		return
	}

	resp, err := http.Get(file.Link(h.bot.Token))
	if err != nil {
		log.Print(err)
		return
	}

	grainNoise, err := imago.NewGrainNoiseFromReader(resp.Body, freq)
	if err != nil {
		log.Print(err)
		return
	}
	grainNoise.Mode = mode

	var buf bytes.Buffer
	err = png.Encode(&buf, grainNoise)
	if err != nil {
		log.Print(err)
		return
	}

	docConfig := tgbotapi.NewDocument(u.CallbackQuery.Message.Chat.ID, tgbotapi.FileReader{
		Name:   u.CallbackQuery.Message.Document.FileName,
		Reader: &buf,
	})
	h.bot.Send(docConfig)
}

func (h *NoisebotHandler) CallbackAnotherPart(u *tgbotapi.Update) bool {
	if u.CallbackData() != ANOTHER_TAG {
		return false
	}

	h.ProcessAndSendPhoto(u, imago.Overlay, 0.1)
	return true
}

func (h *NoisebotHandler) CallbackBackPart(u *tgbotapi.Update) bool {
	if u.CallbackData() != BACK_TAG {
		return false
	}

	h.ProcessAndSendPhoto(u, imago.Custom, 0.6)
	return true
}
