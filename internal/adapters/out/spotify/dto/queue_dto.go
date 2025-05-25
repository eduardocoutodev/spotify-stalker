package dto

type QueueResponse struct {
	CurrentlyPlaying Item   `json:"currently_playing"`
	Queue            []Item `json:"queue"`
}

type QueueOutboundResponse struct {
	CurrentlyPlaying TrackOutbound   `json:"currentlyPlaying"`
	Queue            []TrackOutbound `json:"queue"`
}

type TrackOutbound struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Artists    []string `json:"artists"`
	AlbumName  string   `json:"albumName"`
	AlbumImage string   `json:"albumImage"`
	DurationMs int64    `json:"durationMs"`
	URI        string   `json:"uri"`
	Popularity int64    `json:"popularity"`
	PreviewURL string   `json:"previewUrl,omitempty"`
}
