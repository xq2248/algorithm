package quicksort

import (
	"fmt"
	"testing"
)
import "algorithm/utils"

func TestQuickSort(t *testing.T) {

	var tests = []struct {
		nums   []int
		expect []int
	}{
		{[]int{3, 1, 5, 6, 2}, []int{1, 2, 3, 5, 6}},
		{[]int{33, 44}, []int{33, 44}},
		{[]int{33}, []int{33}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		QuickSort(test.nums)
		isEqual := utils.SliceCmp(test.nums, test.expect)
		if !isEqual {
			fmt.Println("sort res of ", test.nums)
			fmt.Println("expect ", test.expect)
			t.Fail()
		}
	}
}

func TestQuickSort2(t *testing.T) {

	var tests = []struct {
		nums   []int
		expect []int
	}{
		{[]int{3, 1, 5, 6, 2}, []int{1, 2, 3, 5, 6}},
		{[]int{33, 44}, []int{33, 44}},
		{[]int{33}, []int{33}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		QuickSort2(test.nums)
		isEqual := utils.SliceCmp(test.nums, test.expect)
		if !isEqual {
			fmt.Println("sort res of ", test.nums)
			fmt.Println("expect ", test.expect)
			t.Fail()
		}
	}
}
