package handler_spotify

import (
	core "github.com/eduardoooxd/spotify-tracker/core/domain"
)

func TransformTopTracks(spotifyApiResponse *core.TopTracksSpotifyApiResponse) core.TopTracksOutputResponse {
	outputResponse := core.TopTracksOutputResponse{
		Tracks: make([]core.TopTrack, len(spotifyApiResponse.Items)),
	}

	for i, apiTrack := range spotifyApiResponse.Items {
		var imageURL string
		if len(apiTrack.Album.Images) > 0 {
			imageURL = apiTrack.Album.Images[0].URL
		}

		artists := make([]core.SimpleArtist, len(apiTrack.Artists))
		for j, artist := range apiTrack.Artists {
			artists[j] = core.SimpleArtist{
				ID:   artist.ID,
				Name: artist.Name,
				URL:  artist.ExternalURLs.Spotify,
			}
		}

		outputResponse.Tracks[i] = core.TopTrack{
			ID:         apiTrack.ID,
			Name:       apiTrack.Name,
			Artists:    artists,
			Duration:   apiTrack.DurationMs,
			Popularity: apiTrack.Popularity,
			SpotifyUrl: apiTrack.ExternalURLs.Spotify,
			UserRank:   i + 1,
			Album: core.SimpleAlbum{
				ID:          apiTrack.Album.ID,
				Name:        apiTrack.Album.Name,
				ImageURL:    imageURL,
				ReleaseDate: apiTrack.Album.ReleaseDate,
			},
		}
	}

	return outputResponse
}
