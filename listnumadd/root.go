package listnumadd

//作者：大熊
//mail: xq2248@163.com

//From https://blog.csdn.net/DevolperFront/article/details/99688144
//问题2：
//用单向列表保存2个十进制整数
//然后实现相加,注意只能用单向链表

//实现思路：
//将2个链表倒序，相加后，再倒序

type Elem struct {
	value uint32
	next  *Elem
}

func NewElem(value uint32) *Elem {
	return &Elem{value, nil}
}

type List struct {
	number uint32
	head   *Elem
}

func NewList() *List {
	return &List{0, nil}
}

func ConvertSliceToList(a []uint32) *List {
	list := NewList()
	for _, v := range a {
		//ne := NewElem( v )
		list.addTail(v)
	}
	return list
}

func (list *List) addTail(value uint32) {
	ne := NewElem(value)
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

func (list *List) addHead(value uint32) {
	ne := NewElem(value)
	ne.next = list.head
	list.head = ne
	list.number++
}

func (list *List) Traverse() *List {
	newList := NewList()
	cur := list.head
	for ; cur != nil; cur = cur.next {
		newList.addHead(cur.value)
	}
	return newList
}

func ListAdd(a, b *List) *List {
	newList := NewList()
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

func (list *List) ConvertListToSlice() []uint32 {
	resLen := list.number + 1
	resSlice := make([]uint32, 0, resLen)
	cur := list.head
	for ; cur != nil; cur = cur.next {
		resSlice = append(resSlice, cur.value)
	}
	return resSlice
}

//算法入口，对2个整数切片进行按位相加
func CaculateAdd(sa, sb []uint32) []uint32 {
	la1 := ConvertSliceToList(sa)
	lb1 := ConvertSliceToList(sb)
	la2 := la1.Traverse()
	lb2 := lb1.Traverse()
	lc2 := ListAdd(la2, lb2)
	lc1 := lc2.Traverse()
	return lc1.ConvertListToSlice()
}
