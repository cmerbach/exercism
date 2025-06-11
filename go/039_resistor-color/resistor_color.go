package resistorcolor

// Colors returns a slice of all resistor color names
func Colors() []string {
	return []string{
		"black", "brown", "red", "orange", "yellow",
		"green", "blue", "violet", "grey", "white",
	}
}

// ColorCode returns the numerical value for a given color
func ColorCode(color string) int {
	// colorMap stores the mapping of colors to their numerical values
	var colorMap = map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
	return colorMap[color]
}
