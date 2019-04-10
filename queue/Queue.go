package queue

import "fmt"

type IQueue interface {
	// 入队列
	EnQueue(data *Node) *Queue
	// 出队列
	DeQueue() *Node
	// 构造入队列数据节点
	Data(data interface{}) *Node
	// 获取队列第一个元素,不出队列
	Peek() *Node
	// 队列是否为空
	IsEmpty() bool
	// 队列数据数量
	Size() int
	// 打印队列
	Show()
}

type Node struct {
	Data interface{}
	Next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func NewQueue() *Queue {
	return &Queue{}
}
// EnQueue : 入列
func (q *Queue) Data(data interface{}) *Node {
	return &Node{Data:data}
}
// EnQueue : 入列
func (q *Queue) EnQueue(data *Node) *Queue {
	if q.size==0 {
		q.head = data
	} else {
		q.tail.Next = data
	}
	q.tail = data
	q.size++
	return q
}
// DeQueue : 出列
func (q *Queue) DeQueue() *Node {
	tmp := q.head
	if tmp==nil {
		return tmp
	}
	q.head = tmp.Next
	q.size--
	return tmp
}
// Peek : 获取队列首部第一个数据
func (q *Queue) Peek() *Node {
	return q.head
}
// IsEmpty
func (q *Queue) IsEmpty() bool {
	return q.size==0
}
// Show 打印队列
func (q *Queue) Show() {
	current := q.head
	fmt.Printf("%s: ","show")
	for current!=nil {
		fmt.Printf("%d ", current.Data)
		current=current.Next
	}
	fmt.Println("")
}
// Length 队列长度
func (q *Queue) Size() int {
	return q.size
}
