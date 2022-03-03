package pingu

import (
	"NFTProject/internal/bot/discord/pingu/imgtools"
	"bytes"
	"fmt"
	"image/png"
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendHelloMessage(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	user, err := s.User(e.UserID)
	if err != nil {
		log.Print(err)
		return
	}

	img, err := s.UserAvatar(user.ID)
	if err != nil {
		log.Print(err)
		return
	}

	text := fmt.Sprintf("Hello, %s #%s", user.Username, user.Discriminator)
	helloImg, err := imgtools.MakeHelloImage(img, text)
	if err != nil {
		log.Print(err)
	}

	var buff []byte
	rw := bytes.NewBuffer(buff)
	err = png.Encode(rw, helloImg)
	if err != nil {
		log.Print(err)
		return
	}

	filename := fmt.Sprintf("%s#%s.png", user.Username, user.Discriminator)
	message := fmt.Sprintf("%s", user.Mention())
	_, err = s.ChannelFileSendWithMessage(newPinguinsChannelID, message, filename, rw)
	if err != nil {
		log.Print(err)
	}
}
