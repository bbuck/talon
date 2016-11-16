package types_test

import (
	. "github.com/bbuck/talon/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Complex", func() {
	Describe("NewComplex", func() {
		var (
			zero complex128 = 0i
			c64  complex64  = 1 + 2i
			c128 complex128 = 3 + 6i
		)

		Context("when fed a complex64", func() {
			var (
				c64to128 = complex128(c64)
				c        = NewComplex(c64)
			)

			It("returns a wrapped complex128", func() {
				Ω(complex128(c)).Should(Equal(c64to128))
			})
		})

		Context("when fed a complex128", func() {
			c := NewComplex(c128)

			It("wraps the complex128", func() {
				Ω(complex128(c)).Should(Equal(c128))
			})
		})

		Context("when fed a non-complex value", func() {
			c := NewComplex("")

			It("returns a zero complex", func() {
				Ω(complex128(c)).Should(Equal(zero))
			})
		})
	})
})
