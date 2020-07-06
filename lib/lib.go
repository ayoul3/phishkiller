package lib

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

const contentTypeHeader = "content-type"

var Chan chan []*http.Request

func ConfigureProxy(proxy string) {
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		log.Warnf("Error configuring proxy %s: %s", proxy, err)
	}
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}
func CreateNewClient(config *Configuration) HttpAPI {
	if config.Proxy != "" {
		log.Infof("Registering Proxy %s", config.Proxy)
		ConfigureProxy(config.Proxy)
	}
	defaultHeaders := map[string]string{
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.6,image/webp,*/*;q=0.5",
		"user-agent":                "Mozilla/5.0 (Windows NT 8.0; Win64; x64; rv:69.0) Gecko/20100115 Firefox/89.85",
		"accept-language":           "en-US,en;q=0.5",
		"accept-encoding":           "gzip, deflate",
		"dnt":                       "1",
		"connection":                "close",
		"upgrade-insecure-requests": "1",
	}
	return &HttpClient{
		Client:  &http.Client{},
		Headers: defaultHeaders,
	}
}

func MergeURLParams(path string, data []byte) string {
	if len(data) > 0 {
		return fmt.Sprintf("%s?%s", path, data)
	}
	return path
}

func PrepareRequests(client HttpAPI, req Request) (*http.Request, error) {
	switch req.Method {
	case "get":
		_, data := PrepareFormData(req.Headers, req.Params)
		url := MergeURLParams(req.Path, data)
		return client.PrepareGet(url, req.Headers), nil
	case "post":
		headers, data := PrepareData(req.Headers, req.Params)
		return client.PreparePost(req.Path, headers, data), nil
	}
	return nil, fmt.Errorf("Method %s not supported for %s. Only get or post", req.Method, req.Path)
}

func PrepareData(headers map[string]string, params []Param) (map[string]string, []byte) {
	contentType := strings.ToLower(headers[contentTypeHeader])
	if strings.Contains(contentType, "json") {
		return PrepareJson(headers, params)
	}
	return PrepareFormData(headers, params)
}

func PrepareJson(headers map[string]string, params []Param) (map[string]string, []byte) {
	values := make(map[string]string)
	for _, p := range params {
		values[p.Name] = GetParamValue(p)
	}
	jsonValue, _ := json.Marshal(values)
	return headers, jsonValue
}

func PrepareFormData(headers map[string]string, params []Param) (map[string]string, []byte) {
	if headers == nil {
		headers = make(map[string]string)
	}
	form := url.Values{}
	for _, p := range params {
		v := GetParamValue(p)
		form.Add(p.Name, v)
	}
	headers[contentTypeHeader] = "application/x-www-form-urlencoded; charset=UTF-8"
	return headers, []byte(form.Encode())
}

func LoopRequests(client HttpAPI, requests []Request) {
	for {
		var preparedReqs []*http.Request
		for _, rawRequest := range requests {
			req, err := PrepareRequests(client, rawRequest)
			if err != nil {
				log.Warn(err)
				continue
			}
			log.Debugf("Prepared request for %s %s", req.Method, req.URL.Path)
			preparedReqs = append(preparedReqs, req)
		}
		Chan <- preparedReqs
	}
}

func Perform(client HttpAPI) {
	for {
		reqs := <-Chan
		log.Infof("sending %d requests", len(reqs))
		client.Perform(reqs)
	}
}

func GetParamValue(p Param) string {
	if p.Value == "" && p.Type != "" {
		return GenerateFake(p.Type)
	}
	return p.Value
}
