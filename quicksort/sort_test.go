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
	}

	for _, test := range tests {
		res := QuickSort(test.nums)
		isEqual := utils.SliceCmp(res, test.expect)
		if !isEqual {
			fmt.Println("sort res of ", test.nums)
			fmt.Println("expect ", test.expect)
			fmt.Println("res ", res)
			t.Fail()
		}
	}
}
