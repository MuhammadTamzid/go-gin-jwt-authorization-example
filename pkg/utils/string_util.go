package utils

import "strings"

func JoinString(sep string, strs ...string) string  {
	if len(strs) < 2 {
		return strs[0]
	}
	return strings.Join(strs, sep)
}
