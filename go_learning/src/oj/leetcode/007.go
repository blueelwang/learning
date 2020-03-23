package leetcode

import (
	"fmt"
)


func Reverse(x int) int {
    result := 0

    for x != 0 {
		tmp := x % 10
		result = result * 10 + tmp

		if result > 1 << 31 - 1 || result < -1 << 31 {
			return 0
		}
		x = x / 10
	}
	fmt.Println(result)

	return result
}