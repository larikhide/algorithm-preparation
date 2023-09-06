package quickSort

import (
	"testing"
)

type testpair struct {
	data     []int
	expected []int
}

var testCases = []testpair{
	{[]int{1, 2, 3}, []int{1, 2, 3}},
	{[]int{1}, []int{1}},
	{[]int{1, 1, 2, 1}, []int{1, 1, 1, 2}},
	{[]int{4, 2, 1, 6}, []int{1, 2, 4, 6}},
}

func TestQuickSort(t *testing.T) {
	for _, v := range testCases {
		QuickSort(v.data, 0, len(v.data)-1)

		for i, j := range v.expected {
			if j != v.expected[i] {
				t.Fail()
			}

		}
	}

}
