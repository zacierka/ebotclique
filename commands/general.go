package commands

/*
//	FILE: general.go
//
//  PURPOSE: file for holding most command implementations
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/04/03 Implementation
//    21/07/06 Rework on readability
//  ORIGINAL CREDITS: https://github.com/Sdyess/AlfredBot/blob/master/commands/general.go
//
// ----------------------------------------------------------------------------
*/
import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func HandleInfoCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	if hasRole(m.Member.Roles, "549469405306945538") { // admin role
		t1 := time.Now()
		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			fmt.Println("[ERROR] Issue finding channel by ID: ", err)
			return
		}

		channelName := channel.Name
		message := "```txt\n%s\n%s\n%-16s%-20s\n%-16s%-20s\n%-16s%-20s```"
		message = fmt.Sprintf(message, "EBotClique", strings.Repeat("-", len("EBotClique")), "ChannelID", m.ChannelID, "Channel Name", channelName, "Uptime", (t1.Sub(t0).String()))
		s.ChannelMessageSend(m.ChannelID, message)
	}
}

func HandlePingCommand(s *discordgo.Session, m *discordgo.Message) {
	if m.Author.ID == "164833787543420928" { // me
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
