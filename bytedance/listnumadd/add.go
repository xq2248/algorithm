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

//From https://blog.csdn.net/DevolperFront/article/details/99688144
//字节跳动公司算法面试题
//问题2：
//用单向列表保存2个十进制整数
//然后实现相加,注意只能用单向链表

//实现思路：
//将2个链表倒序，相加后，再倒序

type elem struct {
	value uint32
	next  *elem
}

func newElem(value uint32) *elem {
	return &elem{value, nil}
}

type list struct {
	number uint32
	head   *elem
}

func newList() *list {
	return &list{0, nil}
}

func convertSliceToList(a []uint32) *list {
	list := newList()
	for _, v := range a {
		//ne := newElem( v )
		list.addTail(v)
	}
	return list
}

func (list *list) addTail(value uint32) {
	ne := newElem(value)
	if list.head == nil {
		list.head = ne
		list.number = 1
		return
	}
	cur := list.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = ne
	list.number++
}

func (list *list) addHead(value uint32) {
	ne := newElem(value)
	ne.next = list.head
	list.head = ne
	list.number++
}

//traverse 链表反转
func (list *list) traverse() *list {
	newList := newList()
	cur := list.head
	for ; cur != nil; cur = cur.next {
		newList.addHead(cur.value)
	}
	return newList
}

//listAdd 两个链表相加
func listAdd(a, b *list) *list {
	newList := newList()
	count := a.number
	if count < b.number {
		count = b.number
	}
	delta := uint32(0)
	curA := a.head
	curB := b.head
	for i := uint32(0); i < count; i++ {
		sum := delta //计算进位
		delta = 0    //进位清0
		if curA != nil {
			sum += curA.value
			curA = curA.next
		}
		if curB != nil {
			sum += curB.value
			curB = curB.next
		}
		mod := sum % 10
		delta = sum / 10
		newList.addTail(mod)
	}
	if delta > 0 {
		newList.addTail(delta)
	}
	return newList
}

// convertListToSlice convert list to slice
func (list *list) convertListToSlice() []uint32 {
	resLen := list.number + 1
	resSlice := make([]uint32, 0, resLen)
	cur := list.head
	for ; cur != nil; cur = cur.next {
		resSlice = append(resSlice, cur.value)
	}
	return resSlice
}

//CaculateAdd return the addtion result of two integer slices
func CaculateAdd(sa, sb []uint32) []uint32 {
	la1 := convertSliceToList(sa)
	lb1 := convertSliceToList(sb)
	la2 := la1.traverse()
	lb2 := lb1.traverse()
	lc2 := listAdd(la2, lb2)
	lc1 := lc2.traverse()
	return lc1.convertListToSlice()
}
