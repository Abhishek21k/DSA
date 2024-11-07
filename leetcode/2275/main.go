func findMax(arr *[32]int) int {
	maxVal := arr[0]
	for _, value := range arr {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

func largestCombination(candidates []int) int {
	var count [32]int

	for _, val := range candidates {
		i := 0
		for val > 0 {
			count[i] += 1 & val
			i++
			val = val >> 1
		}
	}
	return findMax(&count)
}
