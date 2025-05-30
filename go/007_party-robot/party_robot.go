package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	hello := fmt.Sprintf("Welcome to my party, %s!", name)
	seat := fmt.Sprintf("You have been assigned to table %03d.", table)
	guest := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	position := fmt.Sprintf("Your table is %s, exactly %.1f meters from here.", direction, distance)
	return hello + "\n" + seat + " " + position + "\n" + guest
}

func main() {
	Welcome("Christiane")
	HappyBirthday("Frank", 58)
	AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
}
