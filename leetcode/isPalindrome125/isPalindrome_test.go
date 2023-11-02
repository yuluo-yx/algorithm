package isPalindrome125

import (
	"reflect"
	"testing"
)

var TestData = []struct {
	input    string
	expected bool
	name     string
}{
	{
		input:    "A man, a plan, a canal: Panama",
		expected: true,
		name:     "test - 1",
	},
	{
		input:    "0q",
		expected: false,
		name:     "test - 2",
	},
	{
		input:    " ",
		expected: true,
		name:     "test - 3",
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, data := range TestData {
		t.Run(data.name, func(t *testing.T) {
			actual := isPalindrome(data.input)
			expected := reflect.DeepEqual(data.expected, actual)

			if !expected {
				t.Errorf("test %s failed", data.name)
				t.Errorf("actual %v expected %v", actual, data.expected)
			}
		})
	}
}
