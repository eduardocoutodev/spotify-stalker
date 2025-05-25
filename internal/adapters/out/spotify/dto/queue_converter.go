package dto

func ConvertToQueueOutbound(queueResponse QueueResponse) QueueOutboundResponse {
	return QueueOutboundResponse{
		CurrentlyPlaying: convertItemToTrackOutbound(queueResponse.CurrentlyPlaying),
		Queue:            convertItemsToTrackOutbound(queueResponse.Queue),
	}
}

func convertItemToTrackOutbound(item Item) TrackOutbound {
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

	return TrackOutbound{
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

func convertItemsToTrackOutbound(items []Item) []TrackOutbound {
	tracks := make([]TrackOutbound, len(items))
	for i, item := range items {
		tracks[i] = convertItemToTrackOutbound(item)
	}
	return tracks
}
