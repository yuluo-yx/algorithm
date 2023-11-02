package lengthOfLastWord58

import (
	"reflect"
	"testing"
)

var TestData = []struct {
	input    string
	expected int
	name     string
}{
	{
		input:    "Hello World",
		expected: 5,
		name:     "test-1",
	},
	{
		input:    "   fly me   to   the moon ",
		expected: 4,
		name:     "test-2",
	},
	{
		input:    "luffy is ",
		expected: 2,
		name:     "test-3",
	},
}

func TestLengthOhLastWord(t *testing.T) {
	for _, data := range TestData {
		t.Run(data.name, func(t *testing.T) {
			actual := lengthOfLastWord(data.input)
			expected := reflect.DeepEqual(data.expected, actual)

			if !expected {
				t.Errorf("test %s failed", data.name)
				t.Errorf("actual %v expected %v", actual, data.expected)
			}
		})
	}
}
