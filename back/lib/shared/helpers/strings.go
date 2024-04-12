package helpers

import (
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Capitalize(text string) string {
	caser := cases.Title(language.AmericanEnglish)

	return caser.String(text)
}

func ConvertStringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}
