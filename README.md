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
	EnQueue(data *Node) *Queue
	DeQueue() *Node
	Peek() *Node
	IsEmpty() bool
	Size() int
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
	Add(data int) *AVLTree
	Width() int
	Depth()
	Size() int
	Show()
}
```
