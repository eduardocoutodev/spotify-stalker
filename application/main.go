package main

import (
	"fmt"
	"log"
	"net/http"

	handler_spotify "github.com/eduardoooxd/spotify-tracker/input/handlers/spotify"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})

	mux.HandleFunc("GET /top/tracks", handler_spotify.HandleTopTracks)
}
