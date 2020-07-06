package lib

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

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
	log.Info("sending requests")
	for _, req := range reqs {
		_, err := h.Client.Do(req)
		if err != nil {
			log.Errorf("Error sending reaching %s: %s", req.URL, err)
		}
	}
	time.Sleep(1 * time.Second)
}

func (h *HttpClient) setRandomUA() {
	ff_version := fmt.Sprintf("%d.%d", randRange(69, 82), randRange(69, 99))
	ff_rv := fmt.Sprintf("%d.%d", randRange(58, 99), randRange(0, 9))
	gecko := fmt.Sprintf("20100%03d", randRange(100, 121))
	ua := fmt.Sprintf("Mozilla/5.0 (Windows NT 8.0; Win64; x64; rv:%s) Gecko/%s Firefox/%s", ff_rv, gecko, ff_version)
	h.Headers["user-agent"] = ua
}
