package leetcode

import (
	"math"
	"fmt"
)

func IsPalindrome() {
	x := 423424
	fmt.Println(isPalindrome(x))
}


func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	len := 0
	n := x
	for n > 0 {
		len++
		n = n / 10
	}

	for i := 1; i <= int(math.Ceil(float64(len) / 2)); i++ {
		n1 := x % int(math.Pow(10, float64(i))) / int(math.Pow(10, float64(i - 1)))
		n2 := x / int(math.Pow(10, float64(len - i))) % 10
		if n2 != n1 {
			return false
		}
	}
	return true
}