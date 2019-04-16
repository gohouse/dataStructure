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


	//3. 前序遍历，中序遍历，后序遍历
	//4.分层遍历二叉树（按层次从上往下，从左往右）
	//5. 将二叉查找树变为有序的双向链表
	//6. 求二叉树第K层的节点个数
	//7. 求二叉树中叶子节点的个数
	//8. 判断两棵二叉树是否结构相同
	//9. 判断二叉树是不是平衡二叉树
	//10. 求二叉树的镜像
	//11. 求二叉树中两个节点的最低公共祖先节点
	//12. 求二叉树中节点的最大距离
	//13. 由前序遍历序列和中序遍历序列重建二叉树
	//14.判断二叉树是不是完全二叉树
}

type BinarySearchTree struct {
	Data        int
	Left, Right *BinarySearchTree
}

//var _ IBinarySearchTree = &BinarySearchTree{}

func New() *BinarySearchTree {
	return new(BinarySearchTree)
}

func (b *BinarySearchTree) Depth() int {
	if b.Data==0 {
		return 0
	}

	return 1+max(b.Left.Depth(), b.Right.Depth())
}

func (b *BinarySearchTree) DepthMin() int {
	if b.Data==0 {
		return 0
	}

	return 1+min(b.Left.Depth(), b.Right.Depth())
}

func max(a,b int) int {
	if a>=b {
		return a
	} else {
		return b
	}
}

func min(a,b int) int {
	if a<b {
		return a
	} else {
		return b
	}
}
func (b *BinarySearchTree) Insert(i int) bool {
	if b.Data == 0 {
		b.Data = i
		b.Left = &BinarySearchTree{}
		b.Right = &BinarySearchTree{}
		return true
	}
	if b.Data > i {
		b.Left.Insert(i)
	} else {
		b.Right.Insert(i)
	}
	return false
}

func (b *BinarySearchTree) Search(i int) *BinarySearchTree {
	if b == nil {
		return nil
	}
	if b.Data == i {
		return b
	}
	if b.Data > i {
		return b.Left.Search(i)
	} else {
		return b.Right.Search(i)
	}
}

func (b *BinarySearchTree) MaxNode() *BinarySearchTree {
	if b.Right.Data == 0 {
		return b
	} else {
		return b.Right.MaxNode()
	}
}

func (b *BinarySearchTree) MinNode() *BinarySearchTree {
	if b.Left.Data == 0 {
		return b
	} else {
		return b.Left.MinNode()
	}
}

func (b *BinarySearchTree) PreOrder() []int {
	var res []int
	if b.Data != 0 {
		//fmt.Printf("%v ", b.Data)
		res = append(res, b.Data)
		res = append(res, b.Left.PreOrder()...)
		res = append(res,b.Right.PreOrder()...)
	}

	return res
}

func (b *BinarySearchTree) InOrder() []int {
	var res []int
	if b.Data != 0 {
		res = append(res, b.Left.PreOrder()...)
		//fmt.Printf("%v ", b.Data)
		res = append(res, b.Data)
		res = append(res,b.Right.PreOrder()...)
	}
	return res
}

func (b *BinarySearchTree) PostOrder() []int {
	var res []int
	if b.Data != 0 {
		res = append(res, b.Left.PreOrder()...)
		res = append(res,b.Right.PreOrder()...)
		//fmt.Printf("%v ", b.Data)
		res = append(res, b.Data)
	}
	return res
}

func (b *BinarySearchTree) Delete(i int) bool {
	// 查找
	findTreeNode := b.Search(i)
	if findTreeNode.Data == 0 {
		return false
	}

	// 根据不同情况删除, 1:没有子节点; 2:有一个子节点; 3:有两个子节点
	if findTreeNode.Left.Data == 0 && findTreeNode.Right.Data == 0 {
		// 没有子节点, 直接销毁该节点
		findTreeNode.clear()
	} else if findTreeNode.Left.Data == 0 && findTreeNode.Right.Data != 0 {
		// 只有右节点,当前节点替换为右节点,然后销毁右节点
		findTreeNode.fillByTreeNode(findTreeNode.Right)
		findTreeNode.Right.clear()
	} else if findTreeNode.Left.Data != 0 && findTreeNode.Right.Data == 0 {
		// 只有左节点,当前节点替换为右节点,然后销毁左节点
		findTreeNode.fillByTreeNode(findTreeNode.Left)
		findTreeNode.Left.clear()
	} else {	// 两个子节点都存在
		// 查找右子树中最小的节点
		rightMin := findTreeNode.Right.MinNode()
		// 用右子树的节点值,替换需要删除的节点值
		findTreeNode.Data = rightMin.Data
		// 如果右子树的最小值有右节点(不可能有左节点,因为左节点比该节点还小),则将此右节点迁移到最小值节点
		if rightMin.Right != nil {
			rightMin.fillByTreeNode(rightMin.Right)
		} else {
			rightMin.clear()
		}
	}
	return true
}

// 销毁叶子节点
func (b *BinarySearchTree) clear() {
	b.Data = 0
	b.Left = nil
	b.Right = nil
}

// 填充为另一个节点的值
func (b *BinarySearchTree) fillByTreeNode(node *BinarySearchTree) {
	b.Data = node.Data
	b.Left = node.Left
	b.Right = node.Right
}
