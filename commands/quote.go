package commands

/*
//	FILE: quote.go
//
//  PURPOSE: handles quote command
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/06/26 Implementation
//    21/07/06 Added Adaptive User Functionality
//
// ----------------------------------------------------------------------------
*/
import (
	"database/sql"
	"ebotclique/database"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func HandleQuoteCommand(s *discordgo.Session, m *discordgo.Message, name string, args []string) {

	if len(args) == 0 {
		q := getRandomQuote(name)
		s.ChannelMessageSend(m.ChannelID, q)
		log.Printf("COMMAND::%s:CALLED_BY:%s:SUCCESS", name, m.Author.Username)
	} else if args[0] == "add" {
		if hasRole(m.Member.Roles, "549469405306945538") { // admin role
			quote := strings.Join(args[1:], " ")
			res := QuoteAdd(name, quote)
			if res {
				s.ChannelMessageSend(m.ChannelID, "Added Quote Successfully")
				log.Printf("COMMAND::%s_ADD:CALLED_BY:%s:SUCCESS", name, m.Author.Username)
			} else {
				log.Printf("COMMAND::%s_ADD:CALLED_BY:%s:FAILED", name, m.Author.Username)
			}
		} else {
			reply := fmt.Sprintf("%s, you do not have the permissions to use that.", m.Author.Mention())
			msg, _ := s.ChannelMessageSend(m.ChannelID, reply)
			time.AfterFunc(5*time.Second, func() {
				s.ChannelMessageDelete(m.ChannelID, msg.ID)
			})
		}
	}
}

func getRandomQuote(name string) string {
	var quote sql.NullString
	db := database.Connect()
	query := fmt.Sprintf("SELECT quote from switch_db.ebotclique_quotes_t WHERE user = '%s' ORDER BY RAND() LIMIT 1", name)
	err := db.QueryRow(query).Scan(&quote)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Quote::getRandomQuote: ERROR retrieving database query")
	}
	if !quote.Valid {
		quote.String = "No quotes found for this user"
	}
	db.Close()
	return quote.String
}

func QuoteAdd(name string, quote string) bool {
	var res bool
	var id sql.NullInt64
	db := database.Connect()
	row := db.QueryRow(
		"CALL addquote(?, ?)",
		quote,
		name,
	)

	if err := row.Scan(&id); err != nil {
		res = true
	} else {
		res = false
	}
	db.Close()
	return res
}

func hasRole(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
