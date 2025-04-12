package spotify

import (
	"fmt"
	"net/http"
)

type SpotifyRequestArguments struct {
	Method             string
	Endpoint           string
	Headers            map[string]string
	ExpectedStatusCode int
}

func FetchSpotifyWebAPI(requestArguments SpotifyRequestArguments) (*http.Response, error) {
	req, err := http.NewRequest(requestArguments.Method, requestArguments.Endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for %s: %w", requestArguments.Endpoint, err)
	}

	for key, value := range requestArguments.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to %s: %w", requestArguments.Endpoint, err)
	}

	if resp.StatusCode != requestArguments.ExpectedStatusCode {
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code: got %d, expected %d",
			resp.StatusCode, requestArguments.ExpectedStatusCode)
	}

	return resp, err
}
