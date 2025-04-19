package out

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SpotifyRequestArguments struct {
	Method             string
	Endpoint           string
	Headers            map[string]string
	ExpectedStatusCode int
	Body               url.Values
}

func FetchSpotifyWebAPI(requestArguments SpotifyRequestArguments) (*http.Response, error) {
	var reqBody io.Reader
	if requestArguments.Body != nil {
		reqBody = strings.NewReader(requestArguments.Body.Encode())
	}

	req, err := http.NewRequest(requestArguments.Method, requestArguments.Endpoint, reqBody)
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
