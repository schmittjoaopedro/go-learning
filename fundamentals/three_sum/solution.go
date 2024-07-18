package three_sum

/*
https://leetcode.com/problems/3sum
Start: 7:35
End: 8:00

Original: [-1, 0, 1, 2,-1,-4]
Sorted:   [-4,-1,-1, 0, 1, 2]
*/
import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	combs := map[string]bool{}
	sum3 := [][]int{}
	for i1 := 0; i1 < len(nums)-2; i1++ {
		i2 := i1 + 1
		i3 := len(nums) - 1
		for i2 < i3 {
			sum := nums[i1] + nums[i2] + nums[i3]
			if sum == 0 {
				key := fmt.Sprintf("%d_%d_%d", nums[i1], nums[i2], nums[i3])
				_, ok := combs[key]
				if !ok {
					combs[key] = true
					sum3 = append(sum3, []int{nums[i1], nums[i2], nums[i3]})
				}
				i2++
				i3--
				for i2 < i3 && nums[i2] == nums[i2-1] {
					i2++
				}
				for i2 < i3 && nums[i3] == nums[i3+1] {
					i3--
				}
			} else if sum < 0 {
				i2++
			} else {
				i3--
			}
		}
	}
	return sum3
}

func threeSumBinarySearch(nums []int) [][]int {
	sort.Ints(nums)
	combs := map[string]bool{}
	sum3 := [][]int{}

	for i1 := 0; i1 < len(nums)-2; i1++ {
		for i2 := i1 + 1; i2 < len(nums)-1; i2++ {
			target := -1 * (nums[i1] + nums[i2])
			i3 := binarySearch(nums, target, i2+1, len(nums)-1)
			if i3 != -1 {
				key := fmt.Sprintf("%d_%d_%d", nums[i1], nums[i2], nums[i3])
				_, ok := combs[key]
				if !ok {
					combs[key] = true
					sum3 = append(sum3, []int{nums[i1], nums[i2], nums[i3]})
				}
			}
		}
	}

	return sum3
}

func binarySearch(nums []int, target int, start int, end int) int {
	if start > end {
		return -1
	}
	middle := (end + start) / 2
	if nums[middle] == target {
		return middle
	} else if nums[middle] > target {
		return binarySearch(nums, target, start, middle-1)
	} else {
		return binarySearch(nums, target, middle+1, end)
	}
}
