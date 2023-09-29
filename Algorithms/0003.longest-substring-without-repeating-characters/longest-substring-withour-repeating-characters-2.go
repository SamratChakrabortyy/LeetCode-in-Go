package problem0003

import (
	"math"
	"strings"
)



func lengthOfLongestSubstring2(s string) int {
	n := len(s)
	l, r := 0,0
	freqMap := make(map[string]int)
	sArr := strings.Split(s, "")
	max := 0.0
	for r = 0; r < n; r++ {
		freqMap[sArr[r]] = freqMap[sArr[r]] + 1
		for freqMap[sArr[r]] > 1 && l <= r {
			freqMap[sArr[l]] = freqMap[sArr[l]] - 1
			l++;	
		}
		max = math.Max(max, float64(r - l + 1))
	}
	return int(max)
}