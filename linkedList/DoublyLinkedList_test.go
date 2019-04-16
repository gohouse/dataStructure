package linkedList

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	dl := NewDoublyLinkedList()
	dl.Add(3)
	dl.Add(9)
	dl.Add(92)
	dl.AddToHead(88)
	dl.AddToTail(44)
	dl.AddAt(2,1)

	dl.Show()
	fmt.Println(dl.Length())

	dl.Delete(3)
	dl.Show()
	fmt.Println(dl.Length())

	fmt.Println(dl.Get(3).data)
}
