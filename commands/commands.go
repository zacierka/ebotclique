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

	msg := strings.Split(strings.TrimSpace(m.Content), ".")[1]

	if len(msg) > 2 {
		msg = strings.Split(strings.Split(m.Content, " ")[0], ".")[1]
	}

	switch msg {
	case "info":
		HandleInfoCommand(s, m, t0)
	case "ping":
		HandlePingCommand(s, m)
	case "mark":
		HandleMarkCommand(s, m)
	default:
		//HandleUnknownCommand(s, m, msg) // sends message to users DM of invalid message
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
