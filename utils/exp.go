package utils

const minNameLength = 2

func IsZero(count int) bool {
	return count == 0
}

func IsOutOfRange(min, max, i int) bool {
	return i <= min || i > max
}

func IsNameValid(n string) bool {
	return len(n) < minNameLength
}
