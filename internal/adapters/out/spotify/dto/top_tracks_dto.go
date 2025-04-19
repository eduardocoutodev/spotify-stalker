package dto

type TopTracksSpotifyApiResponse struct {
	Items []TrackSpotifyApiResponse `json:"items"`
}

type TrackSpotifyApiResponse struct {
	ID           string       `json:"id"`
	Album        Album        `json:"album"`
	Artists      []Artist     `json:"artists"`
	DurationMs   int          `json:"duration_ms"`
	Explicit     bool         `json:"explicit"`
	ExternalURLs ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	IsPlayable   bool         `json:"is_playable"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	TrackNumber  int          `json:"track_number"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}
