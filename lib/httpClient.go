package lib

import (
	"bytes"
	"net/http"
	"time"

	"github.com/prometheus/common/log"
)

type HttpAPI interface {
	PrepareGet(url string, headers map[string]string)
	PreparePost(url string, headers map[string]string, body []byte)
	Perform(userAgent string)
	GetRequests() []*http.Request
}

type HttpClient struct {
	Client  *http.Client
	Target  string
	Headers map[string]string
	Reqs    []*http.Request
}

func (h *HttpClient) PrepareGet(url string, extraHeaders map[string]string) {
	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range extraHeaders {
		h.Headers[key] = value
	}
	for key, value := range h.Headers {
		req.Header.Set(key, value)
	}
	h.Reqs = append(h.Reqs, req)
}

func (h *HttpClient) PreparePost(url string, extraHeaders map[string]string, body []byte) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	for key, value := range extraHeaders {
		h.Headers[key] = value
	}
	for key, value := range h.Headers {
		req.Header.Set(key, value)
	}
	h.Reqs = append(h.Reqs, req)
}

func (h *HttpClient) Perform(userAgent string) {
	log.Info("sending requests")
	for _, req := range h.Reqs {
		req.Header.Set("User-Agent", userAgent)
		_, err := h.Client.Do(req)
		if err != nil {
			log.Errorf("Error sending reaching %s: %s", req.URL, err)
		}
	}
	time.Sleep(1 * time.Second)
}
func (h *HttpClient) GetRequests() []*http.Request {
	return h.Reqs
}
