package main

/*
//	FILE: utils.go
//
//  PURPOSE: Utility Methods
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/04/11 Implementation
//
// ----------------------------------------------------------------------------
*/
import (
	"log"

	"github.com/bwmarrin/discordgo"
)

/*
//  METHOD: SetStatus
//
//  PURPOSE: Change bots status through supplied title
//
*/
func SetStatus(s *discordgo.Session, n string) {
	err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{{
			Name: n,
			Type: 3,
		}},
		AFK:    false,
		Status: "",
	})
	if err != nil {
		log.Println("Error attempting to set my status")
		log.Println(err)
	}
}
