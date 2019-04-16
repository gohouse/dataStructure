package binarySearchTree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := New()
	fmt.Println(bst.Data == 0)
	bst.Insert(10)
	bst.Insert(8)
	bst.Insert(21)
	bst.Insert(11)
	bst.Insert(30)
	bst.Insert(28)
	bst.Insert(29)
	bst.Insert(50)

	fmt.Printf("最大:%v,最小:%v", bst.MaxNode().Data, bst.MinNode().Data)
	fmt.Println("\n前序遍历:")
	fmt.Println(bst.PreOrder())
	fmt.Println("中序遍历:")
	fmt.Println(bst.InOrder())
	fmt.Println("后序遍历:")
	fmt.Println(bst.PostOrder())
	fmt.Println("查找 21 是否存在: ", bst.Search(21))
	fmt.Println("删除 21, 继续查看前序遍历结果: ")
	bst.Delete(21)
	fmt.Println(bst.PreOrder())
	fmt.Println("最大深度: ", bst.Depth())
	fmt.Println("最小深度: ", bst.DepthMin())
}
