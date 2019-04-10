# dataStructure
the data structure implementation in golang

## dataStructure index  
* [queue](#queue)  
* [hashTable](#hashTable)  
* [tree](#tree)  
    * [AVL tree](#AVL-tree)  

## queue
```go
package queue

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
