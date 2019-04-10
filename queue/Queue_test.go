package queue

import "testing"

func TestQueue(t *testing.T) {
	var q = NewQueue()
	q.EnQueue(q.Data(1))
	q.EnQueue(q.Data(4))
	q.EnQueue(&Node{Data:8})
	q.EnQueue(q.Data(2))

	q.Show()
	//return

	t.Log(q.Peek())
	t.Log(q.DeQueue())
	t.Log(q.Peek())

	t.Log(q.Size())

	q.Show()
}
