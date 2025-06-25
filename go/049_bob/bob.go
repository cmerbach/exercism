package bob

import (
	"strings"
	"unicode"
)

// Hey determines what Bob will reply based on the input
func Hey(remark string) string {
	// Trim whitespace from the remark
	trimmed := strings.TrimSpace(remark)

	// Check if it's silence (empty or only whitespace)
	if trimmed == "" {
		return "Fine. Be that way!"
	}

	// Check if it's a question (ends with ?)
	isQuestion := strings.HasSuffix(trimmed, "?")

	// Check if it's yelling (has letters and all letters are uppercase)
	hasLetters := false
	isYelling := true

	for _, char := range trimmed {
		if unicode.IsLetter(char) {
			hasLetters = true
			if unicode.IsLower(char) {
				isYelling = false
				break
			}
		}
	}

	// Only consider it yelling if there are letters and all are uppercase
	isYelling = isYelling && hasLetters

	// Determine response based on conditions
	if isQuestion && isYelling {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion {
		return "Sure."
	} else if isYelling {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}
