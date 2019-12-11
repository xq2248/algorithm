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

import (
	"fmt"
	"sort"
)

import "testing"

//golang官方提供的sort使用示例
//示例是定义一个数组，数组的元素是平面中的像素点
//按照离原点(0,0)的距离进行排序

//自定义数组切片
type Points []struct {
	x, y int
}

//需要实现sort.Interface接口的三个方法
//因为切片自带本身就是指针头,因此这里并没有使用指针
func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	sx := p[i].x*p[i].x + p[i].y*p[i].y
	sy := p[j].x*p[j].x + p[j].y*p[j].y
	return sx < sy
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func TestService(t *testing.T) {
	var ps Points = []struct {
		x, y int
	}{
		{1, 2},
		{1, 3},
		{2, 2},
	}

	fmt.Println(ps)
	sort.Sort(ps)
	fmt.Println(ps)
}
