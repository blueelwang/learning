package leetcode

import (
	"math"
	"fmt"
)

func MaxArea() {
	area := maxArea([]int{1,8,6,2,5,4,8,3,7})
	fmt.Println(area)
}

func maxArea(height []int) int {
	maxArea := 0
	l, r := 0, len(height) - 1
	for l < r {
		area := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
		if (area > maxArea) {
			maxArea = area
		}
		if (height[l] > height[r]) {
			r--
		} else {
			l++
		}
	}
	return maxArea
}

func maxArea1(height []int) int {
	maxArea := 0
	l := len(height)
    for i := 0; i < l - 1; i++ {
		for j := i + 1; j < l; j++ {
			area := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}