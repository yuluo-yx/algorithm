package maxProfit121

import (
	"reflect"
	"testing"
)

var TestData = []struct {
	input    []int
	expected int
	name     string
}{
	{
		input:    []int{7, 1, 5, 3, 6, 4},
		expected: 5,
		name:     "test-1",
	},
}

func TestLengthOhLastWord(t *testing.T) {
	for _, data := range TestData {
		t.Run(data.name, func(t *testing.T) {
			actual := maxProfit(data.input)
			expected := reflect.DeepEqual(data.expected, actual)

			if !expected {
				t.Errorf("test %s failed", data.name)
				t.Errorf("actual %v expected %v", actual, data.expected)
			}
		})
	}
}
