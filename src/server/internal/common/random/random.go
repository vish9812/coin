package random

import (
	"math/rand"
	"strings"
)

// Int returns random integer between range [min, max]
// where max > min >= 0
// If min or max are invalid, returned value is between absolute lesser and greater values
func Int(min, max int) int {
	if min < 0 {
		min *= -1
	}
	if max < 0 {
		max *= -1
	}
	if max < min {
		min, max = max, min
	} else if max == min {
		return max
	}

	return rand.Intn(max-min) + min
}

// String returns a random string of length n
// where n > 0
func String(n int) string {
	if n < 1 {
		return ""
	}

	var sb strings.Builder
	max := int('z') + 1
	min := int('a')

	for i := 0; i < n; i++ {
		ascii := Int(min, max)
		sb.WriteRune(rune(ascii))
	}

	return sb.String()
}
