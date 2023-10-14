package sort

import (
	"reflect"
	"testing"
)

var sortTestsData = []struct {
	total    int64
	input    []int
	expected []int
	name     string
}{
	{
		input:    []int{2, 1, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
		name:     "quick-sort-test",
	},
	{
		input:    []int{1},
		expected: []int{1},
		name:     "quick-sort-test",
	},
}

func TestQuickSort(t *testing.T) {

	for _, test := range sortTestsData {
		t.Run(test.name, func(t *testing.T) {
			actual := quickSort(test.input, 0, len(test.input)-1)
			expected := reflect.DeepEqual(actual, test.expected)

			if !expected {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}

}
