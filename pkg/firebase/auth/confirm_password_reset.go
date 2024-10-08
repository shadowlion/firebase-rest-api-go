package auth

import (
	"fmt"
	"net/http"
)

type ConfirmPasswordResetRequest struct {
	OOBCode     string `json:"oobCode"`
	NewPassword string `json:"newPassword"`
}

type ConfirmPasswordResetResponse struct {
	Email       string `json:"email"`
	RequestType string `json:"requestType"`
}

// ConfirmPasswordReset applies a password reset change by issuing an HTTP POST request to the Auth resetPassword endpoint.
//
// Reference: https://firebase.google.com/docs/reference/rest/auth/#section-confirm-reset-password
func (c *client) ConfirmPasswordReset(req *ConfirmPasswordResetRequest) (*ConfirmPasswordResetResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:resetPassword?key=%s", c.apiKey)
	return request[ConfirmPasswordResetRequest, ConfirmPasswordResetResponse](c.httpClient, http.MethodPost, url, req)
}
