package dto

type UserCurrentPlayingSpotifyApiResponse struct {
	Timestamp            int64   `json:"timestamp"`
	Context              Context `json:"context"`
	ProgressMS           int64   `json:"progress_ms"`
	Item                 Item    `json:"item"`
	CurrentlyPlayingType string  `json:"currently_playing_type"`
	Actions              Actions `json:"actions"`
	IsPlaying            bool    `json:"is_playing"`
}

type Actions struct {
	Disallows Disallows `json:"disallows"`
}

type Disallows struct {
	Resuming              bool `json:"resuming"`
	TogglingRepeatContext bool `json:"toggling_repeat_context"`
	TogglingRepeatTrack   bool `json:"toggling_repeat_track"`
	TogglingShuffle       bool `json:"toggling_shuffle"`
}

type Context struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Item struct {
	Album        Album        `json:"album"`
	Artists      []Artist     `json:"artists"`
	DiscNumber   int64        `json:"disc_number"`
	DurationMS   int64        `json:"duration_ms"`
	Explicit     bool         `json:"explicit"`
	ExternalIDS  ExternalIDS  `json:"external_ids"`
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	IsLocal      bool         `json:"is_local"`
	Name         string       `json:"name"`
	Popularity   int64        `json:"popularity"`
	PreviewURL   interface{}  `json:"preview_url"`
	TrackNumber  int64        `json:"track_number"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}
