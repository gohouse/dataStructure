# dataStructure
the data structure implementation in golang

## dataStructure index  
- [queue](#queue)  
- [hashTable](#hashTable)  
- [tree](#tree)  
    - [AVL tree](#AVL-tree)  

## queue
```go
package main
import (
	"fmt"
	"github.com/gohouse/dataStructure/queue"
)
func main() {
	var q = queue.NewQueue()
	
	q.EnQueue(q.Data(1))
	q.EnQueue(q.Data(4))
	q.EnQueue(&queue.Node{Data:8})
	q.EnQueue(q.Data(2))
	
	q.DeQueue()

	q.Show()
}
```

## hashTable

## tree
### AVL tree
