package main

import (
	"NFTProject/internal/bot/telegram"
	"NFTProject/internal/bot/telegram/noisebot"
	"log"
)

func main() {

	bot, err := telegram.NewTelegramBot("5224244154:AAHkrSAKlexF1ICdLV_OGNPjpNJuqFHzmQY")
	if err != nil {
		log.Fatal(err)
	}

	handler := noisebot.NoisebotHandler{}
	bot.BindHandler(&handler)
	bot.Run()
	for {

	}
}
