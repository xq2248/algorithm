package main

import "fmt"

//MaxSubStr: 计算最长的不重复子串
//From leetcode:https://leetcode-cn.com/problems/longest-palindromic-substring/
func lengthOfLongestSubstring(aa string) int {
	count := len(aa)
	if count < 1 {
		return count
	}

	maxSubLen := 1
	start := 0
	end := 0
	subLen := 1
	for i := 1; i < count; i++ {
		//如果aa[i]不在 aa[start] aa[i-1]中
		//fmt.Println("judge", string(aa[i]), "from", aa[start:end+1])
		res, idx := repeated(aa, i, start, end)
		if res {
			//subLen就不再变了
			start = idx + 1
			end = i
			//fmt.Println("repeated: set start", start, "end", end)
		} else { //不重复，则end继续增加
			end = i
			//fmt.Println("not repeated: set start", start, "end", end)
		}
		subLen = end - start + 1
		if maxSubLen < subLen {
			maxSubLen = subLen
			//fmt.Println("new max str", aa[start:end+1])
		}
	}
	return maxSubLen
}

func repeated(aa string, next, start, end int) (res bool, idx int) {
	res = false
	idx = 0
	for i := end; i >= start; i-- {
		if aa[next] == aa[i] {
			res = true
			idx = i
			break
		}
	}
	return res, idx
}

func main() {
	//bb :=
	fmt.Println(lengthOfLongestSubstring("bbbbbbb"))
	//bb := lengthOfLongestSubstring("dvdf")
	//fmt.Println(lengthOfLongestSubstring("dvdf"))

}
