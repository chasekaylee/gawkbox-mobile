package twitch

// Streamers =>
type Streamers struct {
	Users []Streamer `json:"users"`
}

// Streamer => data structure for Users slice
type Streamer struct {
	ID          string `json:"_id"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
}

// StreamData =>
type StreamData struct {
	Data Stream `json:"stream"`
}

// Stream =>
type Stream struct {
	ID         int     `json:"_id"`
	Game       string  `json:"game"`
	Viewers    int     `json:"viewers"`
	VidHeight  int     `json:"video_height"`
	AvgFPS     float64 `json:"average_fps"`
	Delay      int     `json:"delay"`
	CreatedAt  string  `json:"created_at"`
	IsPlaylist bool    `json:"is_playlist"`
	Preview    Preview `json:"preview"`
	Channel    Channel `json:"channel"`
}

// Preview => data structure for nested struct within Stream
type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}

// Channel => data structure for nested struct within Stream
type Channel struct {
	Mature          bool   `json:"mature"`
	Status          string `json:"status"`
	Lang            string `json:"broadcaster_language"`
	ID              int    `json:"_id"`
	Name            string `json:"name"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Partner         bool   `json:"partner"`
	Logo            string `json:"logo"`
	VidBanner       string `json:"video_banner"`
	ProfBanner      string `json:"profile_banner"`
	ProfBannerColor string `json:"profile_banner_background_color"`
	URL             string `json:"url"`
	Views           int    `json:"views"`
	Followers       int    `json:"followers"`
}
