package auth

import (
	"fmt"
	"net/http"
)

type SignUpWithEmailPasswordRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignUpWithEmailPasswordResponse struct {
	IDToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
	Kind         string `json:"kind"`
}

// SignUpWithEmailPassword creates a new email and password user by issuing an HTTP POST request
// to the Auth `signupNewUser` endpoint.
//
//	Reference: https://firebase.google.com/docs/reference/rest/auth#section-create-email-password
func (c *client) SignUpWithEmailAndPassword(req *SignUpWithEmailPasswordRequest) (*SignUpWithEmailPasswordResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=%s", c.apiKey)
	return request[SignUpWithEmailPasswordRequest, SignUpWithEmailPasswordResponse](c.httpClient, http.MethodPost, url, req)
}
