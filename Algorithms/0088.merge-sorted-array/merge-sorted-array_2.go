package problem0088

func merge2(nums1 []int, m int, nums2 []int, n int) {
	// EMPTY nums2
	if n == 0 {
		return
	} else if m == 0 { // EMPTY nums1
		copy(nums1, nums2)
		return
	}
	i, j, k := m-1, n-1, m+n-1

	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
