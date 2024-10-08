package auth

import (
	"fmt"
	"net/http"
)

type SendPasswordResetEmailRequest struct {
	RequestType string `json:"requestType"`
	Email       string `json:"email"`
}

type SendPasswordResetEmailResponse struct {
	Email string `json:"email"`
}

// SendPasswordResetEmail sends a password reset email by issuing an HTTP POST request to the Auth
// getOobConfirmationCode endpoint.
//
// Reference: https://firebase.google.com/docs/reference/rest/auth/#section-send-password-reset-email
func (c *client) SendPasswordResetEmail(req *SendPasswordResetEmailRequest) (*SendPasswordResetEmailResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=%s", c.apiKey)
	return request[SendPasswordResetEmailRequest, SendPasswordResetEmailResponse](c.httpClient, http.MethodPost, url, req)
}
