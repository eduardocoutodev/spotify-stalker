package domain

type UserCurrentPlaying struct {
	Context              *MusicSessionContext `json:"context"`
	IsPlaying            bool                 `json:"isPlaying"`
	CurrentlyPlayingType *string              `json:"currentlyPlayingType"`
	CurrentItemPlaying   *CurrentItemPlaying  `json:"currentItemPlaying"`
}

type MusicSessionContextType string

const (
	Artist   MusicSessionContextType = "artist"
	Playlist MusicSessionContextType = "playlist"
	Album    MusicSessionContextType = "album"
	Show     MusicSessionContextType = "show"
)

type MusicSessionContext struct {
	SpotifyHref string                  `json:"spotifyHref"`
	Type        MusicSessionContextType `json:"type"`
}

type CurrentItemPlaying struct {
	ID          string         `json:"id"`
	Album       SimpleAlbum    `json:"album"`
	Artists     []SimpleArtist `json:"artists"`
	DurationMS  int64          `json:"durationMs"`
	ProgresMS   int64          `json:"progressMs"`
	Explicit    bool           `json:"explicit"`
	SpotifyHref string         `json:"spotifyHref"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
}
