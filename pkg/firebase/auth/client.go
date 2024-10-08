package auth

import (
	"log"
	"net/http"
	"time"
)

type client struct {
	httpClient *http.Client
	apiKey     string
}

// NewClient returns an http client for Firebase Auth REST API calls.
func NewClient(apiKey string) *client {
	if apiKey == "" {
		log.Fatal("no API key set")
	}

	return &client{
		httpClient: &http.Client{
			Timeout: time.Duration(10 * time.Second),
		},
		apiKey: apiKey,
	}
}
