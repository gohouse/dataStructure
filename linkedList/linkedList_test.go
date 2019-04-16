package linkedList

import "testing"

func TestLinkedList_AddToFirst(t *testing.T) {
	var list LinkedList
	list.AddToFirst(1)
	list.AddToFirst(2)
	list.AddToLast(3)

	t.Log("打印初始链表：", list.Show())

	t.Log("打印index为1的值", list.Get(1).data)

	list.Add(1,5)

	t.Log("打印插入index为1，值为5的链表：",list.Show())

	list.Delete(2)

	t.Log("打印删除index为2的链表:",list.Show())

	reverse := list.Reverse()
	t.Log("反转链表结果： ",reverse.Show())
}
