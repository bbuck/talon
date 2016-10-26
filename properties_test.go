package talon_test

import (
	"time"

	. "github.com/bbuck/talon"
	"github.com/bbuck/talon/types"

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

		Context("with a single property", func() {
			BeforeEach(func() {
				p["one"] = "two"
			})

			It("is a key-insertion pairing", func() {
				Ω(p.String()).Should(Equal(`{one: {one}}`))
			})
		})

		Context("with more than one property", func() {
			BeforeEach(func() {
				p["one"] = "two"
				p["three"] = "four"
			})

			It("is a key-insertion pairing", func() {
				Ω(p.String()).Should(Equal(`{one: {one}, three: {three}}`))
			})
		})

		Context("with conflicting keys during merge", func() {
			BeforeEach(func() {
				b := make(Properties)
				b["one"] = "three"
				p["one"] = "two"
				p = p.Merge(b)
			})

			It("is a key-insertion pairing", func() {
				Ω(p.String()).Should(Equal(`{one: {one}}`))
			})
		})
	})

	Describe("StringWithPostfixedProperties", func() {
		var postfix = "node_a"

		Context("when the property map is empty", func() {
			It("is just an empty string", func() {
				Ω(p.StringWithPostfixedProperties(postfix)).Should(Equal(""))
			})
		})

		Context("with a single property", func() {
			BeforeEach(func() {
				p["one"] = "two"
			})

			It("is a key-insertion pairing", func() {
				Ω(p.StringWithPostfixedProperties(postfix)).Should(Equal(`{one: {one$$node_a}}`))
			})
		})

		Context("with more than one property", func() {
			BeforeEach(func() {
				p["one"] = "two"
				p["three"] = "four"
			})

			It("is a key-insertion pairing", func() {
				Ω(p.StringWithPostfixedProperties(postfix)).Should(Equal(`{one: {one$$node_a}, three: {three$$node_a}}`))
			})
		})

		Context("with conflicting keys during merge", func() {
			BeforeEach(func() {
				b := make(Properties)
				b["one"] = "three"
				p["one"] = "two"
				p = p.Merge(b)
			})

			It("is a key-insertion pairing", func() {
				Ω(p.StringWithPostfixedProperties(postfix)).Should(Equal(`{one: {one$$node_a}}`))
			})
		})
	})

	Describe("ForQuery", func() {
		var (
			pfed         Properties
			postfix      = "node_a"
			val          string
			newOk, oldOk bool
		)

		Context("basic use", func() {
			BeforeEach(func() {
				p["one"] = "two"
				pfed = p.ForQuery(postfix)
				var ival interface{}
				ival, newOk = pfed["one$$node_a"]
				val = ival.(string)
				_, oldOk = pfed["one"]
			})

			It("adds postfix to property names", func() {
				Ω(newOk).Should(BeTrue())
			})

			It("returns the expected value from postfixed key", func() {
				Ω(val).Should(Equal("two"))
			})

			It("does not carry over standard keys", func() {
				Ω(oldOk).Should(BeFalse())
			})
		})

		Context("with marshaled types", func() {
			var (
				t      = time.Now()
				ft     = t.Format(types.DefaultTimeFormat)
				typeOk bool
			)

			BeforeEach(func() {
				p["time"] = t
				pfed = p.ForQuery(postfix)
				var ival interface{}
				ival, newOk = pfed["time$$node_a"]
				val, typeOk = ival.(string)
			})

			It("adds postfix to the key", func() {
				Ω(newOk).Should(BeTrue())
			})

			It("converted time.Time to string", func() {
				Ω(typeOk).Should(BeTrue())
			})

			It("marshaled the time correctly", func() {
				Ω(val).Should(Equal(ft))
			})
		})
	})
})
