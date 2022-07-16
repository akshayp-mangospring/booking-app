package utils

import "strings"

const minNameLength = 2

func IsZero(count int) bool {
	return count == 0
}

func IsInRange(min, max, i int) bool {
	return i >= min || i <= max
}

func IsNameValid(n string) bool {
	return len(n) > minNameLength
}

func IsEmailValid(str string) bool {
	return strings.Contains(str, "@") && strings.Contains(str, ".")
}
