package other

import "strings"

func ReplaceSpace(s string) string {
	return strings.Join(strings.Split(s, " "), "%20")
}
