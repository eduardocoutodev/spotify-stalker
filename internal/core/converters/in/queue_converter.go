package in

import "github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/dto"

func ConvertToQueueOutbound(queueResponse dto.QueueResponse) dto.QueueOutboundResponse {
	return dto.QueueOutboundResponse{
		CurrentlyPlaying: convertItemToTrackOutbound(queueResponse.CurrentlyPlaying),
		Queue:            convertItemsToTrackOutbound(queueResponse.Queue),
	}
}

func convertItemToTrackOutbound(item dto.Item) dto.TrackOutbound {
	artists := make([]string, len(item.Artists))
	for i, artist := range item.Artists {
		artists[i] = artist.Name
	}

	var albumImage string
	if len(item.Album.Images) > 0 {
		albumImage = item.Album.Images[0].URL
	}

	var previewURL string
	if item.PreviewURL != nil {
		previewURL = item.PreviewURL.(string)
	}

	return dto.TrackOutbound{
		ID:         item.ID,
		Name:       item.Name,
		Artists:    artists,
		AlbumName:  item.Album.Name,
		AlbumImage: albumImage,
		DurationMs: item.DurationMS,
		URI:        item.URI,
		Popularity: item.Popularity,
		PreviewURL: previewURL,
	}
}

func convertItemsToTrackOutbound(items []dto.Item) []dto.TrackOutbound {
	tracks := make([]dto.TrackOutbound, len(items))
	for i, item := range items {
		tracks[i] = convertItemToTrackOutbound(item)
	}
	return tracks
}
