package dto

type Album struct {
	AlbumType    string       `json:"album_type"`
	Artists      []Artist     `json:"artists"`
	ExternalURLs ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []Image      `json:"images"`
	IsPlayable   bool         `json:"is_playable"`
	Name         string       `json:"name"`
	ReleaseDate  string       `json:"release_date"`
	TotalTracks  int          `json:"total_tracks"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}
