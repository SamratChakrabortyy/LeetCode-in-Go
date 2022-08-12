package problem0001

func twoSum2(nums []int, target int) []int {
	visited := make(map[int]int, len(nums))
	for i, a := range(nums) {
		if j, ok := visited[target - a]; ok {
			return []int {j, i}
		} 
		visited[a] = i
	}
	return nil
}