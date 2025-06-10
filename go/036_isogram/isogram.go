package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	letter := make(map[rune]struct{})
	
	for _, char := range strings.ToLower(word) {
		if char == ' ' || char == '-' {
			continue
		}
		
		if _, exists := letter[char]; exists {
			return false
		}
		
		letter[char] = struct{}{}
	}
	
	return true
}
