package reverse

func Reverse(input string) string {
	runes := []rune(input)

	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// func String(s string) (r string) {
// 	for _, v := range s {
// 		r = string(v) + r
// 	}
// 	return
// }
