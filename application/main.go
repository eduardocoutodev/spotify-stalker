package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/eduardoooxd/spotify-tracker/commons"
	handler_spotify "github.com/eduardoooxd/spotify-tracker/input/handlers/spotify"
	"github.com/eduardoooxd/spotify-tracker/input/middlewares"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", slog.Any("err", err))
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})
	mux.HandleFunc("GET /top/tracks", handler_spotify.HandleTopTracks)

	handler := middlewares.JsonContentTypeMiddleware(mux)

	port := commons.GetEnv("SERVER_PORT", "8080")
	addressToListen := fmt.Sprintf(":%s", port)
	slog.Info("Started HTTP Server", slog.String("address", addressToListen))
	if err := http.ListenAndServe(addressToListen, handler); err != nil {
		slog.Error("Error starting HTTP server", slog.Any("err", err))
	}
}
