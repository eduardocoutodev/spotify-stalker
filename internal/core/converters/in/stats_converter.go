package in

import "github.com/eduardocoutodev/spotify-stalker/internal/core/domain"

func TransformTopTracks(spotifyApiResponse *domain.TopTracksSpotifyApiResponse) domain.TopTracksOutputResponse {
	outputResponse := domain.TopTracksOutputResponse{
		Tracks: make([]domain.TopTrack, len(spotifyApiResponse.Items)),
	}

	for i, apiTrack := range spotifyApiResponse.Items {
		var imageURL string
		if len(apiTrack.Album.Images) > 0 {
			imageURL = apiTrack.Album.Images[0].URL
		}

		artists := make([]domain.SimpleArtist, len(apiTrack.Artists))
		for j, artist := range apiTrack.Artists {
			artists[j] = domain.SimpleArtist{
				ID:   artist.ID,
				Name: artist.Name,
				URL:  artist.ExternalURLs.Spotify,
			}
		}

		outputResponse.Tracks[i] = domain.TopTrack{
			ID:         apiTrack.ID,
			Name:       apiTrack.Name,
			Artists:    artists,
			Duration:   apiTrack.DurationMs,
			Popularity: apiTrack.Popularity,
			SpotifyUrl: apiTrack.ExternalURLs.Spotify,
			UserRank:   i + 1,
			Album: domain.SimpleAlbum{
				ID:          apiTrack.Album.ID,
				Name:        apiTrack.Album.Name,
				ImageURL:    imageURL,
				ReleaseDate: apiTrack.Album.ReleaseDate,
			},
		}
	}

	return outputResponse
}
