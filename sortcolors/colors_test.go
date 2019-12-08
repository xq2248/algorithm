package sortcolors

import "testing"
import "fmt"

func TestSort(t *testing.T) {
	var ss []int = []int{2, 0, 2, 1, 1, 0}
	fmt.Println(ss)
	SortColors(ss)
	fmt.Println(ss)

}
