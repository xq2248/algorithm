package water

import "testing"

func TestSuite(t *testing.T) {
	//arr := []uint32{ 0,1,0,2,1,0,1,3,2,1,2,1}
	arr := []uint32{0, 5, 0, 5}
	RainAll(arr)
}
