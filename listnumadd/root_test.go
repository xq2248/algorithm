package listnumadd

import "testing"

func sliceCmp(c, d []uint32) bool {
	if len(c) != len(d) {
		return false
	}
	for i := 0; i < len(c); i++ {
		if c[i] != d[i] {
			return false
		}
	}
	return true
}

func TestSuite(t *testing.T) {
	var tests = []struct {
		a, b   []uint32
		expect []uint32
	}{
		{[]uint32{9}, []uint32{4, 3, 0}, []uint32{4, 3, 9}},
		{[]uint32{1, 2, 0}, []uint32{5, 6, 7, 3, 0}, []uint32{5, 6, 8, 5, 0}},
		{[]uint32{6, 7, 9}, []uint32{4, 3, 0}, []uint32{1, 1, 0, 9}},
	}
	for _, v := range tests {
		c := CaculateAdd(v.a, v.b)
		if !sliceCmp(c, v.expect) {
			t.Error(v.a, "+", v.b, "=", c, "shoud be", v.expect)
		}
	}
}
