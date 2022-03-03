package main

import (
	"NFTProject/internal/bot/telegram/pengubot"
	"NFTProject/internal/generator/penguin"
	"log"
	"math/rand"
	"time"
)

func main() {

	gen := penguin.NewPinguinGenerator("./assets/gendata")
	rand.Seed(time.Now().Unix())

	// 1139557493:AAHlidCIg0TwzJAiAe1JiNRI-9pnwUazGr8
	// 5138889762:AAFSw8G2k0v3uYZEseJ61PoJ2b-VvUjt3ZM
	bot, err := pengubot.NewPenguBot("5138889762:AAFSw8G2k0v3uYZEseJ61PoJ2b-VvUjt3ZM", gen)
	if err != nil {
		log.Fatal(err)
	}
	err = bot.Run()
	if err != nil {
		log.Fatal(err)
	}
	for {
	}
}
