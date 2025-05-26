package logs

// Application identifies the application emitting the given log.
func Application(log string) string {

	apps := map[rune]string{
		'❗': "recommendation",	// U+2757
		'🔍': "search",					// U+1F50D
		'☀': "weather",					// U+2600
	}
   
	for _, char := range log {
		if app, exists := apps[char]; exists {
			return app
   	}
  }

		return "default"

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	var newLog string

	for _, char := range log {
		if char == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(char)
		}
	}

	return newLog

	// optional:
	// return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	
	count := 0

	for range log {
		count++
	}

	return count <= limit

	// optional
	// return utf8.RuneCountInString(log) <= limit
}
