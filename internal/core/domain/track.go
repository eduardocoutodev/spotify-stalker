package domain

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
