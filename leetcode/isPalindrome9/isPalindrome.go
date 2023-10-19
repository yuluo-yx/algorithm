package isPalindrome9

import (
	"strconv"
)

// isPalindrome 思路为：将前一半字符串反转和后一半进行比较，相同则为回文串。
func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	prefix := str[:len(str)/2]
	var suffix string
	if !isEven(len(str)) {
		suffix = str[len(str)/2+1:]
	} else {
		suffix = str[len(str)/2:]
	}

	if suffix == Reverse(prefix) {
		return true
	}

	return false
}

// isEven 判断是否偶数
func isEven(num int) bool {

	return num%2 == 0
}

// Reverse 反转字符串
func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
