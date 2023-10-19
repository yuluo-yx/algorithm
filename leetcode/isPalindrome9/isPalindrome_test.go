package isPalindrome9

import (
	"reflect"
	"testing"
)

var testData = []struct {
	input    int
	expected bool
	name     string
}{
	{
		input:    0,
		expected: true,
		name:     "isPalindrome-test-00",
	},
	{
		input:    12221,
		expected: true,
		name:     "isPalindrome-test-01",
	},
	{
		input:    121,
		expected: true,
		name:     "isPalindrome-test-02",
	},
	{
		input:    -121,
		expected: false,
		name:     "isPalindrome-test-03",
	},
	{
		input:    10,
		expected: false,
		name:     "isPalindrome-test-04",
	},
	{
		input:    5,
		expected: true,
		name:     "isPalindrome-test-04",
	},
}

func TestIsPalindrome(t *testing.T) {

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			actual := isPalindrome(data.input)
			expected := reflect.DeepEqual(actual, data.expected)

			if !expected {
				t.Errorf("test %s failed", data.name)
				t.Errorf("actual %v expected %v", actual, data.expected)
			}
		})
	}

}
