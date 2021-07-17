// +build pummel
package commands

import (
	"database/sql"
	"ebotclique/database"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func HandlePummelCommand(s *discordgo.Session, m *discordgo.Message, name string, args []string) {
	// no args is boards, winner arg goes to mention user
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

// write sql procedure to return formated string such as switch1, Burnout3, etc...
func getLeaderboard(name string) string {
	var quote sql.NullString
	db := database.Connect()
	query := fmt.Sprintf("SELECT quote from switch_db.ebotclique_quotes_t WHERE user = '%s' ORDER BY RAND() LIMIT 1", name)
	err := db.QueryRow(query)
	if err != nil {
		log.Println("Pummel::getLeaderboard: ERROR retrieving database query")
	}
	if !quote.Valid {
		quote.String = "No quotes found for this user"
	}
	db.Close()
	return quote.String
}

//procedure call to update table. In scratch pad already
func updateLeaderboard(name string) {
	var quote sql.NullString
	db := database.Connect()
	query := fmt.Sprintf("UPDATE switch_db.leaderboards_t SET score = score + 1 WHERE `name` = '%s'", name)
	err := db.QueryRow(query).Scan(&quote)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Quote::getRandomQuote: ERROR retrieving database query")
	}
	if !quote.Valid {
		quote.String = "No quotes found for this user"
	}
	db.Close()
}
