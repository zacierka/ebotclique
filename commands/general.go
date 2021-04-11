package commands

/*
//	FILE: general.go
//
//  PURPOSE: file for holding most command implementations
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/04/03 Implementation
//
//  ORIGINAL CREDITS: https://github.com/Sdyess/AlfredBot/blob/master/commands/general.go
//
// ----------------------------------------------------------------------------
*/
import (
	"bytes"
	"database/sql"
	"ebotclique/database"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

/*
//  METHOD: HandleInfoCommand <prefix>info
//
//  PURPOSE: display a list of channel data
//
//  ARGS: NONE
*/
func HandleInfoCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
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

/*
//  METHOD: HandlePingCommand <prefix>ping
//
//  PURPOSE: display a string to verify connection
//
//  ARGS: NONE
*/
func HandlePingCommand(s *discordgo.Session, m *discordgo.Message) {

	s.ChannelMessageSend(m.ChannelID, "pong")
}

/*
//  METHOD: HandleMarkCommand <prefix>mark
//
//  PURPOSE: display a currated list of quotes
//
//  ARGS: add        add new quote
//        remove     remove a quote
//        list       view all quotes
*/
func HandleMarkCommand(s *discordgo.Session, m *discordgo.Message) {
	mode := strings.Split(m.Content, " ")
	if len(mode) == 1 {
		randomMarkQuote(s, m)
	}
	if len(mode) > 1 {
		if mode[1] == "add" {
			newMarkQuote(s, m)
		} else if mode[1] == "list" {
			listMarkQuote(s, m)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Usage: .mark | .mark add _quote_ | .mark list")
		}
	}
}

func newMarkQuote(s *discordgo.Session, m *discordgo.Message) {
	// create quote
	var id sql.NullInt64
	db := database.Connect()
	quote := strings.TrimSpace(strings.Split(m.Content, ".mark add")[1])
	row := db.QueryRow(
		"CALL newquote(?)",
		quote,
	)

	if err := row.Scan(&id); err != nil {
		s.ChannelMessageSend(m.ChannelID, "Added Quote: "+quote)
		log.Printf("HandleMarkCommand::newMarkQuote(%s)::BY_USER::%s", quote, m.Author.Username)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Error creating quote. Try again.")
		log.Println("[Error] newMarkQuote()::Could not add Quote @", m.Author.Username)
	}
	db.Close()
}

// func removeMarkQuote(s *discordgo.Session, m *discordgo.Message) {
// 	// @TODO placeholder... find a intuitive way to remove (list #? exact? )
// }

func randomMarkQuote(s *discordgo.Session, m *discordgo.Message) {
	var quote string
	db := database.Connect()
	err := db.QueryRow("SELECT quote from switch_db.ebotclique_markquotes_t ORDER BY RAND() LIMIT 1").Scan(&quote)
	if err != nil {
		log.Println("Cannot recall a quote...")
		return
	}
	s.ChannelMessageSend(m.ChannelID, quote)
	db.Close()
}

func listMarkQuote(s *discordgo.Session, m *discordgo.Message) {
	var quotes bytes.Buffer
	quotes.WriteString("```")
	db := database.Connect()
	results, err := db.Query("SELECT quote from switch_db.ebotclique_markquotes_t ORDER BY quote ASC")
	if err != nil {
		log.Println("Error ... Cannot recall quotes")
		return
	}
	counter := 1
	for results.Next() {
		var quote string
		err = results.Scan(&quote)
		if err != nil {
			panic(err.Error())
		}
		quotes.WriteString(strconv.Itoa(counter))
		quotes.WriteString(". ")
		quotes.WriteString(quote)
		quotes.WriteString("\n")
		counter = counter + 1
	}
	quotes.WriteString("```")
	s.ChannelMessageSend(m.ChannelID, quotes.String())
	log.Printf("HandleMarkCommand::listMarkQuote()::BY::%s", m.Author.Username)
	db.Close()
}
