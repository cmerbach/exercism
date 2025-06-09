package parsinglogfiles

import (
	"regexp"
	"fmt"
)

func IsValidLine(text string) bool {
	pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	matched, _ := regexp.MatchString(pattern, text)
	return matched
}

func SplitLogLine(text string) []string {
	pattern := `<[~*=-]*>`
	re := regexp.MustCompile(pattern)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	pattern := `"[^"]*password[^"]*"`
	re := regexp.MustCompile(`(?i)` + pattern) // (?i) macht es case-insensitive
	
	for _, line := range lines {
			if re.MatchString(line) {
					count++
			}
	}
	
	return count
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d+`
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	result := []string{}

	for _, line := range lines {
		founds := re.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		result = append(result, line)
	}
	
	return result
}
