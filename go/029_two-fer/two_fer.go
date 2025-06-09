// Package twofer ...
package twofer

// import (
// 	"fmt"
// )

// ShareWith giving back a string to say hello to a customer
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."

	// if name != "" {
	// 	return fmt.Sprintf("One for %s, one for me.", name)
	// }	else {
	// 	return "One for you, one for me."
	// }
}
