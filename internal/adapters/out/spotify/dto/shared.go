package dto

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type ExternalIDS struct {
	Isrc string `json:"isrc"`
}
