package oj

import (
	"fmt"
)



func Convert() {
	fmt.Println(convert2("LEETCODEISHIRING", 4))
	fmt.Println(convert("LEETCODEISHIRING", 4))
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	numChars := len(s)
	if numChars <= numRows {
		return s
	}

	var result string = ""
	mod := numRows * 2 - 2	
	for row := 1; row <= numRows; row++ {
		for i, c := range s {
			n := i % mod
			if (n + 1 == row) || (mod - n == row - 1) {
				result += string(c)
			}
		}
	}

	return result
}

func convert2(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	numChars := len(s)
	if numChars <= numRows {
		return s
	}

	var result string = ""
	mod := numRows * 2 - 2

	rows := make([][]byte, numRows, numRows)
	row := 0
	direction := -1
	for i, c := range s {
		if rows[row] == nil {
			rows[row] = []byte{}
		}
		rows[row] = append(rows[row], byte(c))

		n := i % mod
		if n == numRows - 1 || n == 0 {
			direction *= -1
		}
		row += direction
	}
	for _, row := range rows {
		result += string(row)
	}
	return result
}