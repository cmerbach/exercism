package grains

import "errors"

// Square calculates the number of grains on a given square
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("square must be between 1 and 64")
	}

	return 1 << (number - 1), nil
}

// Total calculates the total number of grains on the entire chessboard
func Total() uint64 {
	return (1 << 64) - 1
}
