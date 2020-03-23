package leetcode

import (
	"fmt"
)


func LongestCommonPrefix() {
	param := []string{"","b"}
	fmt.Println(longestCommonPrefix(param))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	l := 0
	for i := 0; i < len(strs); i++ {
		if i == 0 || len(strs[i]) < l {
			l = len(strs[i])
		}
	}

	prefixLen := 0
	loop1:
	for i := 0; i < l; i++ {
		var c byte = 0
		for j := 0; j < len(strs); j++ {
			if j == 0 {
				c = strs[j][i]
			} else if c != strs[j][i] {
				break loop1;
			}
		}
		prefixLen++
	}

	return strs[0][0:prefixLen]
}