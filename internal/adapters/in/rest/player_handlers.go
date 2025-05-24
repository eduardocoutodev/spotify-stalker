package in

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	out "github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify"
	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/auth"
	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/dto"
)

func HandleResumeMusic(w http.ResponseWriter, r *http.Request) {
	tokenManager := auth.GetInstance()

	spotifyToken, err := tokenManager.GetAuthToken()
	if err != nil {
		slog.Error("Error getting auth token", slog.Any("err", err))
		http.Error(w, "Error authenticating with Spotify", http.StatusInternalServerError)
		return
	}

	reqHeaders := make(map[string]string)
	reqHeaders["Authorization"] = "Bearer " + spotifyToken

	baseUrl, err := url.Parse("https://api.spotify.com/v1/me/player/play")
	if err != nil {
		slog.Error("Error parsing base url", slog.Any("baseUrl", baseUrl), slog.Any("err", err))
		http.Error(w, "Error parsing base URL to resume music", http.StatusInternalServerError)
		return
	}

	deviceId := r.URL.Query().Get("deviceId")
	if deviceId != "" {
		q := baseUrl.Query()
		q.Set("device_id", deviceId)
		baseUrl.RawQuery = q.Encode()
	}

	resumeEndpoint := baseUrl.String()
	slog.Debug("Resume endpoint generated", slog.String("resumeEndpoint", resumeEndpoint))
	resp, err := out.FetchSpotifyWebAPI(
		out.SpotifyRequestArguments{
			Method:              "PUT",
			Endpoint:            resumeEndpoint,
			Headers:             reqHeaders,
			ExpectedStatusCodes: []int{http.StatusOK, http.StatusForbidden},
		},
	)

	if err != nil {
		slog.Error("Failed making the request to spotify", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to resume music from Spotify API"})
		return
	}

	defer resp.Body.Close()

	// Flow for bad request (trying to stop, when it's already resumed)
	if resp.StatusCode == http.StatusForbidden {
		slog.Warn("Client tried to resume music when is already playing")
		// Ideally this should be a http.StatusConflict, but I will keep consistent with spotify api status code
		w.WriteHeader(http.StatusForbidden)
		errorResponse := dto.ErrorResponse{
			ErrorMessage: "Music is already resumed",
		}
		json.NewEncoder(w).Encode(&errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Success on resuming the player"})
}

func HandlePauseMusic(w http.ResponseWriter, r *http.Request) {
	tokenManager := auth.GetInstance()

	spotifyToken, err := tokenManager.GetAuthToken()
	if err != nil {
		slog.Error("Error getting auth token", slog.Any("err", err))
		http.Error(w, "Error authenticating with Spotify", http.StatusInternalServerError)
		return
	}

	reqHeaders := make(map[string]string)
	reqHeaders["Authorization"] = "Bearer " + spotifyToken

	baseUrl, err := url.Parse("https://api.spotify.com/v1/me/player/pause")
	if err != nil {
		slog.Error("Error parsing base url", slog.Any("baseUrl", baseUrl), slog.Any("err", err))
		http.Error(w, "Error parsing base URL to pause music", http.StatusInternalServerError)
		return
	}

	deviceId := r.URL.Query().Get("deviceId")
	if deviceId != "" {
		q := baseUrl.Query()
		q.Set("device_id", deviceId)
		baseUrl.RawQuery = q.Encode()
	}

	pauseEndpoint := baseUrl.String()
	slog.Debug("Pause endpoint generated", slog.String("pauseEndpoint", pauseEndpoint))
	resp, err := out.FetchSpotifyWebAPI(
		out.SpotifyRequestArguments{
			Method:              "PUT",
			Endpoint:            pauseEndpoint,
			Headers:             reqHeaders,
			ExpectedStatusCodes: []int{http.StatusOK, http.StatusForbidden},
		},
	)

	if err != nil {
		slog.Error("Failed making the request to spotify", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to pause music from Spotify API"})
		return
	}

	defer resp.Body.Close()

	// Flow for bad request (trying to stop, when it's already stopped)
	if resp.StatusCode == http.StatusForbidden {
		slog.Warn("Client tried to stop music already stopped")
		// Ideally this should be a http.StatusConflict, but I will keep consistent with spotify api status code
		w.WriteHeader(http.StatusForbidden)
		errorResponse := dto.ErrorResponse{
			ErrorMessage: "Music is already stopped",
		}
		json.NewEncoder(w).Encode(&errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Success on pausing the player"})
}
