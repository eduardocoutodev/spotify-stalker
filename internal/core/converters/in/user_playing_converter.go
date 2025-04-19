package in

import (
	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/dto"
	"github.com/eduardocoutodev/spotify-stalker/internal/core/domain"
)

func ConvertToUserCurrentPlaying(apiResponse *dto.UserCurrentPlayingSpotifyApiResponse) domain.UserCurrentPlaying {
	return domain.UserCurrentPlaying{
		Context: domain.MusicSessionContext{
			SpotifyHref: apiResponse.Context.Href,
			Type:        domain.MusicSessionContextType(apiResponse.Context.Type),
		},
		IsPlaying:            apiResponse.IsPlaying,
		CurrentlyPlayingType: apiResponse.CurrentlyPlayingType,
		CurrentItemPlaying: domain.CurrentItemPlaying{
			ID:          apiResponse.Item.ID,
			Album:       convertToAlbum(&apiResponse.Item.Album),
			Artists:     ConvertToSimpleArtists(&apiResponse.Item.Artists),
			DurationMS:  apiResponse.Item.DurationMS,
			ProgresMS:   apiResponse.ProgressMS,
			Explicit:    apiResponse.Item.Explicit,
			SpotifyHref: apiResponse.Item.ExternalUrls.Spotify,
			Name:        apiResponse.Item.Name,
			Type:        apiResponse.Item.Type,
		},
	}
}

func convertToAlbum(apiResponseAlbum *dto.Album) domain.SimpleAlbum {
	var imageURL string
	if len(apiResponseAlbum.Images) > 0 {
		imageURL = apiResponseAlbum.Images[0].URL
	}

	return domain.SimpleAlbum{
		ID:          apiResponseAlbum.ID,
		Name:        apiResponseAlbum.Name,
		ImageURL:    imageURL,
		ReleaseDate: apiResponseAlbum.ReleaseDate,
	}
}
