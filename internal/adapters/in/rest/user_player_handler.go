package in

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/auth"
	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/dto"

	out "github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify"
)

func HandleUserPlayerSeek(w http.ResponseWriter, r *http.Request) {
	var rBody dto.UserPlayerSeekRequest

	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		slog.Warn("Bad request to player seek api", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenManager := auth.GetInstance()

	spotifyToken, err := tokenManager.GetAuthToken()
	if err != nil {
		slog.Error("Error getting auth token", slog.Any("err", err))
		http.Error(w, "Error authenticating with Spotify", http.StatusInternalServerError)
		return
	}

	reqHeaders := make(map[string]string)
	reqHeaders["Authorization"] = "Bearer " + spotifyToken

	resp, err := out.FetchSpotifyWebAPI(
		out.SpotifyRequestArguments{
			Method:              "PUT",
			Endpoint:            fmt.Sprintf("https://api.spotify.com/v1/me/player/seek?position_ms=%d", rBody.NewPositionMs),
			Headers:             reqHeaders,
			ExpectedStatusCodes: []int{http.StatusNoContent, http.StatusOK},
		},
	)

	if err != nil {
		slog.Error("Failed making the request to spotify", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch data from Spotify API"})
		return
	}

	defer resp.Body.Close()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Success on updating to new position"})
}
