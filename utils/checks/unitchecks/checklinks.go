package unitchecks

import "strings"

func CheckLinks(s string) bool {
	rooms := strings.Split(s, "-")
	if rooms[0] != rooms[1] {
		return true
	} else {
		return false
	}
}
