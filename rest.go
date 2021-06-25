package goocord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// TODO: docs and ratelimits

type HTTPRestProvider struct {
	Auth   string
	URL    string
	Client *http.Client
}

func NewHTTPRestProvider(auth string) *HTTPRestProvider {
	return &HTTPRestProvider{
		auth,
		"https://discord.com/api/v9",
		&http.Client{},
	}
}

func (h *HTTPRestProvider) UseAuth(auth string) {
	h.Auth = auth
}

func (h *HTTPRestProvider) UseAPI(url string) {
	h.URL = url
}

func (h *HTTPRestProvider) getURL(endpoint string) string {
	return fmt.Sprintf("%s/%s", h.URL, endpoint)
}

func (h *HTTPRestProvider) convertBody(data interface{}) ([]byte, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (h *HTTPRestProvider) setHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

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
