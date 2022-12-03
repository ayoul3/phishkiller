package lib

import (
	"bytes"
	"net/http"
	"time"

	"github.com/corpix/uarand"
	"github.com/prometheus/common/log"
)

type HttpAPI interface {
	PrepareGet(url string, extraHeaders map[string]string) *http.Request
	PreparePost(url string, extraHeaders map[string]string, body []byte) *http.Request
	Perform(reqs []*http.Request)
}

type HttpClient struct {
	Client  *http.Client
	Target  string
	Headers map[string]string
}

func (h *HttpClient) PrepareGet(url string, extraHeaders map[string]string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	h.setRandomUA()
	for key, value := range extraHeaders {
		h.Headers[key] = value
	}
	for key, value := range h.Headers {
		req.Header.Set(key, value)
	}
	return req
}

func (h *HttpClient) PreparePost(url string, extraHeaders map[string]string, body []byte) *http.Request {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	h.setRandomUA()
	for key, value := range extraHeaders {
		h.Headers[key] = value
	}
	for key, value := range h.Headers {
		req.Header.Set(key, value)
	}
	return req
}

func (h *HttpClient) Perform(reqs []*http.Request) {
	for _, req := range reqs {
		go func(req *http.Request) {
			resp, err := h.Client.Do(req)
			if err != nil {
				log.Errorf("Error sending reaching %s: %s", req.URL, err)
				return
			}
			if resp.StatusCode > 310 {
				log.Errorf("Got error code %d reaching %s", resp.StatusCode, req.URL)
			}
		}(req)
	}
	time.Sleep(1 * time.Second)
}

func (h *HttpClient) setRandomUA() {
	h.Headers["user-agent"] = uarand.GetRandom()
}
