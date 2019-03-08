package helper

import "strconv"

// ParseInt cast string to int
func ParseInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
