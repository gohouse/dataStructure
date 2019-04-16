package stack

import "fmt"

type IStack interface {
	Push(data interface{}) bool
	Pop() interface{}
	Length() int
	Show()
}

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func New() *Stack {
	return new(Stack)
}

func (s *Stack) Push(data interface{}) bool {
	tmp := &Node{data:data}
	if s.size==0 {
		s.head = tmp
	} else {
		tmp.next = s.head
		s.head = tmp
	}
	s.size++
	tmp = nil
	return true
}

func (s *Stack) Pop() interface{} {
	if s.size==0 {
		return nil
	} else {
		head := s.head
		s.head = head.next
		s.size--
		head.next = nil
		return head.data
	}
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) Show() {
	current := s.head
	fmt.Printf("%s: ","show")
	for current!=nil {
		fmt.Printf("%d ", current.data)
		current=current.next
	}
	fmt.Println("")
}
