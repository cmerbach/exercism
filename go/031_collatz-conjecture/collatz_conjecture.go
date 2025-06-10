package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {

	// Handle invalid input
	if n <= 0 {
		return 0, errors.New("only positive integers are allowed")
	}

	steps := 0
	
	// Continue until we reach 1
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	
	return steps, nil
}
