package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	// Roll 4 six-sided dice
	dice := make([]int, 4)
	for i := 0; i < 4; i++ {
		dice[i] = rand.Intn(6) + 1 // Random number from 1 to 6
	}

	// Sort the dice to easily find the 3 highest
	sort.Ints(dice)

	// Sum the 3 highest dice (indices 1, 2, 3 after sorting)
	return dice[1] + dice[2] + dice[3]

	// Alternative:
	// return rand.Intn(15) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()

	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
