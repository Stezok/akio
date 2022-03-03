package pingu

import (
	"github.com/bwmarrin/discordgo"
)

const (
	verifyMessageID      = "941794050581872703"
	verifyChannelID      = "941332208252256297"
	newPinguinsChannelID = "941781304876347392"
	memberCountChannelID = "941334534513889290"
	guildID              = "938343272957509672"
)

func (pb *PinguBot) VerifyMessageReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	if e.ChannelID != verifyChannelID || e.MessageID != verifyMessageID || e.GuildID != guildID {
		return
	}

	SendHelloMessage(s, e)
}
