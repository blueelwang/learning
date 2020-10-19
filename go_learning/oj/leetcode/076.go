package leetcode

import (
	"fmt"
)

func MinWindow() {
    fmt.Println(minWindow("aa", "aa"))
}

func minWindow(s string, t string) string {
	result := ""
	if 0 == len(s) || 0 == len(t) {
		return result
	}
	
	charCnt := map[byte]int{}
	nCover := 0
	for _, c := range t {
		charCnt[byte(c)]++
	}

	l, r, sLen, tLen, rangeLen := 0, 0, len(s), len(charCnt), len(s) + 1
	if _, ok := charCnt[s[l]]; ok {
		charCnt[s[l]]--
		if 0 == charCnt[s[l]] {
			nCover++
			if nCover >= tLen && r - l + 1 < rangeLen {
				rangeLen = r - l + 1
				result = s[l:r+1]
			}
		}
	}
	for l < r || r < sLen - 1 {
		if nCover >= tLen && l < r {
			if _, ok := charCnt[s[l]]; ok {
				charCnt[s[l]]++
				if charCnt[s[l]] > 0 {
					nCover--
				}
			}
			l++
			if nCover >= tLen && r - l + 1 < rangeLen {
				rangeLen = r - l + 1
				result = s[l:r+1]
			}
		} else if nCover < tLen && r < sLen - 1 {
			r++
			if _, ok := charCnt[s[r]]; ok {
				charCnt[s[r]]--
				if 0 == charCnt[s[r]] {
					nCover++
				}
			}
			if nCover >= tLen && r - l + 1 < rangeLen {
				rangeLen = r - l + 1
				result = s[l:r+1]
			}
		} else {
			break;
		}
	}

    return result
}
