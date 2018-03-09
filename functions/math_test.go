package functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dzahariev/gosample/functions"
)

var _ = Describe("Math", func() {
	var (
		a int
		b int
	)
	BeforeEach(func() {
		a = 10
		b = 5
	})

	Describe("Check math funtions", func() {
		Context("When adding", func() {
			It("should be 15", func() {
				Expect(Plus(a, b)).To(Equal(15))
			})
		})

		Context("When substract", func() {
			It("should be 5", func() {
				Expect(Minus(a, b)).To(Equal(5))
			})
		})
	})

})
