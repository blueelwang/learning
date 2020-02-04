package oj

import (
	"fmt"
)

func MinWindow() {
    fmt.Println(minWindow("a", "a"))
}

var result string = ""
var source string = ""
var tLen int = 0
var nLetters int = 0
var lettersCnt map[rune]int = map[rune]int{}
var l int = 0
var r int = -1
var change int = 1


func minWindow(s string, t string) string {
	if 0 == len(s) || 0 == len(t) {
		return result
	}
	source = s
	
	for _, c := range t {
		if _, ok := lettersCnt[c]; ok {
			continue
		}
		tLen++
		lettersCnt[c] = 0
	}
	

	for move() {
		if change > 0 {
			incr(rune(s[r]), 1)
		} else {
			incr(rune(s[l-1]), -1)
		}
	}

    return result
}

func incr(c rune, step int) {
	if _, ok := lettersCnt[c]; ok {
		lettersCnt[c] += step
		if step > 0 && lettersCnt[c] == 1 {
			nLetters++
		} else if step < 0 && lettersCnt[c] == 0 {
			nLetters--
		}
	}
	if nLetters >= tLen && ("" == result || len(result) > r - l + 1){
		result = source[l:r+1]
	}
}

func move() bool {
	if r == len(source) - 1 && l >= r {
		return false
	}

	if nLetters < tLen && r < len(source) - 1 {
		r++
		change = 1
		return true
	} else {
		l++
		change = -1
		return true
	}
}