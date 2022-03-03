package pingu

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (pb *PinguBot) UpdateMembersWithInterval(s *discordgo.Session, interval time.Duration) {
	for {
	Anchor:
		select {
		case <-pb.closeChan:
			return
		default:
			lastID := ""
			count := 0
			for {
				members, err := s.GuildMembers(guildID, lastID, 1000)
				if err != nil {
					log.Print(err)
					break Anchor
				}
				lastID = members[len(members)-1].User.ID
				count += len(members)
				time.Sleep(interval)
				if len(members) < 1000 {
					break
				}
			}
			newName := fmt.Sprintf("AKIO FRIENDS: %d", count)
			log.Print(newName)
			_, err := s.ChannelEdit(memberCountChannelID, newName)
			if err != nil {
				log.Print(err)
			}
		}
		time.Sleep(interval)
	}
}
