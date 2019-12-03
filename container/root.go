/*
 *   Copyright (c) 2019
 *   All rights reserved.
 */

package container

//From https://leetcode-cn.com/problems/container-with-most-water/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode/
//盛最多水的容器

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	count := len(height)
	var maxArea int = 0
	for i := 0; i < count; i++ {
		for j := 0; j < i; j++ {
			res := min(height[i], height[j]) * (i - j)
			if maxArea < res {
				maxArea = res
			}
		}
	}
	return maxArea
}
