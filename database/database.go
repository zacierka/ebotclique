package database

/*
//	FILE: database
//
//  PURPOSE: Handle Database Connection
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/04/02 Implementation
//	  21/04/10 Add Test Connection
//    21/07/06 Add Env Vars Usage
//  ORIGINAL CREDITS: https://github.com/Sdyess/AlfredBot/blob/master/database/database.go
//
// ----------------------------------------------------------------------------
*/
import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	Path string
)

func Connect() *sql.DB {
	path, ok := os.LookupEnv("DBPATH")
	if ok {
		Path = path
	} else {
		panic(1)
	}

	db, err := sql.Open("mysql", Path)
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
