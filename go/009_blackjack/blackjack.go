package blackjack

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
		case "ace":
			return 11
		case "king":
			return 10
		case "queen":
			return 10
		case "jack":
			return 10
		case "ten":
			return 10
		case "nine":
			return 9
		case "eight":
			return 8
		case "seven":
			return 7
		case "six":
			return 6
		case "five":
			return 5
		case "four":
			return 4
		case "three":
			return 3
		case "two":
			return 2
		default:
			return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    playerTotal := ParseCard(card1) + ParseCard(card2)
    dealerValue := ParseCard(dealerCard)
    
    switch {
		case ParseCard(card1) == 11 && ParseCard(card2) == 11:
			return "P"
		
		case playerTotal == 21:
			switch {
				case dealerValue >= 10:
					return "S"
				default:
					return "W"
			}
		
		case playerTotal >= 17:
			return "S"
		
		case playerTotal >= 12:
			switch {
				case dealerValue >= 7:
					return "H"
				default:
					return "S"
			}
		
		default:
			return "H"
    }
}

func main() {
	fmt.Println(ParseCard("ace"))
	FirstTurn("ace", "ace", "jack")
}
