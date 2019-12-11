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

package maxsubstr

//From leetcode:https://leetcode-cn.com/problems/longest-palindromic-substring/

//MaxSubStr: 算法入口，计算最长的不重复子串
func LengthOfLongestSubstring(aa string) int {
	count := len(aa)
	if count <= 1 {
		return count
	}

	maxSubLen := 0
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
