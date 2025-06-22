package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)

	for i := 'a'; i <= 'z'; i++ {
		if !strings.ContainsRune(input, i) {
			return false
		}
	}

	return true

	// m := make(map[rune]bool)
	// for _, r := range input {
	// 	if unicode.IsLetter(r) {
	// 		m[unicode.ToLower(r)] = true
	// 	}
	// }
	// return len(m) == 26
}
