package handler_spotify

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"

	core "github.com/eduardoooxd/spotify-tracker/core/domain"
	"github.com/eduardoooxd/spotify-tracker/output/spotify"
)

func HandleTopTracks(w http.ResponseWriter, r *http.Request) {
	spotifyToken := os.Getenv("SPOTIFY_TOKEN")

	reqHeaders := make(map[string]string)
	reqHeaders["Authorization"] = "Bearer " + spotifyToken

	resp, err := spotify.FetchSpotifyWebAPI(
		spotify.SpotifyRequestArguments{
			Method:             "GET",
			Endpoint:           "https://api.spotify.com/v1/me/top/tracks?time_range=long_term&limit=10",
			Headers:            reqHeaders,
			ExpectedStatusCode: http.StatusOK,
		},
	)

	if err != nil {
		slog.Error("Failed making the request to spotify", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch data from Spotify API"})
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed reading response body", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to read response from Spotify"})
		return
	}

	var apiResponse core.TopTracksSpotifyApiResponse
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		slog.Error("Failed reading response body", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to parse response from Spotify"})
		return
	}

	tracksResponseBody := TransformTopTracks(&apiResponse)

	if err := json.NewEncoder(w).Encode(&tracksResponseBody); err != nil {
		slog.Error("Failed writing to the response", slog.Any("err", err))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to encode body from Spotify API"})
		return
	}
}
