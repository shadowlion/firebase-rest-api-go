package auth

import (
	"fmt"
	"net/http"
)

type SignInAnonymouslyRequest struct {
	ReturnSecureToken bool `json:"returnSecureToken"`
}

type SignInAnonymouslyResponse struct {
	IDToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
}

// SignInAnonymously signs in a user anonymously by issuing an HTTP POST request to the Auth `signupNewUser“ endpoint.
//
//	Reference: https://firebase.google.com/docs/reference/rest/auth/#section-sign-in-anonymously
func (c *client) SignInAnonymously(req *SignInAnonymouslyRequest) (*SignInAnonymouslyResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=%s", c.apiKey)
	return request[SignInAnonymouslyRequest, SignInAnonymouslyResponse](c.httpClient, http.MethodPost, url, req)
}
