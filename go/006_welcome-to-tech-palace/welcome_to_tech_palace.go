package techpalace

import "strings"

const greetings = "Welcome to the Tech Palace, "

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return greetings + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func main() {
	WelcomeMessage("Judy")
	AddBorder("Welcome!", 10)

	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	CleanupMessage(message)
}