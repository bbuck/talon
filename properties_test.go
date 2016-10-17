package talon_test

import (
	. "github.com/bbuck/talon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Properties", func() {
	var p Properties

	BeforeEach(func() {
		p = make(Properties)
	})

	Describe("String", func() {
		Context("when the property map is empty", func() {
			It("is just an empty string", func() {
				Ω(p.String()).Should(Equal(""))
			})
		})

		Context("with a single string property", func() {
			BeforeEach(func() {
				p["one"] = "two"
			})

			It("is a key value pairing", func() {
				Ω(p.String()).Should(Equal(`{one: "two"}`))
			})
		})
	})
})
