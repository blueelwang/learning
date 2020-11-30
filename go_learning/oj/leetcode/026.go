package leetcode

func removeDuplicates(nums []int) int {

	i := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}