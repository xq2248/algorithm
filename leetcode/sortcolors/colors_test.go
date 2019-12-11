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

import "testing"
import "algorithm/utils"

func TestSort(t *testing.T) {

	var tcs = []struct {
		nums   []int
		expect []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, tc := range tcs {
		SortColors(tc.nums)
		if !utils.SliceCmp(tc.nums, tc.expect) {
			t.Fail()
		}
	}
}

func TestSort2(t *testing.T) {

	var tcs = []struct {
		nums   []int
		expect []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, tc := range tcs {
		SortColors2(tc.nums)
		if !utils.SliceCmp(tc.nums, tc.expect) {
			t.Fail()
		}
	}
}
