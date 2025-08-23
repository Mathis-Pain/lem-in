package checks

func CheckLinks(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && i > 0 && i < len(s)-1 {
			if s[i-1] != s[i+1] {
				return true
			}
		}
	}
	return false
}
