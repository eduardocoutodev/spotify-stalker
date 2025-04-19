package in

import (
	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/dto"
	"github.com/eduardocoutodev/spotify-stalker/internal/core/domain"
)

func ConvertToSimpleArtists(artistsApiResponse *[]dto.Artist) []domain.SimpleArtist {
	artists := make([]domain.SimpleArtist, len(*artistsApiResponse))

	for j, artist := range *artistsApiResponse {
		artists[j] = domain.SimpleArtist{
			ID:   artist.ID,
			Name: artist.Name,
			URL:  artist.ExternalUrls.Spotify,
		}
	}

	return artists
}
