package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	appRune := []rune(log)
    for _, char := range appRune {
		if char == 0x2757 {
            return "recommendation"
        } else if char == 0x1F50D {
            return "search"
        } else if char == 0x2600 {
            return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	actualRune := []rune(log)
    for i, char := range actualRune {
        if char == oldRune {
            actualRune[i] = newRune
        }
    }
    return string(actualRune)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    numberOfRunes := utf8.RuneCountInString(log)
    if numberOfRunes <= limit {
        return true
    }
    return false
}
