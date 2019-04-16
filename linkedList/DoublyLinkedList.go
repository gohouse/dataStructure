package linkedList

import "fmt"

type DublyNode struct {
	data interface{}
	prev *DublyNode
	next *DublyNode
}

type DoublyLinkedList struct {
	head *DublyNode
	tail *DublyNode
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return new(DoublyLinkedList)
}

func (d *DoublyLinkedList) Length() int {
	return d.size
}

func (d *DoublyLinkedList) Add(data interface{}) *DoublyLinkedList {
	tmp := &DublyNode{data:data,prev:d.tail}
	if d.size==0 {
		d.head = tmp
		d.tail = tmp
	} else {
		d.tail.next = tmp
		d.tail = tmp
	}

	d.size++
	return d
}

func (d *DoublyLinkedList) AddToHead(data interface{}) *DoublyLinkedList {
	tmp := &DublyNode{data:data,next:d.head}

	d.head = tmp

	d.size++
	return d
}

func (d *DoublyLinkedList) AddToTail(data interface{}) *DoublyLinkedList {
	tmp := &DublyNode{data:data,prev:d.tail}
	if d.size==0 {
		d.head = tmp
		d.tail = tmp
	} else {
		d.tail.next = tmp
		d.tail = tmp
	}

	d.size++
	return d
}

func (d *DoublyLinkedList) AddAt(index int, data interface{}) *DoublyLinkedList {
	if index==0 {
		return d.AddToHead(data)
	} else if index==d.size {
		return d.AddToTail(data)
	}
	
	findNode := d.Get(index)
	if findNode==nil {
		return nil
	}
	
	prev := findNode.prev
	newNode := &DublyNode{data:data,prev:prev,next:findNode}
	prev.next = newNode
	findNode.prev = newNode

	d.size++
	return d
}

func (d *DoublyLinkedList) Get(index int) *DublyNode {
	if d.size==0 || index<0 || index>d.size-1 {
		return nil
	}
	
	if index==0 {
		return d.head
	} else if index==d.size-1{
		return d.tail
	}
	
	var current *DublyNode
	// 判断在前半区间, 还是后半区间
	if index<=(d.size-1)/2 {
		var step = 0
		current = d.head
		for step<index {
			current = current.next
			step++
		}
		return current
	} else {
		var step = d.size-1
		current = d.tail
		for step>index {
			current=current.prev
			step--
		}
		return current
	}
}

func (d *DoublyLinkedList) Delete(index int) bool {
	findNode := d.Get(index)
	if findNode==nil {
		return false
	}
	prev := findNode.prev
	prev.next = findNode.next
	findNode.next.prev = findNode.prev

	findNode = nil

	d.size--
	return true
}

func (d *DoublyLinkedList) Show() {
	fmt.Printf("show: ")
	current := d.head
	for current!=nil {
		fmt.Printf("%v ",current.data)
		current=current.next
	}
	fmt.Println("")
}


