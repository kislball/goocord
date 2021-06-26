package goocord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// TODO: ratelimits

// HTTPRestProvider is a basic RestProvider used by default.
// Uses HTTP to communicate with Discord's API
type HTTPRestProvider struct {
	Auth   string       // Authentication header used by this HTTPRestProvider
	URL    string       // Base API url
	Client *http.Client // Client used for requests
}

// NewHTTPRestProvider creates a new HTTPRestProvider
func NewHTTPRestProvider(auth string) *HTTPRestProvider {
	return &HTTPRestProvider{
		auth,
		"https://discord.com/api/v9",
		&http.Client{},
	}
}

// UseAuth sets a new Authorization header
func (h *HTTPRestProvider) UseAuth(auth string) {
	h.Auth = auth
}

// UseAPI changes API url
func (h *HTTPRestProvider) UseAPI(url string) {
	h.URL = url
}

// getURL concats endpoint and base API url
func (h *HTTPRestProvider) getURL(endpoint string) string {
	return fmt.Sprintf("%s/%s", h.URL, endpoint)
}

// convertBody transforms data to array of bytes
func (h *HTTPRestProvider) convertBody(data interface{}) ([]byte, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// setHeaders takes map of headers and applies them to the req
func (h *HTTPRestProvider) setHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

// transformResponse transforms http.Response into RestResponse
func (h *HTTPRestProvider) transformResponse(resp *http.Response) (*RestResponse, error) {
	headers := make(map[string]string)

	for k, v := range resp.Header {
		headers[k] = strings.Join(v, " ")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var body interface{}
	err2 := json.Unmarshal(b, &body)
	if err2 != nil {
		return nil, err2
	}

	return &RestResponse{
		resp.StatusCode,
		headers,
		body,
	}, nil
}

// Request sends request to Discord
func (h *HTTPRestProvider) Request(method string, endpoint string, headers map[string]string, body interface{}) (*RestResponse, error) {
	url := h.getURL(endpoint)
	d, err := h.convertBody(body)
	if err != nil {
		return nil, err
	}
	req, errreq := http.NewRequest(method, url, bytes.NewBuffer(d))
	if errreq != nil {
		return nil, errreq
	}
	h.setHeaders(req, headers)
	req.Header.Set("Authorization", h.Auth)
	req.Header.Set("User-Agent", fmt.Sprintf("DiscordBot (https://github.com/kislball/goocord, %s)", VERSION))
	resp, err3 := h.Client.Do(req)
	if err3 != nil {
		return nil, err3
	}

	trresp, err4 := h.transformResponse(resp)
	if err4 != nil {
		return nil, err
	}

	return trresp, nil
}
