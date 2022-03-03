package main

import (
	"NFTProject/internal/bot/discord/pingu"
	"log"
)

func main() {

	bot, err := pingu.NewPinguBot("OTQzMjk1OTQ1ODkyODIzMTAw.Ygw-zw.XRGJVA9I9HzDxLZ5KU56yXGfIlc")
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
