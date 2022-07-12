package strs

import "strings"

func PrefixFold(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(prefix, s[:len(prefix)])
}
