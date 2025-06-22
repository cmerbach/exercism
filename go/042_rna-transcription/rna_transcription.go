package strand

import "strings"

func ToRNA(dna string) string {
	return strings.NewReplacer("G", "C", "C", "G", "T", "A", "A", "U").Replace(dna)
}
