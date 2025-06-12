package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	// Initialize histogram with all valid nucleotides set to 0
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	// Count each nucleotide in the DNA string
	for _, nucleotide := range d {
		// Check if nucleotide is valid
		if nucleotide != 'A' && nucleotide != 'C' && nucleotide != 'G' && nucleotide != 'T' {
			return nil, fmt.Errorf("invalid nucleotide: %c", nucleotide)
		}

		// Increment count for this nucleotide
		h[nucleotide]++
	}

	return h, nil
}
