package main

import (
	"NFTProject/internal/bot/telegram/metabot"
	"log"
)

func main() {

	metaBot, err := metabot.NewMetaBot("5013386434:AAFNWsrZ1p8_kMLV6cYhth9ohvcoYMrLix0")
	if err != nil {
		log.Fatal(err)
	}

	metaBot.Run()
	for {
	}
}
