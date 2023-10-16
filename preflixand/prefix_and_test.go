package prefixand

import (
	"reflect"
	"testing"
)

var prefixAndTestData = []struct {
	input    []int
	qu       [][]int
	n, m     int
	expected []int
	name     string
}{
	{
		// 数据长度
		n: 5,
		// 数据
		input: []int{2, 1, 3, 6, 4, 0},
		// 询问数组
		qu: [][]int{{1, 2}, {1, 3}, {2, 4}},
		// 询问个数
		m: 3,
		// 预期结果
		expected: []int{3, 6, 10},
		name:     "prefix and",
	},
}

func TestPrefixAnd(t *testing.T) {

	for _, test := range prefixAndTestData {
		t.Run(test.name, func(t *testing.T) {

			actual := prefixAnd(test.input, test.n, test.qu, test.m)
			expected := reflect.DeepEqual(actual, test.expected)

			if !expected {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}

}
