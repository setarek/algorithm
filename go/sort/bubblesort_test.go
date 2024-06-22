package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int64
		expected []int64
	}{
		{[]int64{5, 2, 9, 1, 5, 6}, []int64{1, 2, 5, 5, 6, 9}},
		{[]int64{3, 0, 2, 5, -1, 4, 1}, []int64{-1, 0, 1, 2, 3, 4, 5}},
		{[]int64{2, 1}, []int64{1, 2}},
		{[]int64{}, []int64{}},
		{[]int64{1}, []int64{1}},
	}

	for _, test := range tests {
		result := BubbleSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
