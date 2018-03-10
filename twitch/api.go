package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

// GetStreamer => takes in a search query username
// returns a Streamer struct for that specified username
func GetStreamer(name string) Streamer {
	var streamer Streamers

	twitClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/users?login="+name, nil)
	checkErr(err)

	req.Header.Set("Accept", accept)
	req.Header.Add("Client-ID", tID)

	res, getErr := twitClient.Do(req)
	checkErr(getErr)

	body, readErr := ioutil.ReadAll(res.Body)
	checkErr(readErr)
	res.Body.Close()

	err = json.Unmarshal([]byte(body), &streamer)
	checkErr(err)

	t := streamer.Users[0]

	return t
}

// GetUserStream => takes in a channel ID
// returns a Stream struct for that specified channel ID
func GetUserStream(id string) (StreamData, bool) {
	var stream StreamData
	var online = true

	twitClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/streams/"+id, nil)
	checkErr(err)

	req.Header.Set("Accept", accept)
	req.Header.Add("Client-ID", tID)

	res, getErr := twitClient.Do(req)
	checkErr(getErr)

	body, readErr := ioutil.ReadAll(res.Body)
	checkErr(readErr)
	res.Body.Close()
	if string(body) == `{"stream":null}` {
		online = false
	}

	err = json.Unmarshal([]byte(body), &stream)
	checkErr(err)

	return stream, online
}
