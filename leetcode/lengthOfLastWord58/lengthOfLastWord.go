package lengthOfLastWord58

import (
	"strings"
)

func lengthOfLastWord(s string) int {

	return len(strings.Fields(s)[len(strings.Fields(s))-1])

}
