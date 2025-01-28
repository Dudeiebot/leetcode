package heap

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Len() int { return len(*h) }

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }

func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func LastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}

	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}

	for maxHeap.Len() > 1 {
		stone1 := heap.Pop(maxHeap).(int)
		stone2 := heap.Pop(maxHeap).(int)

		if stone1 != stone2 {
			heap.Push(maxHeap, stone1-stone2)
		}
	}

	res := 0
	if maxHeap.Len() == 1 {
		res = (*maxHeap)[0]
	}
	return res
}
