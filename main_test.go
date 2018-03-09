package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dzahariev/gosample"
)

var _ = Describe("Main", func() {
	var (
		dummystring string
	)
	BeforeEach(func() {
		dummystring = "dummystring"
	})

	Describe("Check main funtions", func() {
		Context("Dummy function", func() {
			It("should be dummystring", func() {
				Expect(Dummy("dummystring")).To(Equal("dummystring"))
			})
		})
	})

})
