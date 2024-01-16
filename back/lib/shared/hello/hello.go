package hello

import (
	"strings"
)

func HelloWorld() (string, int) {
	return "HelloWorld", 25
}

func returnLastname(fullName string) string {
	return strings.Split(fullName, " ")[1]
}
