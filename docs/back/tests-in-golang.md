# Tests in Golang

In Golang, there are multiple packages that we can use to have tests. In our philsophy we would like to make our tests manually to be able to completely understand what our code really does. That's why we'll be using Gingko and Gomega

## Ginkgo:

Ginkgo is a Behavior Driven Development (BDD) testing framework designed for the Go programming language. It structures tests in a manner that aligns with the principles of behavior-driven development, allowing developers to express expected behaviors in a clear and organized way. Ginkgo provides a framework for writing descriptive test specifications, making it easier to understand and maintain complex test suites.

## Gomega

Gomega, on the other hand, is a powerful and expressive matcher library that seamlessly integrates with Ginkgo. It enhances the testing experience by offering a wide array of matchers, allowing developers to write succinct and readable assertions in their tests. Gomega's rich set of comparison functions covers various data types, contributing to clearer and more maintainable test code. Together with Ginkgo, Gomega forms a comprehensive testing solution that facilitates the creation of well-structured and easily comprehensible test suites in Go.

---

### How to create a test file in Go

In Go, creating a test is easy and simple.

1. Name the file ending with '\_test.go'
2. In the file, you'll need to have these imports :

```go
  "testing"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

```

The '.' dots are really important in Golang, they are very helpful if you want to import everything from this package without having the need to specify which package your importing a function from.

#### Context: Without the dots in import = The wrong way ❌

Your test functions would look like this:

```go
var _ = gingko.Describe("HelloTest", func() {
	gingko.Context("When we want to print hello", func() {
		word, number := HelloWorld()
		gingko.It("should return a string with hello world and 25 as an integer", func() {
			gomega.Expect(word).To(gomega.Equal("HelloWorld"))
			gomega.Expect(number).To(gomega.Equal(13))
		})
	})
})
```

This is very bothersome as you'll need to write `ginkgo` or `gomega` each time.

#### Context: With the dot in your import = The right way ✅

Your test functions would look like this:

```go
var _ = Describe("HelloTest", func() {
	Context("When we want to print hello", func() {
		word, number := HelloWorld()
		It("should return a string with hello world and 25 as an integer", func() {
			Expect(word).To(Equal("HelloWorld"))
			Expect(number).To(Equal(13))
		})
	})
})
```

This is much cleaner, less useless specifications on imports as we already know this is a test file with imports for unit testing.

3. Make the XML file to export the tests results

```go
func TestNameOfFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "NameOfFile", []Reporter{reporters.NewJUnitReporter("test_report_NameOfFile.xml")})
}
```

### How to make a unit test

1. Make your function in another file, the one ending without "\_test.go"

```go
package hello

import (
	"strings"
)

func HelloWorld() (string, int) {
	return "HelloWorld", 25
}
```

2. Test the behavior of your function with a template like this. Here's what a simple test would look like.

```go
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
	RunSpecsWithDefaultAndCustomReporters(t, "Hello", []Reporter{reporters.NewJUnitReporter("test_report_hello.xml")})
}

var _ = Describe("HelloTest", func() {
	Context("When we want to print hello", func() {
		word, number := HelloWorld()
		It("should return a string with hello world and 25 as an integer", func() {
			Expect(word).To(Equal("HelloWorld"))
			Expect(number).To(Equal(13))
		})
	})
})
```

### How to run tests

- To run all tests in your current directory: `go test`
- To run all tests in your current directory and sub-directories: `go test -v ./...`
