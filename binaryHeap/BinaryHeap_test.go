package binaryHeap

import (
	"fmt"
	"testing"
)

func TestBinaryHeap(t *testing.T) {

	var data = &BinaryHeap{data:[]int{6,9,2,22,44,8,12}}

	data.Show()
	data.Add(88)
	data.AddSlice([]int{3,28})
	data.Show()

	fmt.Println("\n----------max")
	data.BuildMax(data.Length())
	data.Show()

	fmt.Println("\n----------min")
	data.BuildMin(data.Length())
	data.Show()

	fmt.Println("\n----------asc")
	data.SortAsc()
	data.Show()

	fmt.Println("\n----------desc")
	data.SortDesc()
	data.Show()
}
