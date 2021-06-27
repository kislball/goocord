package goocord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

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
func (h *HTTPRestProvider) convertBody(data interface{}) (d []byte, err error) {
	d, err = json.Marshal(data)
	return
}

// setHeaders takes map of headers and applies them to the req
func (h *HTTPRestProvider) setHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

// transformResponse transforms http.Response into RestResponse
func (h *HTTPRestProvider) transformResponse(resp *http.Response) (ret *RestResponse, err error) {
	headers := make(map[string]string)

	for k, v := range resp.Header {
		headers[k] = strings.Join(v, " ")
	}

	b, err := ioutil.ReadAll(resp.Body)

	var body interface{}
	err = json.Unmarshal(b, &body)

	ret = &RestResponse{
		resp.StatusCode,
		headers,
		body,
	}

	return
}

// Request sends requests to Discord
func (h *HTTPRestProvider) Request(method string, endpoint string, headers map[string]string, body interface{}) (resp *RestResponse, err error) {
	url := h.getURL(endpoint)
	d, err := h.convertBody(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(d))

	h.setHeaders(req, headers)
	req.Header.Set("Authorization", h.Auth)
	req.Header.Set("User-Agent", fmt.Sprintf("DiscordBot (https://github.com/kislball/goocord, %s)", VERSION))

	respRaw, err := h.Client.Do(req)
	if resp.StatusCode == 429 {
		retry, _ := strconv.Atoi(respRaw.Header.Get("Retry-After"))
		time.Sleep(time.Duration(retry) * time.Second)
	}

	resp, err = h.transformResponse(respRaw)

	return
}
