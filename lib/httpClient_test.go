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
			client.PrepareGet("https://test.com/someURL", headers)
			Expect(len(client.GetRequests())).To(Equal(1))
			Expect(client.GetRequests()[0].Header.Get("type")).To(Equal("value"))
			Expect(client.GetRequests()[0].Method).To(Equal("GET"))
			Expect(client.GetRequests()[0].URL.Host).To(Equal("test.com"))
			Expect(client.GetRequests()[0].URL.Path).To(Equal("/someURL"))
		})
	})
	Describe("When preparing a POST request", func() {
		It("should return populate a request with headers", func() {
			client := lib.CreateNewClient(&lib.Configuration{})
			headers := map[string]string{"type": "value"}
			params := []byte("param1=value1&param2=value2")
			client.PreparePost("https://test.com/someURL", headers, params)
			Expect(len(client.GetRequests())).To(Equal(1))
			Expect(client.GetRequests()[0].Header.Get("type")).To(Equal("value"))
			Expect(client.GetRequests()[0].Method).To(Equal("POST"))
			Expect(client.GetRequests()[0].URL.Host).To(Equal("test.com"))
			Expect(client.GetRequests()[0].URL.Path).To(Equal("/someURL"))
		})
	})
})
