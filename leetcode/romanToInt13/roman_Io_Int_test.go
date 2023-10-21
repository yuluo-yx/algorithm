package romanToInt13

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
		input:    "III",
		expected: 3,
		name:     "Roman To int test-01",
	},
	{
		input:    "LVIII",
		expected: 58,
		name:     "Roman To int test-02",
	},
	{
		input:    "MCMXCIV",
		expected: 1994,
		name:     "Roman To int test-03",
	},
}

func TestRomanToInt(t *testing.T) {
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			actual := romanToInt(data.input)
			expected := reflect.DeepEqual(actual, data.expected)

			if !expected {
				t.Errorf("test %s failed", data.name)
				t.Errorf("actual %v expected %v", actual, data.expected)
			}
		})
	}
}
