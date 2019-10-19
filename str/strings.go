package str

import (
	"strings"
)

var words = []string{"Id", "Url"}

func ToAbbreviation(s string) string {
	ret := s
	for _, word := range words {
		ret = matchUpperCase(ret, word)
	}

	return ret
}

// matchUpperCase that to upper case when `word` match with `s` suffix.
// eg: XXXId -> XXXID
func matchUpperCase(s, word string) string {
	idLen := len(word)
	last := len(s)
	idIndex := last - idLen

	if len(s) >= idLen {
		if s[idIndex:] == word {
			return s[:idIndex] + strings.ToUpper(word)
		}
	}

	return s
}
