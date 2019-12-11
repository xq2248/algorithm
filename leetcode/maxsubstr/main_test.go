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

import (
	"fmt"
	"testing"
)

func TestMaxSubStr(t *testing.T) {

	var tests = []struct {
		str    string
		expect int
	}{
		{"bbbbbbb", 1},
		{"dvdf", 3},
		{"abcde", 5},
		{"a", 1},
		{"", 0},
	}

	for _, test := range tests {
		res := LengthOfLongestSubstring(test.str)
		if res != test.expect {
			fmt.Println("sublen of ", test.str, "expect ", test.expect, "res ", res)
			t.Fail()
		}
	}

}
