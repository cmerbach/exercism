package space

// Planet represents a planet with its orbital period in Earth years
type Planet string

const (
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Earth   Planet = "Earth"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// Orbital periods in Earth years
var orbitalPeriods = map[Planet]float64{
	Mercury: 0.2408467,
	Venus:   0.61519726,
	Earth:   1.0,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132,
}

// Earth year in seconds (365.25 days)
const earthYearSeconds = 31557600

// Age calculates the age on a given planet based on age in seconds
func Age(seconds float64, planet Planet) float64 {
	// Check if planet exists in our map
	orbitalPeriod, exists := orbitalPeriods[planet]
	if !exists {
		return -1
	}

	// Convert seconds to Earth years
	earthYears := seconds / earthYearSeconds

	// Convert Earth years to planet years
	return earthYears / orbitalPeriod
}

// func Age(seconds float64, planet Planet) float64 {
// 	earthAge := seconds / 31557600
// 	switch planet {
// 	case "Mercury":
// 		return earthAge / 0.2408467
// 	case "Venus":
// 		return earthAge / 0.61519726
// 	case "Earth":
// 		return earthAge
// 	case "Mars":
// 		return earthAge / 1.8808158
// 	case "Jupiter":
// 		return earthAge / 11.862615
// 	case "Saturn":
// 		return earthAge / 29.447498
// 	case "Uranus":
// 		return earthAge / 84.016846
// 	case "Neptune":
// 		return earthAge / 164.79132
// 	default:
// 		return -1
// 	}
// }
