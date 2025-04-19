package in

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"

	out "github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify"
	converters "github.com/eduardocoutodev/spotify-stalker/internal/core/converters/in"
	"github.com/eduardocoutodev/spotify-stalker/internal/core/domain"
)

func HandleTopTracks(w http.ResponseWriter, r *http.Request) {
	spotifyToken := os.Getenv("SPOTIFY_TOKEN")

	reqHeaders := make(map[string]string)
	reqHeaders["Authorization"] = "Bearer " + spotifyToken

	resp, err := out.FetchSpotifyWebAPI(
		out.SpotifyRequestArguments{
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

	var apiResponse domain.TopTracksSpotifyApiResponse
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		slog.Error("Failed reading response body", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to parse response from Spotify"})
		return
	}

	tracksResponseBody := converters.TransformTopTracks(&apiResponse)

	if err := json.NewEncoder(w).Encode(&tracksResponseBody); err != nil {
		slog.Error("Failed writing to the response", slog.Any("err", err))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to encode body from Spotify API"})
		return
	}
}
