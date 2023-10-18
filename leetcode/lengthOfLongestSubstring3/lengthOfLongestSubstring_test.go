package lengthOfLongestSubstring3

import (
	"reflect"
	"testing"
)

var testData = []struct {
	input    string
	expected int
	name     string
}{
	{
		input:    "bbbbb",
		expected: 1,
		name:     "lengthOfLongestSubstring3-1",
	},
	{
		input:    "abcabcbb",
		expected: 3,
		name:     "lengthOfLongestSubstring3-2",
	},
	{
		input:    "pwwkew",
		expected: 3,
		name:     "lengthOfLongestSubstring3-3",
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			actual := lengthOfLongestSubstring(data.input)
			expected := reflect.DeepEqual(actual, data.expected)

			if !expected {
				t.Errorf("test %s failed", data.name)
				t.Errorf("actual %v expected %v", actual, data.expected)
			}
		})
	}

}
