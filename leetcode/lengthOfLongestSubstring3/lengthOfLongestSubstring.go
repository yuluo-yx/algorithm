package lengthOfLongestSubstring3

func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	set := make(map[byte]bool)

	l, r, maxLength := 0, 0, 0

	for r < len(s) {
		if set[s[r]] {
			delete(set, s[l])
			l++
		} else {
			set[s[r]] = true
			maxLength = getMax(maxLength, r-l+1)
			r++
		}
	}

	return maxLength

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
