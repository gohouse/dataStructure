package avl

import (
	"fmt"
	"github.com/gohouse/dataStructure/queue"
)

type AvlNode struct {
	data   int
	left   *AvlNode
	right  *AvlNode
}

type AVLTree struct {
	root *AvlNode
	size int
}

func NewAVL() *AVLTree {
	return new(AVLTree)
}

func (avl *AVLTree) Add(data int) *AVLTree {
	avl.root = avl._add(avl.root, data)
	avl.size++
	return avl
}

func (avl *AVLTree) _add(node *AvlNode, data int) *AvlNode {
	if node == nil {
		node = &AvlNode{data: data}
	} else if data < node.data {
		node.left = avl._add(node.left, data)
		if node.left.depth()-node.right.depth() >= 2 {
			if data < node.left.data {
				node = avl.LL_Rotate(node)
			} else {
				node = avl.LR_Rotate(node)
			}
		}
	} else {
		node.right = avl._add(node.right, data)
		if node.right.depth()-node.left.depth() >= 2 {
			if data > node.right.data {
				node = avl.RR_Rotate(node)
			} else {
				node = avl.RL_Rotate(node)
			}
		}
	}

	return node
}

func (avl *AVLTree) Width() int {
	if avl.size == 0 {
		return 0
	}

	var maxWidth = 0
	var q queue.Queue
	q.EnQueue(q.Data(avl.root))

	for {
		size := q.Size()
		if size == 0 {
			break
		}
		for size > 0 {
			node := q.DeQueue()
			size--

			avlNode := node.Data.(*AvlNode)
			if avlNode.left != nil {
				q.EnQueue(q.Data(avlNode.left))
			}
			if avlNode.right != nil {
				q.EnQueue(q.Data(avlNode.right))
			}
		}
		maxWidth = max(maxWidth, q.Size())
	}

	return maxWidth
}

func (avl *AVLTree) Depth() int {
	return avl.root.depth()
}

func (an *AvlNode) depth() int {
	if an == nil {
		return 0
	}
	return 1 + max(an.left.depth(), an.right.depth())
}

func (avl *AVLTree) Size() int {
	return avl.size
}

func (avl *AVLTree) LL_Rotate(t *AvlNode) *AvlNode {
	headNode := t.left
	t.left = headNode.right
	headNode.right = t
	return headNode
}

func (avl *AVLTree) RR_Rotate(t *AvlNode) *AvlNode {
	headNode := t.right
	t.right = headNode.left
	headNode.left = t
	return headNode
}

func (avl *AVLTree) LR_Rotate(t *AvlNode) *AvlNode {
	t.left = avl.RR_Rotate(t.left)
	return avl.LL_Rotate(t)
}

func (avl *AVLTree) RL_Rotate(t *AvlNode) *AvlNode {
	t.right = avl.LL_Rotate(t.right)
	return avl.RR_Rotate(t)
}

func (avl *AVLTree) Show() {
	if avl.Size() == 0 {
		return
	}
	avl.root.show()
}

func (an *AvlNode) show() {
	var q = queue.NewQueue()
	q.EnQueue(q.Data(an))

	for {
		size := q.Size()
		if size == 0 {
			break
		}
		for size > 0 {
			cur := q.DeQueue()
			curNode := cur.Data.(*AvlNode)

			fmt.Printf("%v, ", curNode.data)
			size--

			if curNode.left != nil {
				q.EnQueue(q.Data(curNode.left))
			}
			if curNode.right != nil {
				q.EnQueue(q.Data(curNode.right))
			}
		}
		fmt.Println()
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

