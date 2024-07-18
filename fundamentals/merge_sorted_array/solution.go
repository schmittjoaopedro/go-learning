package merge_sorted_array

// https://leetcode.com/problems/merge-sorted-array

func Merge(nums1 []int, m int, nums2 []int, n int) {
	// There's nothing to merge, nums1 is already sorted, just return
	if n <= 0 {
		return
	}

	length := m + n
	for i := length - 1; i >= 0; i-- {
		if m < 1 {
			// nums1 is empty, just copy over all elements form nums2 into nums1
			nums1[i] = nums2[n-1]
			n--
		} else {
			if n > 0 && nums2[n-1] >= nums1[m-1] {
				nums1[i] = nums2[n-1]
				n--
			} else {
				nums1[i] = nums1[m-1]
				m--
			}
		}
	}
}
