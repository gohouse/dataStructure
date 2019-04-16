package binaryHeap

import "fmt"

type IBinaryHeap interface {
	Add(data int)
	AddSlice(data []int)
	Length() int
	BuildMax(dataLen int)
	BuildMin(dataLen int)
	SortAsc()
	SortDesc()
	Show()
}

type BinaryHeap struct {
	data []int
}

func (b *BinaryHeap) Add(data int) {
	if b.Length() == 0 {
		b.data = append(b.data, 1)
	}
	b.data = append(b.data, data)
	b.data[0]++
}

func (b *BinaryHeap) AddSlice(data []int) {
	for _, item := range data {
		b.Add(item)
	}
}

func (b *BinaryHeap) Length() int {
	return len(b.data)
}

// BuildMax : 构建大根堆
func (b *BinaryHeap) BuildMax(dataLen int) {
	// 循环所有的父节点(从后向前循环)
	for i := (dataLen - 1) / 2; i >= 1; i-- {
		// 找出最大的儿子节点
		var maxIndex = i * 2 // 假设最大的就是左儿子(因为是父节点,必定存在左儿子)
		// 如果有右儿子, 并且右儿子更大,则记录
		if (maxIndex+1) < dataLen && b.data[maxIndex+1] > b.data[maxIndex] {
			maxIndex = i*2 + 1
		}
		// 与父节点比较, 如果更大就交换
		for b.data[maxIndex] > b.data[i] {
			b.data[maxIndex], b.data[i] = b.data[i], b.data[maxIndex]
		}
	}
}

// BuildMin : 构建小根堆
func (b *BinaryHeap) BuildMin(dataLen int) {
	// 循环所有的父节点(从后向前循环)
	for i := (dataLen - 1) / 2; i >= 1; i-- {
		// 找出最大的儿子节点
		var minIndex = i * 2 // 假设最大的就是左儿子(因为是父节点,必定存在左儿子)
		// 如果有右儿子, 并且右儿子更大,则记录
		if (minIndex+1) < dataLen && b.data[minIndex+1] < b.data[minIndex] {
			minIndex++
		}
		// 与父节点比较, 如果更小就交换
		for b.data[minIndex] < b.data[i] {
			b.data[minIndex], b.data[i] = b.data[i], b.data[minIndex]
		}
	}
}

// SortRise : 对堆做升序排列
// 1. 组合大顶堆
// 2. 根节点与最后一个叶子节点互换
// 3. 在除最后一个叶子节点的堆中继续执行1和2步骤, 到堆中剩余1个节点为止
func (b *BinaryHeap) SortAsc() {
	var dataLen = len(b.data)
	for dataLen>2 {
		b.BuildMax(dataLen)

		b.data[1],b.data[dataLen-1] = b.data[dataLen-1],b.data[1]

		dataLen--
	}
}

// SortRise : 对堆做降序排列
func (b *BinaryHeap) SortDesc() {
	var dataLen = len(b.data)
	for dataLen>2 {
		b.BuildMin(dataLen)

		b.data[1],b.data[dataLen-1] = b.data[dataLen-1],b.data[1]

		dataLen--
	}
}

func (b *BinaryHeap) Show() {
	for k, item := range b.data {
		if k>0 {
			fmt.Printf("%v,", item)
		}
	}
	fmt.Println("")
}
