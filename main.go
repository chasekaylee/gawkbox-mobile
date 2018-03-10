package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chasekaylee/gawkbox-mobile/twitch"
)

/*
	1) search for live creators based on username
	POST - with query as body
	/api/search
	2) click on stream and deliver individual results
	GET - send streamer name
	/api/stream
*/

func main() {
	fmt.Println("Booting the server...")
	twitch.SetID("io3r29b0slg7j9qf2woudfzdqshhmh")

	// Initial Route
	http.HandleFunc("/", serveHome)

	// search route
	http.HandleFunc("/api/search", handleSearch)

	// Run your server
	http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GawkBox!")
}

/*
	handleSearch - A search handler function for the route /search for HTTP server expects a query string (username) from client
	on success returns stream information for user + 200 status code
	if Streamer is offline "user offline" is returned
*/
func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Fprintf(w, "Recieved the following request: %s\n", query.Get("query"))
	t := twitch.GetStreamer(query.Get("query"))
	us, online := twitch.GetUserStream(t.ID)
	if online {
		s, err := json.Marshal(us)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(s)
	} else {
		io.WriteString(w, `{"online": false}`)
	}
}
