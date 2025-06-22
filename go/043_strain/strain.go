package strain

// Generics Syntax erklärt:
// [T any] bedeutet: T ist ein Type Parameter, der jeden Typ (any) annehmen kann
// []T bedeutet: eine Slice vom Typ T
// func(T) bool bedeutet: eine Funktion die T nimmt und bool zurückgibt

// Keep returns a new slice containing elements where the predicate returns true
func Keep[T any](collection []T, predicate func(T) bool) []T {
	var result []T // result ist eine Slice vom gleichen Typ T

	for _, value := range collection {
		if predicate(value) {
			result = append(result, value)
		}
	}

	return result
}

// Discard returns a new slice containing elements where the predicate returns false
func Discard[T any](collection []T, predicate func(T) bool) []T {
	var result []T

	for _, value := range collection {
		if !predicate(value) { // Negiere das Prädikat
			result = append(result, value)
		}
	}

	return result
}
