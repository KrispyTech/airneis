package helpers

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Capitalize(text string) string {
	caser := cases.Title(language.AmericanEnglish)

	return caser.String(text)
}
