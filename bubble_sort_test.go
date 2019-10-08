package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{2, 1, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2}, []int{2}},
		{[]int{1, 3, 4, 2, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{4, 5, 2, 1, 3,}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 5, 3, 4,}, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		bubbleSort(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("slice %v should be sorted as %v", test.input, test.want)
		}
	}

}
