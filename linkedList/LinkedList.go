package linkedList

type ILinkedList interface {
	AddToFirst(data interface{})
	AddToLast(data interface{})
	Show() (res []interface{})
	Get(index int) *Node
	GetPrev(index int) *Node
	Add(index int, data interface{}) bool
	Delete(index int) bool
	Reverse() *LinkedList
	ValueOf(data interface{}) (index int)
}

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) AddToFirst(data interface{}) {
	tmp := &Node{data: data, next: list.head}
	list.head = tmp
	list.size++
}

func (list *LinkedList) AddToLast(data interface{}) {
	current := list.head
	for current.next != nil {
		current = current.next
	}
	if current.next == nil {
		current.next = &Node{data: data}
	}
	list.size++
}

func (list *LinkedList) Show() (res []interface{}) {
	var current = list.head
	for current != nil {
		res = append(res, current.data)
		current = current.next
	}
	return
}

func (list *LinkedList) Get(index int) *Node {
	if list.size <= index+1 || list.size == 0 {
		return nil
	}
	var i = 0
	current := list.head
	for i != index {
		current = current.next
		i++
	}
	return current
}

func (list *LinkedList) GetPrev(index int) *Node {
	if list.size < index || list.size <=1 || index==0 {
		return nil
	}
	var i = 0
	current := list.head
	for i != index-1 {
		current = current.next
		i++
	}
	return current
}

func (list *LinkedList) Add(index int, data interface{}) bool {
	if index == 0 {
		list.AddToFirst(data)
	} else if index == list.size {
		list.AddToLast(data)
	} else if index>list.size {
		return false
	}

	prevNode := list.GetPrev(index)

	if prevNode==nil {
		return false
	}
	prevNode.next = &Node{data,prevNode.next}

	list.size++

	return true
}

func (list *LinkedList) Delete(index int) bool {
	if list.size==1 {
		list.head = nil
		return true
	}
	prevNode := list.GetPrev(index)

	if prevNode==nil {
		return false
	}

	prevNode.next = prevNode.next.next

	list.size--

	return true
}
// Reverse 反转链表
// 第一种方案: 原理, 从第二个开始, 依次放入到前边即可, 空间复杂度为O(1)
// 第二种方案: 原理, 循环构建新的链表, 空间复杂度为O(n)
func (list *LinkedList) Reverse() *LinkedList {
	if list==nil {
		return list
	}
	var tmp = &Node{-1, list.head}

	var prev = tmp.next
	var current = prev.next

	for current!=nil {
		prev.next = current.next
		current.next = tmp.next
		tmp.next = current
		current = prev.next
	}

	return &LinkedList{tmp.next,list.size}
}

func (list *LinkedList) ValueOf(data interface{}) (index int) {
	if list.size==0 {
		return -1
	}
	i:=0
	current := list.head

	for current.data!=data {
		current=current.next
		i++
	}

	return i
}
