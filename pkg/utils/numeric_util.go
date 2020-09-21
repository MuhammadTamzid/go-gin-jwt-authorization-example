package utils

import "strconv"

func CastStringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}
