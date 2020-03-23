package leetcode


import (
	"fmt"
)


func LetterCombinations() {
	res := letterCombinations("23")
	fmt.Println(res)
}

func letterCombinations(digits string) []string {
    table := map[rune][]rune{
		'2' : {'a', 'b', 'c'},
		'3' : {'d', 'e', 'f'},
		'4' : {'g', 'h', 'i'},
		'5' : {'j', 'k', 'l'},
		'6' : {'m', 'n', 'o'},
		'7' : {'p', 'q', 'r', 's'},
		'8' : {'t', 'u', 'v'},
		'9' : {'w', 'x', 'y', 'z'},
	}
	l := len(digits)
	if l == 0 {
		return []string{}
	}

	result := []string{}

	pos := make([]int, l)
	for i, c := range digits {
		pos[i] = len(table[c]) - 1
	}
	
	for {
		s := make([]rune, l)
		for i, p := range pos {
			c := rune(digits[i])
			s[i] = table[c][p]
		}
		result = append(result, string(s))

		if zero(pos) {
			break
		}
		pos = decr(pos, digits, table)
	}

	return result
}

func decr(arr []int, digits string, table map[rune][] rune) []int {
	l := len(arr)
	for i := l - 1; i >= 0; i-- {
		if arr[i] > 0 {
			arr[i]--
			return arr
		} else {
			arr[i] = len(table[rune(digits[i])]) - 1
		}
	}
	return arr
}

func zero(arr []int) bool {
	for _, n := range arr {
		if n != 0 {
			return false
		}
	}
	return true
}