package TwoSum1

import (
	"reflect"
	"testing"
)

var testData = []struct {
	input    []int
	target   int
	expected []int
	name     string
}{
	{
		input:    []int{2, 5, 5, 11},
		target:   10,
		expected: []int{1, 2},
		name:     "test-two-num-01",
	},
	{
		input:    []int{3, 3},
		target:   6,
		expected: []int{0, 1},
		name:     "test-two-num-02",
	},
	{
		input:    []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
		name:     "test-two-num-03",
	},
	{
		input:    []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
		name:     "test-two-num-04",
	},
}

func TestTwoNumPlus(t *testing.T) {

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			actual := twoNumPlus(data.input, data.target)
			expected := reflect.DeepEqual(actual, data.expected)

			if !expected {
				t.Errorf("test %s failed", data.name)
				t.Errorf("actual %v expected %v", actual, data.expected)
			}
		})
	}

}
