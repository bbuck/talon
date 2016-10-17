package talon_test

import (
	. "github.com/bbuck/talon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NodeBuilder", func() {
	var n *NodeBuilder

	BeforeEach(func() {
		n = Node()
	})

	It("generates the proper string", func() {
		Ω(n.String()).Should(Equal("()"))
	})

	Describe("Named", func() {
		BeforeEach(func() {
			n.Named("a")
		})

		It("generates the proper string", func() {
			Ω(n.String()).Should(Equal("(a)"))
		})
	})

	Describe("Labeled", func() {
		Context("with one label", func() {
			BeforeEach(func() {
				n.Labeled("One")
			})

			It("generates the proper string", func() {
				Ω(n.String()).Should(Equal("(:One)"))
			})
		})

		Context("with two labels", func() {
			BeforeEach(func() {
				n.Labeled("One", "Two")
			})

			It("generates the proper string", func() {
				Ω(n.String()).Should(Equal("(:One:Two)"))
			})
		})

		Context("with two labels added seperately", func() {
			BeforeEach(func() {
				n.Labeled("One")
				n.Labeled("Two")
			})

			It("generates the proper string", func() {
				Ω(n.String()).Should(Equal("(:One:Two)"))
			})
		})
	})

	Describe("Named + Labeled", func() {
		BeforeEach(func() {
			n.Named("a")
			n.Labeled("One")
		})

		It("generates the proper string", func() {
			Ω(n.String()).Should(Equal("(a:One)"))
		})
	})
})
