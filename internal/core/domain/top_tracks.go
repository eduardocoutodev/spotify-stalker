package domain

type TopTracksOutputResponse struct {
	Tracks []TopTrack `json:"tracks"`
}

type TopTrack struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Artists    []SimpleArtist `json:"artists"`
	Album      SimpleAlbum    `json:"album"`
	Duration   int            `json:"durationMs"`
	Popularity int            `json:"popularity"`
	UserRank   int            `json:"userRank"`
	SpotifyUrl string         `json:"spotifyUrl"`
}

type SimpleArtist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SimpleAlbum struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"imageUrl"`
	ReleaseDate string `json:"releaseDate"`
}

type TopTracksSpotifyApiResponse struct {
	Items []TrackSpotifyApiResponse `json:"items"`
}

type TrackSpotifyApiResponse struct {
	ID           string              `json:"id"`
	Album        AlbumApiResponse    `json:"album"`
	Artists      []ArtistApiResponse `json:"artists"`
	DurationMs   int                 `json:"duration_ms"`
	Explicit     bool                `json:"explicit"`
	ExternalURLs ExternalURL         `json:"external_urls"`
	Href         string              `json:"href"`
	IsPlayable   bool                `json:"is_playable"`
	Name         string              `json:"name"`
	Popularity   int                 `json:"popularity"`
	TrackNumber  int                 `json:"track_number"`
	Type         string              `json:"type"`
	URI          string              `json:"uri"`
}

type AlbumApiResponse struct {
	AlbumType    string              `json:"album_type"`
	Artists      []ArtistApiResponse `json:"artists"`
	ExternalURLs ExternalURL         `json:"external_urls"`
	Href         string              `json:"href"`
	ID           string              `json:"id"`
	Images       []Image             `json:"images"`
	IsPlayable   bool                `json:"is_playable"`
	Name         string              `json:"name"`
	ReleaseDate  string              `json:"release_date"`
	TotalTracks  int                 `json:"total_tracks"`
	Type         string              `json:"type"`
	URI          string              `json:"uri"`
}

type ArtistApiResponse struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Href         string      `json:"href"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type ExternalURL struct {
	Spotify string `json:"spotify"`
}
