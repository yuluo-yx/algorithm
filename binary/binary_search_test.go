package binary

import (
	"reflect"
	"testing"
)

var binarySearchTestData = []struct {
	n, m     int
	input    []int
	qu       []int
	expected [][]int
	name     string
}{
	{
		// 输入数据长度
		n: 6,
		// 查找值个数
		m: 3,
		// 输入数据
		input: []int{1, 2, 2, 3, 3, 4},
		// 查找数据
		qu: []int{3, 4, 5},
		// 期望值
		expected: [][]int{{3, 4}, {5, 5}, {-1, -1}},
		name:     "binary_search_test",
	},
}

func TestBinarySearch(t *testing.T) {

	for _, test := range binarySearchTestData {
		t.Run(test.name, func(t *testing.T) {

			for i := 0; i < test.m; i++ {
				actual := binarySearch(test.input, test.n, test.qu[i])
				expected := reflect.DeepEqual(actual, test.expected[i])

				if !expected {
					t.Errorf("test %s failed", test.name)
					t.Errorf("actual %v expected %v", actual, test.expected[i])
				}
			}
		})
	}

}
