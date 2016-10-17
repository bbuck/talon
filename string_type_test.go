package talon_test

import (
	. "github.com/bbuck/talon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String Type", func() {
	var (
		bs  []byte
		err error
	)

	AfterEach(func() {
		bs = make([]byte, 0)
		err = nil
	})

	Describe("String", func() {
		var (
			input  = "test"
			output = `"test"`
			s      String
		)

		AfterEach(func() {
			s = String("")
		})

		Describe("MarshalTalon", func() {
			BeforeEach(func() {
				s = String(input)
				bs, err = s.MarshalTalon()
			})

			It("doesn't fail", func() {
				Ω(err).Should(BeNil())
			})

			It("returns the correct string, with quotes", func() {
				Ω(string(bs)).Should(Equal(output))
			})
		})

		Describe("UnmarshalTalon", func() {
			BeforeEach(func() {
				err = s.UnmarshalTalon([]byte(input))
			})

			It("doesn't fail", func() {
				Ω(err).Should(BeNil())
			})

			It("sets it's value correctly", func() {
				Ω(string(s)).Should(Equal(input))
			})
		})
	})
})
