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

package quicksort

//partition 算法导论中标准实现，好处是简单，坏处是，如果有很多相等的元素，数据交换比较多
func partition(data []int, left, right int) int {
	pivot := data[right] //提取右侧的元素当做标尺
	i := left
	for j := left; j <= right-1; j++ { //从左到右遍历元素
		if data[j] <= pivot { //如果元素是小于标尺的
			data[i], data[j] = data[j], data[i] //将该数据与第一个
			i++                                 //到此保证了data[i]左边的元素都是小于标尺的
		}
	}
	data[i], data[right] = data[right], data[i] //将标尺与第一个大于标尺的元素互换
	return i
}

//quickSort 算法递归
func quickSort(data []int, left, right int) {
	if left >= right {
		return
	}
	idx := partition(data, left, right)
	quickSort(data, left, idx-1)
	quickSort(data, idx+1, right)
}

//QuickSort 快速排序的实现
//时间复杂度最小O(n*log(n))，最大O(n*n),空间复杂度为%(n)
func QuickSort(data []int) {
	count := len(data)
	quickSort(data, 0, count-1)
}

//partition2 参考了网上多层嵌套循环方法，避免元素交换过多
func partition2(data []int, left, right int) int {
	pivot := data[right] //提取右侧的元素当做标尺

	for left < right {
		//从左边开始寻找第一个大于标尺的数据
		for left < right && data[left] <= pivot {
			left++
		}
		//将这个data[left]放到右边，原来标尺的位置,从而保证接下来的循环
		data[right] = data[left]

		//由于此时data[right]肯定是大于标志,那么循环肯定执行
		for left < right && data[right] >= pivot {
			right--
		}
		//将小于标尺的data[right]放在左边，保证下一次循环
		data[left] = data[right]
	}
	//退出循环的时候left肯定是等于right，
	data[right] = pivot
	return right
}

func quickSort2(data []int, left, right int) {
	if left >= right {
		return
	}
	idx := partition2(data, left, right)
	quickSort2(data, left, idx-1)
	quickSort2(data, idx+1, right)
}

//QuickSort 快速排序实现
func QuickSort2(data []int) {
	count := len(data)
	quickSort2(data, 0, count-1)
}
