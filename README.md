# dataStructure
the data structure implementation in golang (数据结构的go语言实现, 队列:queue; 散列表:hashtable; 二叉平衡树:avl-tree......)


## dataStructure index  
* [linkedList](#linkedList)  
* [queue](#queue)  
* [hashTable](#hashTable)  
* [tree](#tree)  
    * [AVL tree](#AVL-tree)  
    * [binarySearchTree](#binarySearchTree)
* [stack](#stack)
* [binaryHeap](#binaryHeap)

## linkedList
```go
package linkedList

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

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
```

## queue
```go
package queue

type Node struct {
	Data interface{}
	Next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

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
```

## hashTable
```go
package hashTable

type Options struct {
	// hashtable容量，设置默认桶容量
	Capacity uint
	// 负载因子 0<= x <=1
	LoadFactor float64
	// 是否记录扩容log
	Debug bool
}

type Entry struct {
	key   interface{}
	value interface{}
	next  *Entry
}

type IHashTable interface {
	// 添加
	Put(key interface{}, value interface{}) error
	// 获取key对应的值
	Get(key interface{}) interface{}
	// 删除
	Remove(key interface{}) error
	// 当前hashtable的数据量
	Size() int
	// 是否为空
	IsEmpty() bool
	// 打印每一个hash槽内的内容
	Show()
}
```

## tree
### AVL tree
```go
package avl

type AvlNode struct {
	data   int
	left   *AvlNode
	right  *AvlNode
}

type AVLTree struct {
	root *AvlNode
	size int
}

type IAVLTree interface {
	// 添加节点
	Add(data int) *AVLTree
	// 树宽
	Width() int
	// 树的最大深度
	Depth()
	// 树的节点数
	Size() int
	// 横向按层打印树
	Show()
}
```

### binarySearchTree
```go
package binarySearchTree

//const btdemo = `二叉树示例
//	10
//8		21
//	11		30
//		28		50
//			29
//`
type IBinarySearchTree interface {
	// 插入
	Insert(i int) bool
	// 查找
	Search(i int) *BinarySearchTree
	// 最大
	MaxNode() *BinarySearchTree
	// 最小
	MinNode() *BinarySearchTree
	// 前序遍历
	PreOrder() []int
	// 中序遍历
	InOrder() []int
	// 后续遍历
	PostOrder() []int
	// 删除
	Delete(i int) bool
	// 最大深度
	Depth() int
	// 最小深度
	DepthMin() int
	// 宽度
}

type BinarySearchTree struct {
	Data        int
	Left, Right *BinarySearchTree
}
```

## stack
```go
package stack

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
```

## binaryHeap
```go
package binaryHeap

type IBinaryHeap interface {
	Add(data int)
	AddSlice(data []int)
	Length() int
	BuildMax(dataLen int)
	BuildMin(dataLen int)
	SortAsc()
	SortDesc()
	Show()
}

type BinaryHeap struct {
	data []int
}
```
