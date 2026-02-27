package strings_helpers


import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleCase(str string) string {
	caser := cases.Title(language.English)
	return caser.String(str)
}

func LowerCase(str string) string {
	caser := cases.Lower(language.English)
	return caser.String(str)
}

func ContainsIgnoreCase(strs []string, str string) bool {
	lower_str := LowerCase(str)
	for _, s := range strs {
		if LowerCase(s) == lower_str {
			return true
		}
	}
	return false
}