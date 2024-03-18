package helpers

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
)

func Capitalize(text string) string {
	caser := cases.Title(language.AmericanEnglish)

	return caser.String(text)
}

func ConvertStringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}
