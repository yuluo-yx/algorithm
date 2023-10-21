package sort

var sortTestsData = []struct {
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
		input:    []int{2, 1, 3, 6, 5, 4, 7},
		expected: []int{1, 2, 3, 4, 5, 6, 7},
		name:     "quick-sort-test",
	},
	{
		input:    []int{1},
		expected: []int{1},
		name:     "quick-sort-test",
	},
}
