package commands

/*
//	FILE: commands.go
//
//  PURPOSE: Router module for commands
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/04/02 Implementation
//
// ----------------------------------------------------------------------------
*/
import (
	"strings"

	"time"

	"github.com/bwmarrin/discordgo"
)

//ExecuteCommand Parses and executes the command from the calling code
func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	rawcmd := strings.TrimSpace(m.Content)
	temp := strings.Split(rawcmd, " ")
	args := temp[1:]

	switch name := strings.Split(temp[0], ".")[1]; name {
	case "switch", "cloak", "jack", "lit", "mark", "burnout":
		HandleQuoteCommand(s, m, name, args)
	case "uptime":
		HandleInfoCommand(s, m, time.Now())
	}
}

//HandleUnknownCommand is the default for any commands not listed
func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not recognized.")
}

func HandleWrongPermissions(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not available to you.")
}
