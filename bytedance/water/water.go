package water

//作者：大熊
//mail: xq2248@163.com

import (
	"fmt"
)

//From https://blog.csdn.net/DevolperFront/article/details/99688144
//字节跳动公司算法面试题
//问题3：
//有一组不同高度的台阶，用一个数组标识，数组中每个数是台阶的高度
//如果开始下雨，那么台阶之间的水坑会积水多少呢
//实现：找到最高水位
//左右2测分别设置一个水位，逐渐提升，对于提升后，低于低水位的，补充到水位
func GetMax(stairs []uint32) (maxIdx int, maxValue uint32) {
	maxIdx = 0
	maxValue = uint32(0)
	for i, v := range stairs {
		if maxValue < v {
			maxValue = stairs[i]
			maxIdx = i
		}
	}
	return maxIdx, maxValue
}

func RainAll(stairs []uint32) {

	//找最大的台阶和位置
	maxIdx, _ := GetMax(stairs)
	delta := make([]uint32, len(stairs))

	var sum uint32 = 0
	var leftLevel uint32 = 0
	for i := 0; i < maxIdx; i++ {
		if stairs[i] > leftLevel {
			leftLevel = stairs[i]
		} else {
			delta[i] = leftLevel - stairs[i]
			sum += delta[i]
		}
	}

	var rightLevel uint32 = 0
	for i := len(stairs) - 1; i > maxIdx; i-- {
		if stairs[i] > rightLevel {
			rightLevel = stairs[i]
		} else {
			delta[i] = rightLevel - stairs[i]
			sum += delta[i]
		}
	}

	fmt.Println(delta)
	fmt.Println(sum)
}
