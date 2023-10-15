package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {

	for _, test := range sortTestsData {
		t.Run(test.name, func(t *testing.T) {
			actual := mergeSort(test.input, make([]int, N), 0, len(test.input)-1)
			expected := reflect.DeepEqual(actual, test.expected)

			if !expected {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}

}
