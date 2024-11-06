func canSortArray(nums []int) bool {
	n := len(nums) - 1
	for j := 0; j < n; j++ {
		for i := 0; i <= n-j-1; i++ {
			if nums[i] > nums[i+1] {
				if countBits(nums[i]) == countBits(nums[i+1]) {
					nums[i] += nums[i+1]
					nums[i+1] = nums[i] - nums[i+1]
					nums[i] = nums[i] - nums[i+1]
				} else {
					return false
				}
			}
		}
	}
	return true
}

func countBits(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}
