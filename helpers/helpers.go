package helpers

import (
	"fmt"
	"time"
)

func ContainsAny(string string, chars string) bool {
	for _, char := range string {
		for _, c := range chars {
			if char == c {
				return true
			}
		}
	}
	return false
}

func DurationToString(duration time.Duration) string {
	if duration.Seconds() < 1 {
		return fmt.Sprintf("%d", duration.Milliseconds()) + "ms"
	}
	return fmt.Sprintf("%f", duration.Seconds()) + "s"
}

func IndexOf(string string, character byte) int {
	for i := 0; i < len(string); i++ {
		if string[i] == character {
			return i
		}
	}
	return -1
}
