package lib

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"syreclabs.com/go/faker"
)

const contentTypeHeader = "content-type"

var Chan chan bool

func CreateNewClient(config *Configuration) HttpAPI {
	proxyUrl, err := url.Parse(config.Proxy)
	if err != nil {
		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	defaultHeaders := map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.6,image/webp,*/*;q=0.5",
		"User-Agent":                "Mozilla/5.0 (Windows NT 8.0; Win64; x64; rv:69.0) Gecko/20100115 Firefox/89.85",
		"Accept-Language":           "en-US,en;q=0.5",
		"Accept-Encoding":           "gzip, deflate",
		"DNT":                       "1",
		"Connection":                "close",
		"Upgrade-Insecure-Requests": "1",
	}
	return &HttpClient{
		Client:  &http.Client{},
		Headers: defaultHeaders,
	}
}

func PrepareRequests(client HttpAPI, req Request) {
	switch req.Method {
	case "get":
		_, data := PrepareFormData(req.Headers, req.Params)
		client.PrepareGet(fmt.Sprintf("%s?%s", req.Path, data), req.Headers)
	case "post":
		headers, data := PrepareData(req.Headers, req.Params)
		client.PreparePost(req.Path, headers, data)
	}
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
	form := url.Values{}
	for _, p := range params {
		v := GetParamValue(p)
		form.Add(p.Name, v)
	}
	headers[contentTypeHeader] = "application/x-www-form-urlencoded; charset=UTF-8"
	return headers, []byte(form.Encode())
}

func Perform(client HttpAPI) {
	for {
		Chan <- true
		userAgent := GetRandomUA()
		fmt.Println("here")
		client.Perform(userAgent)
		<-Chan
	}
}

func GetParamValue(p Param) string {
	if p.Value == "" && p.Type != "" {
		return GenerateFake(p.Type)
	}
	return p.Value
}

func GetRandomUA() string {
	ff_version := fmt.Sprintf("%d.%d", randRange(69, 82), randRange(69, 99))
	ff_rv := fmt.Sprintf("%d.%d", randRange(58, 99), randRange(0, 9))
	gecko := fmt.Sprintf("20100%03d", randRange(100, 121))
	ua := fmt.Sprintf("Mozilla/5.0 (Windows NT 8.0; Win64; x64; rv:%s) Gecko/%s Firefox/%s", ff_rv, gecko, ff_version)
	return ua
}

func GenerateFake(t string) string {
	switch t {
	case "cardExpiry":
		return fmt.Sprintf("%d/%d", rand.Int31n(13), time.Now().Year()+rand.Intn(3))
	case "cardToken":
		return faker.Numerify("#########")
	case "creditCard":
		return strings.ReplaceAll(faker.Business().CreditCardNumber(), "-", "")
	case "cvv":
		return faker.Numerify("###")
	case "email":
		return faker.Internet().Email()
	case "ipv4":
		return faker.Internet().IpV4Address()
	case "ipv6":
		return faker.Internet().IpV6Address()
	case "ip":
		return fakeIP()
	case "name":
		return faker.Name().Name()
	case "title":
		return faker.Name().Title()
	case "password":
		return faker.Internet().Password(8, 14)
	case "phone":
		return faker.PhoneNumber().CellPhone()
	case "url":
		return faker.Internet().Url()
	case "username":
		return faker.Internet().UserName()
	default:
		return faker.RandomString(10)
	}
}

func fakeEmail() string {
	if rand.Intn(10) < 5 {
		return faker.Internet().Email()
	}
	return faker.Internet().FreeEmail()
}

func fakeIP() string {
	if rand.Intn(10) < 5 {
		return faker.Internet().IpV6Address()
	}
	return faker.Internet().IpV4Address()
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
