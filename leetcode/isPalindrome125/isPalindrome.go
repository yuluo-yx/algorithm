package isPalindrome125

import (
	"regexp"
	"strings"
)

// isPalindrome 双指针解法
func isPalindrome(s string) bool {

	var flag = true

	s = convert(s)

	if len(s) == 1 {
		return true
	}

	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			flag = false
		}
	}

	return flag

}

func convert(s string) string {

	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	reg := regexp.MustCompile(`[^a-z0-9]`)

	return reg.ReplaceAllString(s, "")

}
