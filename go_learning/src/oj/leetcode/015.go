package leetcode

import (
	"sort"
)


func threeSum(nums []int) [][]int {
	sort.Ints(nums)
    res := [][]int{}
    for i := range nums {
        if i == 0 || nums[i]>nums[i-1] {
			l := i+1
			r := len(nums)-1
            for l < r {
                s := nums[i] + nums[l] +nums[r]
                if s ==0 {
                    res = append(res, []int{nums[i],nums[l],nums[r]})
                    l +=1
                    r -=1
                    for l < r && nums[l] == nums[l-1] {
						l += 1
					}
                    for r > l && nums[r] == nums[r+1] {
						r -= 1
					}
				} else if s>0 {
					r -=1
				} else {
					l +=1
				}
			}
		}
	}
    return res
}