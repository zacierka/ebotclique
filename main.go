package main

/*
//	FILE: main.go
//
//  PURPOSE: main for ebotclique
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/04/02 Implementation
//	  21/04/11 Add Features
// ----------------------------------------------------------------------------
*/
import (
	"database/sql"
	"ebotclique/commands"
	"ebotclique/database"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Token = DEV_KEY
	BotID string
)

var db *sql.DB
var t0 time.Time

func init() {

	//flag.StringVar(&Token, "t", "", "Bot Token")
	//flag.Parse()
}
func main() {

	t0 = time.Now()
	db = database.Connect()
	if db == nil {
		return
	}
	database.TestConnection()
	defer db.Close()

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("Error creating discord session.", err)
		return
	}
	log.Println("Session Created.")

	u, err := dg.User("@me")
	if err != nil {
		log.Println("A problem occurred while obtaining account details.", err)
		return
	}

	BotID = u.ID

	dg.AddHandler(messageCreate)

	SetStatus(dg, DEFAULT_STATUS)

	err = dg.Open()
	if err != nil {
		log.Println("A problem occurred while opening a connection.", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if s == nil || m == nil {
		return
	}

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "" {
		return
	}

	if m.Content[0] == PREFIX && strings.Count(m.Content, string(PREFIX)) < 2 {

		commands.ExecuteCommand(s, m.Message, t0)
		return
	}
}
