package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// HTTPHelper is a utility for making HTTP requests
type HTTPHelper struct {
	Client *http.Client
}

// NewHTTPHelper creates a new HTTPHelper with a default timeout
func NewHTTPHelper(timeout time.Duration) *HTTPHelper {
	return &HTTPHelper{
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Post performs an HTTP POST request
func (h *HTTPHelper) Post(url string, body []byte, headers map[string]string) ([]byte, int, error) {
	return h.doRequest("POST", url, body, headers)
}

// Get performs an HTTP GET request
func (h *HTTPHelper) Get(url string, headers map[string]string) ([]byte, int, error) {
	return h.doRequest("GET", url, nil, headers)
}

// Put performs an HTTP PUT request
func (h *HTTPHelper) Put(url string, body []byte, headers map[string]string) ([]byte, int, error) {
	return h.doRequest("PUT", url, body, headers)
}

func (h *HTTPHelper) Path(url string, body []byte, headers map[string]string) ([]byte, int, error) {
	return h.doRequest("PATH", url, body, headers)
}

// Delete performs an HTTP DELETE request
func (h *HTTPHelper) Delete(url string, body []byte, headers map[string]string) ([]byte, int, error) {
	return h.doRequest("DELETE", url, body, headers)
}

// doRequest is a helper method for making HTTP requests
func (h *HTTPHelper) doRequest(method, url string, body []byte, headers map[string]string) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return responseBody, resp.StatusCode, nil
}

// BuildQueryParams builds a query string from a map
func BuildQueryParams(params any) string {
	query := url.Values{}
	//for key, value := range params {
	//	query.Add(key, value)
	//}
	return query.Encode()
}
