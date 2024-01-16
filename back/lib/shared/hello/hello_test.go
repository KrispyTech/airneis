package hello_test

import (
	. "airneis/lib/shared/hello"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Hello", []Reporter{reporters.NewJUnitReporter("test_report-hello.xml")})
}

var _ = Describe("HelloTest", func() {
	Context("When we want to print hello", func() {
		word, number := HelloWorld()
		It("should return a string with hello world and 25 as an integer", func() {
			Expect(word).To(Equal("HelloWorld"))
			Expect(number).To(Equal(25))
		})
	})
})
