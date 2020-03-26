package leetcode

import (
	"fmt"
)

func SearchMatrix() {
	matrix := [][]int{
		{1,   3,  5,  7},
  		{10, 11, 16, 20},
  		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 3))
}


func searchMatrix(matrix [][]int, target int) bool {
	if (len(matrix) == 0 || len(matrix[0]) == 0) {
		return false;
	}

	firstNums := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		firstNums[i] = matrix[i][0]
	}
	find, index := binarySearch(firstNums, target)
	if find {
		return true
	}
	if (index < 0) {
		return false
	}
	
	find, _ = binarySearch(matrix[index], target)
	return find
}

func binarySearch(nums []int, target int) (bool, int) {
	l, r, direction := 0, len(nums) - 1, 0
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true, mid
		} else if nums[mid] < target {
			l = mid + 1
			direction = 0
		} else {
			r = mid - 1
			direction = 1
		}
	}

	if 0 == direction {
		l--
	}

	if (l < 0) {
		return false, l
	} else if nums[l] < target {
		return false, l
	} else {
		return false, l - 1
	}
}
