package lib_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ayoul3/phishkiller/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type jsonData struct {
	Param1 string
	Ip     string
}

var _ = Describe("PrepareData", func() {
	Describe("When data is JSON", func() {
		It("should return proper headers", func() {
			var output jsonData
			headers := map[string]string{"content-type": "application/json"}
			params := []lib.Param{
				{Name: "Param1", Value: "value1"},
				{Name: "Ip", Type: "ipv4"}}
			headers, data := lib.PrepareData(headers, params)
			err := json.Unmarshal(data, &output)
			Expect(len(headers)).To(Equal(1))
			Expect(err).ToNot(HaveOccurred())
			Expect(strings.Count(output.Ip, ".")).To(Equal(3))
		})
	})
	Describe("When data is to be form encoded", func() {
		It("should return proper headers", func() {
			headers := map[string]string{}
			params := []lib.Param{
				{Name: "Param1", Value: "value1"},
				{Name: "ip", Type: "ipv6"}}
			headers, data := lib.PrepareData(headers, params)
			Expect(strings.Count(string(data), "%3A") > 4).To(BeTrue())
			Expect(headers["content-type"]).To(ContainSubstring("x-www-form-urlencoded"))
		})
	})
})

var _ = Describe("GetRandomUA", func() {
	It("should return a different UA each time", func() {
		ua1 := lib.GetRandomUA()
		ua2 := lib.GetRandomUA()
		fmt.Println(ua1)
		fmt.Println(ua2)
		Expect(ua1).ToNot(Equal(ua2))
	})
})
