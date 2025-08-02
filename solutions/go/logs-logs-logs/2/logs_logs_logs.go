package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	logRune := []rune(log)
    for _, char := range logRune {
		switch char{
            case 0x2757 :
            	return "recommendation"
            case 0x1F50D :
            	return "search"
            case 0x2600 :
            	return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	logRune := []rune(log)
    for i, char := range logRune {
        if char == oldRune {
            logRune[i] = newRune
        }
    }
    return string(logRune)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    numberOfRunes := utf8.RuneCountInString(log)
    return numberOfRunes <= limit
}
