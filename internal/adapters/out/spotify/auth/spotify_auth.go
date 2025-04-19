package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	out "github.com/eduardocoutodev/spotify-stalker/internal/adapters/out/spotify"
)

var instance *TokenManager
var once sync.Once

func GetInstance() *TokenManager {
	once.Do(func() {
		instance = &TokenManager{
			refreshToken: os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		}
	})
	return instance
}

type TokenManager struct {
	accessToken  string
	refreshToken string
	expiresAt    time.Time
	mutex        sync.Mutex
}

func (tm *TokenManager) GetAuthToken() (string, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	if tm.accessToken != "" && time.Now().Before(tm.expiresAt) {
		return tm.accessToken, nil
	}

	token, expiresIn, err := refreshAccessToken(tm.refreshToken)
	if err != nil {
		return "", fmt.Errorf("error refreshing token: %v", err)
	}

	tm.accessToken = token
	// Refresh 5 minutes before expiring
	tm.expiresAt = time.Now().Add(time.Duration(expiresIn-300) * time.Second)
	return token, nil
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

func refreshAccessToken(refreshToken string) (string, int, error) {
	reqHeaders := make(map[string]string)

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	reqHeaders["Authorization"] = "Basic " + auth
	reqHeaders["Content-Type"] = "application/x-www-form-urlencoded"

	reqBody := url.Values{}
	reqBody.Set("grant_type", "refresh_token")
	reqBody.Set("refresh_token", refreshToken)

	resp, err := out.FetchSpotifyWebAPI(
		out.SpotifyRequestArguments{
			Method:             "POST",
			Endpoint:           "https://accounts.spotify.com/api/token",
			Headers:            reqHeaders,
			ExpectedStatusCode: http.StatusOK,
			Body:               reqBody,
		},
	)

	if err != nil {
		slog.Error("Failed making the request to spotify", slog.Any("err", err))
		return "", 0, fmt.Errorf("error creating token request: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed reading response body", slog.Any("err", err))
		return "", 0, err
	}

	var apiResponse tokenResponse
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		slog.Error("Failed reading response body", slog.Any("err", err))
		return "", 0, err
	}

	return apiResponse.AccessToken, apiResponse.ExpiresIn, nil
}

func ExchangeCodeForToken(code string) (string, string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectUri := os.Getenv("SPOTIFY_REDIRECT_URI")

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", "", fmt.Errorf("error creating token request: %v", err)
	}

	// Set headers
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("error requesting token: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", "", fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error reading token response: %v", err)
	}

	var tokenResponse tokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", "", fmt.Errorf("error parsing token response: %v", err)
	}

	return tokenResponse.AccessToken, tokenResponse.RefreshToken, nil
}
