package leetcode

import (
	"fmt"
)

func IsInterleave() {
	fmt.Println(isInterleave("bab", "abbbcba", "babbbabcba"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}


func isInterleave(s1 string, s2 string, s3 string) bool {
	len1, len2, len3 := len(s1), len(s2), len(s3)
	if len1 + len2 != len3 {
		return false
	}

	for index1, index2, index3 := 0, 0, 0; index3 < len3; index3++ {
		if index1 < len1 && s3[index3] == s1[index1] {
			index1++
		} else if index2 < len2 && s3[index3] == s2[index2] {
			index2++
		} else {
			if index1 - 1 >= 0 && index2 < len2 && s1[index1 - 1] == s2[index2] {
				index1--
				index2++
				index3--
			} else {
				return false;
			}
		}
	}
	return true;
}