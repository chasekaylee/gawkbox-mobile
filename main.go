package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chasekaylee/gawkbox-mobile/twitch"
)

/*
	1) search for live creators based on username
	GET - with query as body
	/api/search
	2) deliever top 10 featured streamers when req made
	GET
	/api/featured
*/

func main() {
	fmt.Println("Booting the server...")
	twitch.SetID("io3r29b0slg7j9qf2woudfzdqshhmh")

	// Initial Route
	http.HandleFunc("/", serveHome)

	// search route
	http.HandleFunc("/api/search", handleSearch)

	http.HandleFunc("/api/featured", handleFeatured)

	// Run your server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to GawkBox!")
}

/*
	handleSearch - A search handler function for the route /search for HTTP server expects a query string (username) from client
	on success returns stream information for user + 200 status code
	if Streamer is offline "Stream is currently offline" is returned + 404 status code
*/
func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println("Recieved the following request:", query.Get("query"))
	t, err := twitch.GetStreamer(query.Get("query"))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	us, err := twitch.GetUserStream(t.ID)
	if err == nil {
		s, err := json.Marshal(us)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(s)
	} else {
		http.Error(w, err.Error(), 404)
	}
}

/*
	handleFeatured - A featured handler function for the route /featured for HTTP server
	on success returns stream slice of information for top 10 featured streams + 200 status code
	on failure returns error + 500 status code
*/
func handleFeatured(w http.ResponseWriter, r *http.Request) {
	f, err := twitch.GetFeaturedStreams()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	fs, err := json.Marshal(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fs)
}
