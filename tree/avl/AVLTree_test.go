package avl

import (
	"fmt"
	"testing"
)

func TestAVL_Add(t *testing.T) {
	var avl = NewAVL()

	avl.Add(11)
	avl.Add(22)
	avl.Add(6)
	avl.Add(16)
	avl.Add(3)
	avl.Add(9)
	avl.Add(33)
	avl.Add(5)
	avl.Add(2)
	avl.Add(15)
	avl.Add(7)
	avl.Add(17)
	avl.Add(3)
	avl.Add(1)
	avl.Add(10)

	avl.Show()

	t.Log(avl.Size())
	t.Log("deep: ", avl.Depth())
	t.Log("width: ", avl.Width())

}

func TestAvlNode_LL_Rotate(t *testing.T) {
	var avl = NewAVL()

	avl.Add(11)
	avl.Add(6)
	avl.Add(16)

	// L
	avl.Add(5)
	//avl.Add(8)
	// LL
	avl.Add(4)
	avl.Add(3)
	avl.Add(2)
	avl.Add(1)

	// 旋转
	fmt.Println("-------旋转之后的树--------")
	avl.Show()
}
func TestAvlNode_RR_Rotate(t *testing.T) {
	var avl = NewAVL()

	avl.Add(11)
	avl.Add(6)
	avl.Add(16)

	// R
	avl.Add(14)
	avl.Add(18)
	// RR
	avl.Add(17)
	avl.Add(19)

	// 旋转
	fmt.Println("-------旋转之后的树--------")
	avl.Show()
}
func TestAvlNode_LR_Rotate(t *testing.T) {
	var avl = NewAVL()

	avl.Add(11)
	avl.Add(6)
	avl.Add(16)

	// L
	avl.Add(4)
	avl.Add(8)
	// LR
	avl.Add(7)
	avl.Add(9)

	// 旋转
	fmt.Println("-------旋转之后的树--------")
	avl.Show()
}
func TestAvlNode_RL_Rotate(t *testing.T) {
	var avl = NewAVL()

	avl.Add(11)
	avl.Add(6)
	avl.Add(16)

	// R
	avl.Add(14)
	avl.Add(18)
	// RL
	avl.Add(13)
	avl.Add(15)

	// 旋转
	fmt.Println("-------旋转之后的树--------")
	avl.Show()
}
