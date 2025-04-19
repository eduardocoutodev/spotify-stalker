package main

import (
	"fmt"
	"log/slog"
	"net/http"

	in "github.com/eduardocoutodev/spotify-stalker/internal/adapters/in/rest"
	middlewares "github.com/eduardocoutodev/spotify-stalker/internal/adapters/in/rest/middleware"
	"github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify/auth"
	"github.com/eduardocoutodev/spotify-stalker/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", slog.Any("err", err))
	}

	tokenManager := auth.GetInstance()
	_, err = tokenManager.GetAuthToken()
	if err != nil {
		slog.Error("Error getting auth spotify token", slog.Any("err", err))
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", in.HandleHealthCheck)
	mux.HandleFunc("GET /spotify/callback", in.HandleSpotifyAuthFlowCallback)
	mux.HandleFunc("GET /stats/tracks", in.HandleTopTracks)
	mux.HandleFunc("GET /user/music/current", in.HandleUserCurrentPlaying)

	handler := middlewares.JsonContentTypeMiddleware(mux)

	port := config.GetEnv("SERVER_PORT", "8080")
	addressToListen := fmt.Sprintf(":%s", port)
	slog.Info("Started HTTP Server", slog.String("address", addressToListen))
	if err := http.ListenAndServe(addressToListen, handler); err != nil {
		slog.Error("Error starting HTTP server", slog.Any("err", err))
	}

}
