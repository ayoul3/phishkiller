package lib_test

import (
	"github.com/ayoul3/phishkiller/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HttpClient", func() {
	Describe("When preparing a GET request", func() {
		It("should return populate a request with headers", func() {
			client := lib.CreateNewClient(&lib.Configuration{})
			headers := map[string]string{"type": "value"}
			req := client.PrepareGet("https://test.com/someURL", headers)

			Expect(req.Header.Get("type")).To(Equal("value"))
			Expect(req.Method).To(Equal("GET"))
			Expect(req.URL.Host).To(Equal("test.com"))
			Expect(req.URL.Path).To(Equal("/someURL"))
		})
	})
	Describe("When preparing a POST request", func() {
		It("should return populate a request with headers", func() {
			client := lib.CreateNewClient(&lib.Configuration{})
			headers := map[string]string{"type": "value"}
			params := []byte("param1=value1&param2=value2")
			req := client.PreparePost("https://test.com/someURL", headers, params)
			Expect(req.Header.Get("type")).To(Equal("value"))
			Expect(req.Method).To(Equal("POST"))
			Expect(req.URL.Host).To(Equal("test.com"))
			Expect(req.URL.Path).To(Equal("/someURL"))
		})
	})
})
