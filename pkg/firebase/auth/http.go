package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// makeRequestBody simply checks if we're running a GET request. Otherwise, we
// return a request with a []byte body payload attached.
func makeRequestBody(method, url string, requestBody []byte) (*http.Request, error) {
	switch method {
	case http.MethodGet:
		return http.NewRequest(method, url, nil)
	default:
		return http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	}
}

// do attaches a `Content-Type` header to the request before sending.
func do(client *http.Client, req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}

// parseResponse conforms the response body to one of three elements: an okay response, the standard
// firebase error response given in the documentation, or an error related to creating the request.
func parseResponse[ResponseBody interface{}](resp *http.Response) (*ResponseBody, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if resp.StatusCode >= 400 {
		var errResponse ErrorResponse

		if err := json.Unmarshal(body, &errResponse); err != nil {
			return nil, errors.New("error parsing ERROR response body")
		}

		return nil, errResponse.ToCustomError()
	}

	var okResponse ResponseBody

	if err := json.Unmarshal(body, &okResponse); err != nil {
		return nil, errors.New("error parsing OK response body")
	}

	return &okResponse, nil
}

// request creates a generic http request that includes a payload (for our purposes here, we're mainly focused on
// GET, POST, and PATCH)
func request[RequestBody, ResponseBody interface{}](client *http.Client, method, url string, payload *RequestBody) (*ResponseBody, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := makeRequestBody(method, url, rb)
	if err != nil {
		return nil, err
	}

	resp, err := do(client, req)
	if err != nil {
		return nil, err
	}

	return parseResponse[ResponseBody](resp)
}
