package database

/*
//	FILE: commands.go
//
//  PURPOSE: Router module for commands
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/04/02 Implementation
//	  21/04/10 Add Test Connection
//  ORIGINAL CREDITS: https://github.com/Sdyess/AlfredBot/blob/master/database/database.go
//
// ----------------------------------------------------------------------------
*/
import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", DB_DEV)
	if err != nil {
		fmt.Println("[ERROR] Unable to connect to database: ", err)
		return nil
	}
	return db
}

func TestConnection() {
	var quote string
	db := Connect()
	err := db.QueryRow("SELECT `quote` from switch_db.ebotclique_markquotes_t ORDER BY RAND() LIMIT 1").Scan(&quote)
	if err != nil {
		log.Println("Cannot retrieve local connection")
		return
	}
	log.Println("Database::testConnection::SUCCESSFUL")
	db.Close()
}
