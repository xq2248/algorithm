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

package listnumadd

import "testing"

// sliceCmp 比较2个整数切片
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
