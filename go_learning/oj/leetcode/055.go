package leetcode

import (
	"fmt"
)

func CanJump() {
	fmt.Println(canJump([]int{3,2,1,0,4}))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	r := len(nums) - 1
	l := r

	for l > 0 {
		l--
		if (r - l <= nums[l]) {
			r = l
		}
		if l == 0 && r - l <= nums[l] {
			return true
		}
	}
	return false
}
