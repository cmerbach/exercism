package luhn

func Valid(id string) bool {
	// Remove spaces from input
	cleaned := ""
	for _, r := range id {
		if r != ' ' {
			cleaned += string(r)
		}
	}

	// Strings of length 1 or less are invalid
	if len(cleaned) <= 1 {
		return false
	}

	sum := 0
	double := false

	// Process digits from right to left
	for i := len(cleaned) - 1; i >= 0; i-- {
		// Check if character is a digit
		if cleaned[i] < '0' || cleaned[i] > '9' {
			return false
		}

		digit := int(cleaned[i] - '0')
		// Double every second digit from the right
		if double {
			digit *= 2
			// If result > 9, subtract 9
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double // Toggle for next iteration
	}

	// Valid if sum is divisible by 10
	return sum%10 == 0
}
