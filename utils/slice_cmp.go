package utils

// sliceCmp 比较2个整数切片
func SliceCmp(c, d []int) bool {
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
