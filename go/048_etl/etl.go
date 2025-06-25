package etl

import "strings"

// Transform converts the old scoring format to the new format
// Input: map[int][]string - scores mapped to lists of letters
// Output: map[string]int - individual letters mapped to their scores
func Transform(old map[int][]string) map[string]int {
	result := make(map[string]int)

	// Iterate through each score and its associated letters
	for score, letters := range old {
		// For each letter in the group
		for _, letter := range letters {
			// Convert to lowercase and store with its score
			result[strings.ToLower(letter)] = score
		}
	}

	return result
}
