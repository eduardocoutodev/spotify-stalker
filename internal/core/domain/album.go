package domain

type SimpleAlbum struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"imageUrl"`
	ReleaseDate string `json:"releaseDate"`
}
