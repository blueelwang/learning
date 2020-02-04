package oj

import (
	"fmt"
	"strings"
)


func IntToRoman() {
	n := 1994
	res := intToRoman(n)
	fmt.Println(res)
	fmt.Println(romanToInt(res))
}

func intToRoman(num int) string {
    table := [][]string {
		{"I", "V"},
		{"X", "L"},
		{"C", "D"},
		{"M"},
	}

	if num < 1 || num > 3999 {
		return ""
	}
	res := ""
	for i := 0; num > 0; i++ {
		n := num % 10
		num = num / 10
		switch {
		case n == 0:
			continue
		case n <= 3:
			res = strings.Repeat(table[i][0], n) + res
		case n == 4:
			res = table[i][0] + table[i][1] + res
		case n < 9:
			res = table[i][1] + strings.Repeat(table[i][0], n - 5) + res
		default:
			res = table[i][0] + table[i+1][0] + res
		}
	}
	return res
}

func romanToInt(s string) int {
	l := len(s)
    if (l <= 0) {
		return 0
	}
	table := map[rune]int {
		'I' : 1, 
		'V' : 5,
		'X' : 10,
		'L' : 50,
		'C' : 100,
		'D' : 500,
		'M' : 1000,
	}

	res := 0
	last := 0
	for i := 0; i < l; i++ {
		num := table[rune(s[i])]
		if (i == 0) {
			res += num;
		} else if last >= num {
			res += num
		} else {
			res += num - 2 * last
		}
		last = num;
	}
	return res
}