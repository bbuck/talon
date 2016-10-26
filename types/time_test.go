// Copyright (c) 2016 Brandon Buck

package types_test

import (
	"time"

	. "github.com/bbuck/talon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TimeType", func() {
	var (
		bs   []byte
		err  error
		t    Time
		date = time.Date(1986, time.November, 12, 1, 2, 3, 4, time.Local)
		test = date.Format(DefaultTimeFormat)
	)

	BeforeEach(func() {
		bs = make([]byte, 0)
		err = nil
	})

	Describe("MarshalTalon", func() {
		BeforeEach(func() {
			t = NewTime(date)
			bs, err = t.MarshalTalon()
		})

		It("doesn't fail", func() {
			Ω(err).Should(BeNil())
		})

		It("produces the correct string", func() {
			Ω(string(bs)).Should(Equal(test))
		})
	})

	Describe("UnmarshalTalon", func() {
		BeforeEach(func() {
			t = Time{OutputFormat: DefaultTimeFormat}
			err = t.UnmarshalTalon([]byte(test))
		})

		It("doesn't fail", func() {
			Ω(err).Should(BeNil())
		})

		It("parsed correct year", func() {
			Ω(t.Time.Year()).Should(Equal(1986))
		})

		It("parsed correct month", func() {
			Ω(t.Time.Month()).Should(Equal(time.November))
		})

		It("parsed the correct day", func() {
			Ω(t.Time.Day()).Should(Equal(12))
		})

		It("parsed the correct hour", func() {
			Ω(t.Time.Hour()).Should(Equal(1))
		})

		It("parsed the correct minute", func() {
			Ω(t.Time.Minute()).Should(Equal(2))
		})

		It("parsed the correct second", func() {
			Ω(t.Time.Second()).Should(Equal(3))
		})

		It("parsed the correct nsec", func() {
			Ω(t.Time.Nanosecond()).Should(Equal(4))
		})
	})
})
