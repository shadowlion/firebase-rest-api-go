package auth

import (
	"fmt"
	"net/http"
)

type FetchProvidersForEmailRequest struct {
	Identifier  string `json:"identifier"`
	ContinueURI string `json:"continueUri"`
}

type FetchProvidersForEmailResponse struct {
	AllProviders []string `json:"allProviders"`
	Registered   bool     `json:"registered"`
}

// FetchProvidersForEmail looks at all providers associated with a specified email by issuing an HTTP POST request to
// the Auth `createAuthUri` endpoint.
//
// Reference: https://firebase.google.com/docs/reference/rest/auth/#section-send-password-reset-email
func (c *client) FetchProvidersForEmail(req *FetchProvidersForEmailRequest) (*FetchProvidersForEmailResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:createAuthUri?key=%s", c.apiKey)
	return request[FetchProvidersForEmailRequest, FetchProvidersForEmailResponse](c.httpClient, http.MethodPost, url, req)
}
