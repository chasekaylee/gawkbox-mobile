package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// GetStreamer => takes in a search query username
// returns a Streamer struct for that specified username
// & an error relative to where the function fails
func GetStreamer(name string) (Streamer, error) {
	var streamer Streamers

	twitClient := http.Client{
		Timeout: time.Second * 200, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/users?login="+name, nil)
	if err != nil {
		return streamer.Users[0], err
	}

	req.Header.Set("Accept", accept)
	req.Header.Add("Client-ID", tID)

	res, getErr := twitClient.Do(req)
	if getErr != nil {
		return streamer.Users[0], getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return streamer.Users[0], readErr
	}

	res.Body.Close()

	err = json.Unmarshal([]byte(body), &streamer)
	if err != nil {
		return streamer.Users[0], err
	}

	t := streamer.Users[0]

	return t, nil
}

// GetUserStream => takes in a channel ID
// returns a Stream struct for that specified channel ID
// & an error relative to where the function fails
func GetUserStream(id string) (StreamData, error) {
	var stream StreamData

	twitClient := http.Client{
		Timeout: time.Second * 200, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/streams/"+id, nil)
	if err != nil {
		return stream, err
	}

	req.Header.Set("Accept", accept)
	req.Header.Add("Client-ID", tID)

	res, getErr := twitClient.Do(req)
	if getErr != nil {
		return stream, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return stream, readErr
	}

	res.Body.Close()
	if string(body) == `{"stream":null}` {
		return stream, errors.New("Stream is currently offline")
	}

	err = json.Unmarshal([]byte(body), &stream)
	if err != nil {
		return stream, err
	}

	return stream, nil
}

// GetFeaturedStreams =>
// returns a Featured struct for the top 10 featured streams
// & an error relative to where the function fails
func GetFeaturedStreams() (Featured, error) {
	var featured Featured

	twitClient := http.Client{
		Timeout: time.Second * 200, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/streams/featured?limit=10", nil)
	if err != nil {
		return featured, err
	}

	req.Header.Set("Accept", accept)
	req.Header.Add("Client-ID", tID)

	res, getErr := twitClient.Do(req)
	if getErr != nil {
		return featured, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return featured, readErr
	}
	res.Body.Close()

	err = json.Unmarshal([]byte(body), &featured)
	if err != nil {
		return featured, err
	}

	return featured, nil
}
