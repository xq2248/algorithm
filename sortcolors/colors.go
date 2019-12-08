package sortcolors

//https://leetcode-cn.com/problems/sort-colors/

//SortColors: 实现
func SortColors(nums []int) {
	left, right := -1, len(nums)
	hasSwitch := true
	for hasSwitch {
		hasSwitch = false
		for i := left + 1; i < right; i++ {
			switch nums[i] {
			case 0:
				nums[left+1], nums[i] = nums[i], nums[left+1]
				left++
				hasSwitch = true
				break
			case 2:
				nums[right-1], nums[i] = nums[i], nums[right-1]
				right--
				hasSwitch = true
				break
			}
		}
	}
}

func SortColors2(nums []int) {
	count := len(nums)
	left, right := 0, count-1

	for i := left; i <= right; {
		switch nums[i] {
		case 0:
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		case 1:
			i++
		case 2:
			nums[right], nums[i] = nums[i], nums[right]
			right--
		}
	}

}
