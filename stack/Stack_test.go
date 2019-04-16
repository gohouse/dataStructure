package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	s.Push(3)
	s.Push(8)
	s.Push(82)
	s.Push(22)

	t.Log(s.Length())
	s.Show()

	s.Pop()
	s.Pop()
	s.Show()
	t.Log(s.Length())
}
