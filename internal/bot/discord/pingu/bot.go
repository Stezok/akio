package pingu

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type PinguBot struct {
	session *discordgo.Session

	closeChan chan struct{}
}

func (pb *PinguBot) Run() error {

	pb.session.Debug = true

	pb.session.AddHandler(pb.VerifyMessageReactionAdd)
	pb.session.Identify.Intents = discordgo.IntentsAll

	pb.UpdateMembersWithInterval(pb.session, 20*time.Second)

	return pb.session.Open()
}

func NewPinguBot(token string) (*PinguBot, error) {
	authStr := fmt.Sprintf("Bot %s", token)
	session, err := discordgo.New(authStr)
	if err != nil {
		return nil, err
	}

	return &PinguBot{
		session:   session,
		closeChan: make(chan struct{}, 1),
	}, nil
}
