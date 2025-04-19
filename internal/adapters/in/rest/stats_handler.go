package in

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	out "github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify"
	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/auth"
	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/dto"
	converters "github.com/eduardocoutodev/spotify-stalker/internal/core/converters/in"
)

func HandleTopTracks(w http.ResponseWriter, r *http.Request) {
	tokenManager := auth.GetInstance()

	spotifyToken, err := tokenManager.GetAuthToken()
	if err != nil {
		slog.Error("Error getting auth token", slog.Any("err", err))
		http.Error(w, "Error authenticating with Spotify", http.StatusInternalServerError)
		return
	}

	reqHeaders := make(map[string]string)
	reqHeaders["Authorization"] = "Bearer " + spotifyToken

	timeRange := "long_term"
	if r.URL.Query().Has("time_range") {
		timeRange = r.URL.Query().Get("time_range")
	}

	limit := 10
	if r.URL.Query().Has("limit") {
		limitInt, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			slog.Warn("operation='HandleTopTracks', message='Invalid Number on Limit' ", slog.Any("err", err), slog.Any("userInput", r.URL.Query().Get("limit")))
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Limit query should be a number"})
			return
		}
		limit = limitInt
	}

	offset := 0
	if r.URL.Query().Has("offset") {
		offsetInt, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			slog.Warn("operation='HandleTopTracks', message='Invalid Number on Offset' ", slog.Any("err", err), slog.Any("userInput", r.URL.Query().Get("offset")))
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Offset query should be a number"})
			return
		}
		offset = offsetInt
	}

	resp, err := out.FetchSpotifyWebAPI(
		out.SpotifyRequestArguments{
			Method:              "GET",
			Endpoint:            fmt.Sprintf("https://api.spotify.com/v1/me/top/tracks?time_range=%s&limit=%d&offset=%d", timeRange, limit, offset),
			Headers:             reqHeaders,
			ExpectedStatusCodes: []int{http.StatusOK},
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

	var apiResponse dto.TopTracksSpotifyApiResponse
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
