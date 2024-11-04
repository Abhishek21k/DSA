func compressedString(word string) string {
	n := len(word)

	if n == 0 {
		return ""
	}

	var comp strings.Builder
	sum := 1

	for i := 1; i < n; i++ {
		if word[i-1] == word[i] && sum < 9 {
			sum++
		} else {
			comp.WriteString(strconv.Itoa(sum))
			comp.WriteByte(word[i-1])
			sum = 1
		}

	}
	comp.WriteString(strconv.Itoa(sum))
	comp.WriteByte(word[n-1])
	return comp.String()

}