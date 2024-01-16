package hello_test

import (
	"airneis/lib/shared/hello"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HelloTest", func() {
	Context("When we want to print hello", func() {
		word, number := hello.HelloWorld()
		It("should return a string with hello world and 25 as an integer", func() {
			Expect(word).To(Equal("HelloWorld"))
			Expect(number).To(Equal(28))
		})
	})
})
