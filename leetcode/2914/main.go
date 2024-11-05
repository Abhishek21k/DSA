func minChanges(s string) int {
	n := len(s)
	change := 0
	for i := 0; i < n-1; i += 2 {
		if s[i] != s[i+1] {
			change++
		}
	}
	return change
}
