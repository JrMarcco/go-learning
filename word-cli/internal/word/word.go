package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderscoreToUpperCamel(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.ToTitle(s)
	return strings.Replace(s, " ", "", -1)
}

func UnderscoreToLowerCamel(s string) string {
	s = UnderscoreToUpperCamel(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
