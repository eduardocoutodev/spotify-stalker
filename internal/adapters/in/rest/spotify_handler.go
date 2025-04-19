package in

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/auth"
)

// To get refresh token visit: https://spotify-refresh-token-generator.netlify.app
func HandleSpotifyAuthFlowCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	slog.Warn("Got code", slog.Any("Code", code))

	_, refreshToken, err := auth.ExchangeCodeForToken(code)
	if err != nil {
		slog.Error("Error exchanging code for token:", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to exchange token from Spotify API"})
		return
	}

	slog.Info("operation='HandleSpotifyAuthFlowCallback', message='Got refresh Token', ", slog.String("refreshToken", refreshToken))
	w.WriteHeader(http.StatusOK)
}
