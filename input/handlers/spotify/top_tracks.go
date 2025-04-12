package handler_spotify

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/eduardoooxd/spotify-tracker/output/spotify"
)

func HandleTopTracks(w http.ResponseWriter, r *http.Request) {
	spotifyToken := os.Getenv("SPOTIFY_TOKEN")

	reqHeaders := make(map[string]string)
	reqHeaders["Authorization"] = "Bearer " + spotifyToken

	resp, err := spotify.FetchSpotifyWebAPI(
		spotify.SpotifyRequestArguments{
			Method:             "GET",
			Endpoint:           "https://api.spotify.com/v1/me/top/tracks?time_range=long_term&limit=5",
			Headers:            reqHeaders,
			ExpectedStatusCode: http.StatusOK,
		},
	)
	if err != nil {
		log.Fatalf("ERROR making the request to spotify %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ERROR reading body of response: %v", err)
	}
	log.Println(bodyBytes)
}
