package auth

import (
	"fmt"
	"net/http"
)

type VerifyPasswordResetCodeRequest struct {
	OOBCode string `json:"oobCode"`
}

type VerifyPasswordResetCodeResponse struct {
	Email       string `json:"email"`
	RequestType string `json:"requestType"`
}

// VerifyPasswordResetCode verify a password reset code by issuing an HTTP POST request to the Auth resetPassword
// endpoint.
//
// Reference: https://firebase.google.com/docs/reference/rest/auth/#section-verify-password-reset-code
func (c *client) VerifyPasswordResetCode(req *VerifyPasswordResetCodeRequest) (*VerifyPasswordResetCodeResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:resetPassword?key=%s", c.apiKey)
	return request[VerifyPasswordResetCodeRequest, VerifyPasswordResetCodeResponse](c.httpClient, http.MethodPost, url, req)
}
