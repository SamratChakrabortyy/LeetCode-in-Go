package problem0027

func removeElement2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			if nums[j] != val {
				nums[i] = nums[j]
				i++
			}
			j--
		} else {
			i++
		}
	}
	return i
}
