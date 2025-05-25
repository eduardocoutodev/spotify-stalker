package dto

type QueueResponse struct {
	CurrentlyPlaying Item   `json:"currently_playing"`
	Queue            []Item `json:"queue"`
}
