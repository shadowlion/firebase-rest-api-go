package auth

import (
	"fmt"
	"net/http"
)

type SignInWithEmailPasswordRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignInWithEmailPasswordResponse struct {
	IDToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
	Registered   bool   `json:"registered"`
}

// SignInWithEmailPassword signs in a user with an email and password by issuing an HTTP POST request to the Auth
// `signInWithPassword` endpoint.
//
// Reference: https://firebase.google.com/docs/reference/rest/auth#section-create-email-password
func (c *client) SignInWithEmailAndPassword(req *SignInWithEmailPasswordRequest) (*SignInWithEmailPasswordResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", c.apiKey)
	return request[SignInWithEmailPasswordRequest, SignInWithEmailPasswordResponse](c.httpClient, http.MethodPost, url, req)
}
