package twitch

import (
	"fmt"
	"log"
)

var tID string

const accept string = "application/vnd.twitchtv.v5+json"

// SetID => takes clientID and sets variable for use in header
func SetID(clientID string) {
	fmt.Println("Setting clientID")
	tID = clientID
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
