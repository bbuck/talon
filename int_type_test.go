package talon_test

import (
	. "github.com/bbuck/talon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IntType", func() {
	var (
		bs  []byte
		err error
	)

	AfterEach(func() {
		bs = make([]byte, 0)
		err = nil
	})

	Describe("Int", func() {
		var (
			input  int64 = 10
			output       = "10"
			i      Int
		)

		AfterEach(func() {
			i = Int(0)
		})

		Describe("MarshalTalon", func() {
			BeforeEach(func() {
				i = Int(input)
				bs, err = i.MarshalTalon()
			})

			It("doesn't fail", func() {
				Ω(err).Should(BeNil())
			})

			It("returns the correct value", func() {
				Ω(string(bs)).Should(Equal(output))
			})
		})

		Describe("UnmarshalTalon", func() {
			Context("with valid input", func() {
				BeforeEach(func() {
					err = i.UnmarshalTalon([]byte(output))
				})

				It("doesn't fail", func() {
					Ω(err).Should(BeNil())
				})

				It("sets itself to the correct value", func() {
					Ω(int64(i)).Should(Equal(input))
				})
			})

			Context("with invalid input", func() {
				BeforeEach(func() {
					err = i.UnmarshalTalon([]byte("aaa"))
				})

				It("fails", func() {
					Ω(err).ShouldNot(BeNil())
				})

				It("doesn't assign a different value", func() {
					Ω(int64(i)).Should(Equal(int64(0)))
				})
			})
		})
	})
})
