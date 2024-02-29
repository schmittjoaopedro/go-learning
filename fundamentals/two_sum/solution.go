package two_sum

func TwoSum(nums []int, target int) []int {

	var index = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		idx, ok := index[target-nums[i]]
		if ok {
			return []int{idx, i}
		} else {
			index[nums[i]] = i
		}
	}

	return nil
}
