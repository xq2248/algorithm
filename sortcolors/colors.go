/*
 * Copyright xq2248
 * mail: xq2248@163.com
 * github: xq2248
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sortcolors

//https://leetcode-cn.com/problems/sort-colors/
//算法描述：有三种颜色的数组，将数组按照红、白、蓝排列
//最简单的解法，肯定是排序算法，但是排序算法时间复杂度很高
//利用数组中的元素只有3种值，可以实现更高效的排序算法

//SortColors: 我的实现
func SortColors(nums []int) {
	left, right := -1, len(nums) //缺省认为-1是红色，最后追加一个蓝色
	hasSwitch := true
	for hasSwitch { //如果没有数据交换，则说明已经排好
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

//SortColors: 官方推荐的实现
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
